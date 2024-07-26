package cmd

import (
	"io"
	"os"

	"cosmossdk.io/simapp/params"
	dbm "github.com/cometbft/cometbft-db"
	tmcfg "github.com/cometbft/cometbft/config"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	ethermintclient "github.com/evmos/ethermint/client"
	"github.com/evmos/ethermint/encoding"
	ethermintserver "github.com/evmos/ethermint/server"
	servercfg "github.com/evmos/ethermint/server/config"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	genutilcli "iritamod.bianjie.ai/modules/genutil/client/cli"
	"iritamod.bianjie.ai/modules/node"

	"github.com/bianjieai/irita/app"
	"github.com/bianjieai/irita/crypto/hd"
	evmclient "github.com/bianjieai/irita/modules/evm/client"
)

// NewRootCmd creates a new root command for simd. It is called once in the main function.
func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	//encodingConfig := app.MakeEncodingConfig()
	hd.SetSupportedAlgorithms()
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)

	initClientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastSync).
		WithHomeDir(app.DefaultNodeHome).
		WithKeyringOptions(hd.KeyringOption()).
		WithViper("")

	rootCmd := &cobra.Command{
		Use:   "irita",
		Short: "Irita app command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Open when debug
			//algo.Algo = algo.SM2
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}
			if err = client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			// TODO: define our own token
			customAppTemplate, customAppConfig := servercfg.AppConfig(ethermint.AttoPhoton)
			return server.InterceptConfigsPreRunHandler(cmd, customAppTemplate, customAppConfig, tmcfg.DefaultConfig())
		},
	}
	cfg := sdk.GetConfig()
	cfg.Seal()

	initRootCmd(rootCmd, encodingConfig)
	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	rootCmd.AddCommand(
		ethermintclient.ValidateChainID(
			genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		),
		//genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		//genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		//genutilcli.MigrateGenesisCmd(),
		GenRootCert(app.DefaultNodeHome),
		//genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(app.DefaultNodeHome, app.DefaultNodeHome),
		AddGenesisValidatorCmd(app.ModuleBasics, node.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultNodeHome),
		genutilcli.GenKey(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
		config.Cmd(),
		NewSnapshotCmd(),
	)

	ac := appCreator{encodingConfig}
	ethermintserver.AddCommands(rootCmd, ethermintserver.NewDefaultStartOptions(ac.newApp, app.DefaultNodeHome), ac.appExport, addModuleInitFlags)

	// add keybase, auxiliary RPC, query, and tx child commands
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		evmclient.KeyCommands(app.DefaultNodeHome),
	)
}

func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
	app.AddStartFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
		flags.LineBreak,
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg params.EncodingConfig
}

func (ac appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)
	return app.NewIritaAppV2(
		logger,
		db,
		traceStore,
		true,
		app.DefaultDepinjectOptions(),
		appOpts,
		baseappOptions...,
	)
}

// createIrisappAndExport creates a new irisapp (optionally at a given height) and exports state.
func (ac appCreator) appExport(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (
	servertypes.ExportedApp, error,
) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}

	var loadLatest bool
	if height == -1 {
		loadLatest = true
	}

	iritaApp := app.NewIritaApp(
		logger,
		db,
		traceStore,
		loadLatest,
		ac.encCfg,
		appOpts,
	)

	if height != -1 {
		if err := iritaApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return iritaApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
