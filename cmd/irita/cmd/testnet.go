package cmd

// DONTCOVER

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"os"
	"path/filepath"

	evmhd "github.com/tharsis/ethermint/crypto/hd"

	"github.com/tendermint/tendermint/crypto/algo"

	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	ethermint "github.com/tharsis/ethermint/types"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tmconfig "github.com/tendermint/tendermint/config"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	"github.com/tendermint/tendermint/libs/tempfile"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"

	evmutils "github.com/bianjieai/irita/modules/evm/utils"

	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/bianjieai/iritamod/modules/genutil"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/perm"
	"github.com/bianjieai/iritamod/utils"

	evmosConfig "github.com/tharsis/ethermint/server/config"

	opbtypes "github.com/bianjieai/irita/modules/opb/types"
)

const (
	nodeDirPerm         = 0755
	DefaultPointDenom   = "point"
	DefaultPointMinUnit = "upoint"
)

var (
	flagNodeDirPrefix     = "node-dir-prefix"
	flagNumValidators     = "v"
	flagOutputDir         = "output-dir"
	flagNodeDaemonHome    = "node-daemon-home"
	flagNodeCLIHome       = "node-cli-home"
	flagStartingIPAddress = "starting-ip-address"
)

// get cmd to initialize all files for tendermint testnet and application
func testnetCmd(mbm module.BasicManager, genBalIterator banktypes.GenesisBalancesIterator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Initialize files for a irita testnet",
		Long: "testnet will create \"v\" number of directories and populate each with " +
			"necessary files (private validator, genesis, config, etc.).\n" +
			"Note, strict routability for addresses is turned off in the config file.",
		Example: "irita testnet --v 4 --output-dir ./output --starting-ip-address 192.168.10.2",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			outputDir := viper.GetString(flagOutputDir)
			chainID := viper.GetString(flags.FlagChainID)
			minGasPrices := viper.GetString(server.FlagMinGasPrices)
			nodeDirPrefix := viper.GetString(flagNodeDirPrefix)
			nodeDaemonHome := viper.GetString(flagNodeDaemonHome)
			nodeCLIHome := viper.GetString(flagNodeCLIHome)
			startingIPAddress := viper.GetString(flagStartingIPAddress)
			numValidators := viper.GetInt(flagNumValidators)
			algo, _ := cmd.Flags().GetString(flags.FlagKeyAlgorithm)

			return InitTestnet(
				clientCtx, cmd, config, mbm, genBalIterator, outputDir, chainID, minGasPrices,
				nodeDirPrefix, nodeDaemonHome, nodeCLIHome, startingIPAddress, numValidators, algo,
			)
		},
	}

	cmd.Flags().Int(flagNumValidators, 4, "Number of validators to initialize the testnet with")
	cmd.Flags().StringP(flagOutputDir, "o", "./mytestnet", "Directory to store initialization data for the testnet")
	cmd.Flags().String(flagNodeDirPrefix, "node", "Prefix the directory name for each node with (node results in node0, node1, ...)")
	cmd.Flags().String(flagNodeDaemonHome, "irita", "Home directory of the node's daemon configuration")
	cmd.Flags().String(flagNodeCLIHome, "iritacli", "Home directory of the node's cli configuration")
	cmd.Flags().String(flagStartingIPAddress, "192.168.0.1", "Starting IP address (192.168.0.1 results in persistent peers list ID0@192.168.0.1:46656, ID1@192.168.0.2:46656, ...)")
	cmd.Flags().String(flags.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(server.FlagMinGasPrices, fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom), "Minimum gas prices to accept for transactions; All fees in a tx must meet this minimum (e.g. 0.01photino,0.001stake)")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	cmd.Flags().String(flags.FlagKeyAlgorithm, string(hd.Sm2Type), "Key signing algorithm to generate keys for")
	return cmd
}

// Initialize the testnet
func InitTestnet(
	clientCtx client.Context, cmd *cobra.Command,
	config *tmconfig.Config, mbm module.BasicManager,
	genBalIterator banktypes.GenesisBalancesIterator,
	outputDir, chainID, minGasPrices, nodeDirPrefix,
	nodeDaemonHome, nodeCLIHome, startingIPAddress string,
	numValidators int, algoStr string,
) error {
	algo.Algo = algo.SM2
	if chainID == "" {
		chainID = fmt.Sprintf("chain_%d-1", tmrand.Int63n(9999999999999)+1)
	}
	evmutils.SetEthermintSupportedAlgorithms()

	monikers := make([]string, numValidators)
	nodeIDs := make([]string, numValidators)
	valCerts := make([]string, numValidators)

	iritaConfig := evmosConfig.DefaultConfig()
	iritaConfig.MinGasPrices = minGasPrices
	iritaConfig.API.Enable = true
	iritaConfig.Telemetry.Enabled = true
	iritaConfig.Telemetry.PrometheusRetentionTime = 60
	iritaConfig.Telemetry.EnableHostnameLabel = false
	iritaConfig.Telemetry.GlobalLabels = [][]string{{"chain_id", chainID}}

	//nolint:prealloc
	var (
		genAccounts []authtypes.GenesisAccount
		genBalances []banktypes.Balance
		genFiles    []string
	)

	inBuf := bufio.NewReader(cmd.InOrStdin())

	rootKeyPath := filepath.Join(outputDir, "root_key.pem")
	rootCertPath := filepath.Join(outputDir, "root_cert.pem")

	if err := os.MkdirAll(outputDir, nodeDirPerm); err != nil {
		_ = os.RemoveAll(outputDir)
		return err
	}

	utils.GenRootCert(rootKeyPath, rootCertPath, "/C=CN/ST=root/L=root/O=root/OU=root/CN=root")
	// generate private keys, node IDs, and initial transactions
	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		clientDir := filepath.Join(outputDir, nodeDirName, nodeCLIHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")

		config.SetRoot(nodeDir)
		config.RPC.ListenAddress = "tcp://0.0.0.0:26657"

		if err := os.MkdirAll(filepath.Join(nodeDir, "config"), nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		if err := os.MkdirAll(clientDir, nodeDirPerm); err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		monikers[i] = nodeDirName
		config.Moniker = nodeDirName

		ip, err := getIP(i, startingIPAddress)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		nodeKey, filePv, err := genutil.InitializeNodeValidatorFiles(config)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		nodeIDs[i] = string(nodeKey.ID())

		key, err := genutil.Genkey(filePv.Key.PrivKey)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}
		keyPath := filepath.Join(nodeDir, "config", "key.pem")
		cerPath := filepath.Join(nodeDir, "config", "cer.pem")
		certPath := filepath.Join(nodeDir, "config", "cert.pem")
		if err = tempfile.WriteFileAtomic(keyPath, key, 0600); err != nil {
			return err
		}

		utils.GenCertRequest(keyPath, cerPath, fmt.Sprintf("/C=CN/ST=test/L=test/O=test/OU=test/CN=%s", nodeDirName))
		utils.IssueCert(cerPath, rootCertPath, rootKeyPath, certPath)
		valCerts[i] = certPath

		memo := fmt.Sprintf("%s@%s:26656", nodeIDs[i], ip)
		genFiles = append(genFiles, config.GenesisFile())

		kb, err := keyring.New(
			sdk.KeyringServiceName(),
			viper.GetString(flags.FlagKeyringBackend),
			clientDir,
			inBuf,
			evmhd.EthSecp256k1Option(),
		)
		if err != nil {
			return err
		}

		keyringAlgos, _ := kb.SupportedAlgorithms()
		signAlgo, err := keyring.NewSigningAlgoFromString(algoStr, keyringAlgos)
		if err != nil {
			return err
		}

		addr, secret, err := server.GenerateSaveCoinKey(kb, nodeDirName, true, signAlgo)
		if err != nil {
			_ = os.RemoveAll(outputDir)
			return err
		}

		info := map[string]string{"secret": secret}

		cliPrint, err := json.Marshal(info)
		if err != nil {
			return err
		}

		// save private key seed words
		if err := writeFile(fmt.Sprintf("%v.json", "key_seed"), clientDir, cliPrint); err != nil {
			return err
		}

		accTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
		accPointTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
		accNativeTokens := sdk.TokensFromConsensusPower(5000, sdk.DefaultPowerReduction)
		coins := sdk.Coins{
			sdk.NewCoin(fmt.Sprintf("%stoken", nodeDirName), accTokens),
			sdk.NewCoin(DefaultPointMinUnit, accPointTokens),
			sdk.NewCoin(tokentypes.GetNativeToken().MinUnit, accNativeTokens),
		}

		genBalances = append(genBalances, banktypes.Balance{Address: addr.String(), Coins: coins.Sort()})

		genAccounts = append(genAccounts, authtypes.NewBaseAccount(addr, nil, 0, 0))

		cert, err := ioutil.ReadFile(certPath)
		if err != nil {
			return err
		}
		msg := node.NewMsgCreateValidator(nodeDirName, nodeDirName, string(cert), 100, addr)

		txBuilder := clientCtx.TxConfig.NewTxBuilder()
		if err := txBuilder.SetMsgs(msg); err != nil {
			return err
		}

		txBuilder.SetMemo(memo)

		txFactory := tx.Factory{}
		txFactory = txFactory.
			WithChainID(chainID).
			WithMemo(memo).
			WithKeybase(kb).
			WithTxConfig(clientCtx.TxConfig)

		if err := tx.Sign(txFactory, nodeDirName, txBuilder, true); err != nil {
			return err
		}

		txBz, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
		if err != nil {
			return err
		}

		if err := writeFile(fmt.Sprintf("%v.json", nodeDirName), gentxsDir, txBz); err != nil {
			return err
		}

		customAppTemplate, customAppConfig := evmosConfig.AppConfig(ethermint.AttoPhoton)
		srvconfig.SetConfigTemplate(customAppTemplate)
		if err := server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig); err != nil {
			return err
		}

		iritaConfigFilePath := filepath.Join(nodeDir, "config/app.toml")
		srvconfig.WriteConfigFile(iritaConfigFilePath, iritaConfig)
	}

	if err := initGenFiles(DefaultPointMinUnit, clientCtx, mbm, chainID, genAccounts, genBalances, genFiles, numValidators,
		monikers, nodeIDs, rootCertPath); err != nil {
		return err
	}

	if err := collectGenFiles(
		clientCtx, config, chainID, monikers, nodeIDs, valCerts, numValidators,
		outputDir, nodeDirPrefix, nodeDaemonHome,
	); err != nil {
		return err
	}

	cmd.PrintErrf("Successfully initialized %d node directories\n", numValidators)
	return nil
}

func initGenFiles(
	coinDenom string, clientCtx client.Context, mbm module.BasicManager, chainID string,
	genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance,
	genFiles []string, numValidators int, monikers []string, nodeIDs []string,
	rootCertPath string,
) error {
	rootCertBz, err := ioutil.ReadFile(rootCertPath)
	if err != nil {
		return fmt.Errorf("failed to read root certificate: %s", err.Error())
	}
	rootCert := string(rootCertBz)

	jsonMarshaler := clientCtx.Codec

	appGenState := mbm.DefaultGenesis(jsonMarshaler)

	var nodeGenState node.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[node.ModuleName], &nodeGenState)

	nodeGenState.RootCert = rootCert

	nodeGenState.Nodes = make([]node.Node, len(nodeIDs))
	for i, nodeID := range nodeIDs {
		nodeGenState.Nodes[i].Id = nodeID
		nodeGenState.Nodes[i].Name = monikers[i]
	}

	appGenState[node.ModuleName] = jsonMarshaler.MustMarshalJSON(&nodeGenState)

	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return err
	}

	authGenState.Accounts = accounts
	appGenState[authtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = genBalances
	appGenState[banktypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&bankGenState)

	// set the point token in the genesis state
	var tokenGenState tokentypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[tokentypes.ModuleName], &tokenGenState)

	pointToken := tokentypes.Token{
		Symbol:        DefaultPointDenom,
		Name:          "Irita point token",
		Scale:         6,
		MinUnit:       DefaultPointMinUnit,
		InitialSupply: 1000000000,
		MaxSupply:     math.MaxUint64,
		Mintable:      true,
		Owner:         genAccounts[0].GetAddress().String(),
	}

	tokenGenState.Tokens = append(tokenGenState.Tokens, pointToken)
	tokenGenState.Params.IssueTokenBaseFee = sdk.NewCoin(DefaultPointDenom, sdk.NewInt(60000))
	appGenState[tokentypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&tokenGenState)

	// modify the native token denoms in the opb genesis
	var opbGenState opbtypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[opbtypes.ModuleName], &opbGenState)

	opbGenState.Params.BaseTokenDenom = tokentypes.GetNativeToken().MinUnit
	opbGenState.Params.PointTokenDenom = DefaultPointMinUnit
	appGenState[opbtypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&opbGenState)

	// modify the constant fee denoms in the crisis genesis
	var crisisGenState crisistypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[crisistypes.ModuleName], &crisisGenState)

	crisisGenState.ConstantFee.Denom = tokentypes.GetNativeToken().MinUnit
	appGenState[crisistypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&crisisGenState)

	// modify the constant fee denoms in the crisis genesis
	var serviceGenState servicetypes.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[servicetypes.ModuleName], &serviceGenState)

	serviceGenState.Params.MinDeposit = sdk.NewCoins(sdk.NewCoin(DefaultPointMinUnit, sdk.NewInt(5000)))
	serviceGenState.Params.BaseDenom = DefaultPointMinUnit
	appGenState[servicetypes.ModuleName] = jsonMarshaler.MustMarshalJSON(&serviceGenState)

	var evmGenState evmtypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appGenState[evmtypes.ModuleName], &evmGenState)

	evmGenState.Params.EvmDenom = coinDenom
	appGenState[evmtypes.ModuleName] = clientCtx.Codec.MustMarshalJSON(&evmGenState)

	// add all genesis accounts as root admins
	var permGenState perm.GenesisState
	jsonMarshaler.MustUnmarshalJSON(appGenState[perm.ModuleName], &permGenState)
	for _, account := range genAccounts {
		permGenState.RoleAccounts = append(
			permGenState.RoleAccounts,
			perm.RoleAccount{
				Address: account.GetAddress().String(),
				Roles:   []perm.Role{perm.RoleRootAdmin},
			},
		)
	}
	appGenState[perm.ModuleName] = jsonMarshaler.MustMarshalJSON(&permGenState)

	appGenStateJSON, err := json.MarshalIndent(appGenState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < numValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}
	return nil
}

func collectGenFiles(
	clientCtx client.Context, config *tmconfig.Config, chainID string,
	monikers, nodeIDs []string, valCerts []string,
	numValidators int, outputDir, nodeDirPrefix, nodeDaemonHome string,
) error {
	var appState json.RawMessage
	genTime := tmtime.Now()

	for i := 0; i < numValidators; i++ {
		nodeDirName := fmt.Sprintf("%s%d", nodeDirPrefix, i)
		nodeDir := filepath.Join(outputDir, nodeDirName, nodeDaemonHome)
		gentxsDir := filepath.Join(outputDir, "gentxs")
		moniker := monikers[i]
		config.Moniker = nodeDirName

		config.SetRoot(nodeDir)

		initCfg := genutil.NewInitConfig(chainID, gentxsDir, moniker, nodeIDs[i])

		genDoc, err := types.GenesisDocFromFile(config.GenesisFile())
		if err != nil {
			return err
		}

		nodeAppState, err := genutil.GenAppStateFromConfig(clientCtx.Codec, clientCtx.TxConfig, config, initCfg, *genDoc)
		if err != nil {
			return err
		}

		if appState == nil {
			// set the canonical application state (they should not differ)
			appState = nodeAppState
		}

		genFile := config.GenesisFile()

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, chainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

func getIP(i int, startingIPAddr string) (ip string, err error) {
	if len(startingIPAddr) == 0 {
		ip, err = server.ExternalIP()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return calculateIP(startingIPAddr, i)
}

func calculateIP(ip string, i int) (string, error) {
	ipv4 := net.ParseIP(ip).To4()
	if ipv4 == nil {
		return "", fmt.Errorf("%v: non ipv4 address", ip)
	}

	for j := 0; j < i; j++ {
		ipv4[3]++
	}

	return ipv4.String(), nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	if err := tmos.EnsureDir(writePath, 0700); err != nil {
		return err
	}

	return tmos.WriteFile(file, contents, 0600)
}
