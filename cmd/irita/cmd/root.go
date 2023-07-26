package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

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

	genutilcli "github.com/bianjieai/iritamod/modules/genutil/client/cli"
	"github.com/bianjieai/iritamod/modules/node"

	ethermintclient "github.com/evmos/ethermint/client"
	evmclient "github.com/evmos/ethermint/client"
	"github.com/evmos/ethermint/crypto/hd"
	evmserver "github.com/evmos/ethermint/server"
	servercfg "github.com/evmos/ethermint/server/config"
	ethermint "github.com/evmos/ethermint/types"

	"github.com/bianjieai/irita/app"
	"github.com/bianjieai/irita/encoding"
)

// NewRootCmd creates a new root command for simd. It is called once in the main function.
func NewRootCmd() (*cobra.Command, encoding.Config) {
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
		WithKeyringOptions(hd.EthSecp256k1Option()).
		WithViper("")

	rootCmd := &cobra.Command{
		Use:   "irita",
		Short: "Irita app command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// set the default command outputs
			cmd.SetOut(cmd.OutOrStdout())
			cmd.SetErr(cmd.ErrOrStderr())

			initClientCtx, err := client.ReadPersistentCommandFlags(initClientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			initClientCtx, err = config.ReadFromClientConfig(initClientCtx)
			if err != nil {
				return err
			}
			if err = client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			// TODO: define our own token
			customAppTemplate, customAppConfig := servercfg.AppConfig(ethermint.AttoPhoton)

			handleRequestPreRun(cmd, args)
			handleResponsePreRun(cmd)
			return server.InterceptConfigsPreRunHandler(
				cmd,
				customAppTemplate,
				customAppConfig,
				tmcfg.DefaultConfig(),
			)
		},
		PersistentPostRun: func(cmd *cobra.Command, _ []string) {
			handleResponsePostRun(encodingConfig.Codec, cmd)
		},
	}
	cfg := sdk.GetConfig()
	cfg.Seal()

	initRootCmd(rootCmd, encodingConfig)
	return rootCmd, encodingConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig encoding.Config) {
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
		AddGenesisValidatorCmd(
			app.ModuleBasics,
			node.AppModuleBasic{},
			app.DefaultNodeHome,
			app.DefaultNodeHome,
		),
		genutilcli.GenKey(app.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
		config.Cmd(),
		SnapshotCmd(),
	)

	ac := appCreator{encodingConfig}
	evmserver.AddCommands(
		rootCmd,
		evmserver.NewDefaultStartOptions(ac.newApp, app.DefaultNodeHome),
		ac.appExport,
		addModuleInitFlags,
	)

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
	encCfg encoding.Config
}

func (ac appCreator) newApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	baseappOptions := server.DefaultBaseappOptions(appOpts)
	return app.NewIritaApp(
		logger,
		db,
		traceStore,
		true,
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
	baseappOptions := server.DefaultBaseappOptions(appOpts)
	iritaApp := app.NewIritaApp(
		logger,
		db,
		traceStore,
		true,
		appOpts,
		baseappOptions...,
	)

	if height != -1 {
		if err := iritaApp.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return iritaApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
}
