package app

import (
	"io"
	"os"

	"cosmossdk.io/depinject"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/streaming"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	cosmosparamstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"

	identitykeeper "iritamod.bianjie.ai/modules/identity/keeper"
	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	paramskeeper "iritamod.bianjie.ai/modules/params/keeper"
	slashingkeeper "iritamod.bianjie.ai/modules/slashing/keeper"
	upgradekeeper "iritamod.bianjie.ai/modules/upgrade/keeper"

	mtkeeper "mods.irisnet.org/modules/mt/keeper"
	nftkeeper "mods.irisnet.org/modules/nft/keeper"
	oraclekeeper "mods.irisnet.org/modules/oracle/keeper"
	randomkeeper "mods.irisnet.org/modules/random/keeper"
	recordkeeper "mods.irisnet.org/modules/record/keeper"
	servicekeeper "mods.irisnet.org/modules/service/keeper"
	tokenkeeper "mods.irisnet.org/modules/token/keeper"

	"github.com/bianjieai/irita/crypto/hd"
	"github.com/bianjieai/irita/wrapper"
)

var _ servertypes.Application = (*IritaAppV2)(nil)

// IritaAppV2 extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type IritaAppV2 struct {
	*runtime.App
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	txConfig          client.TxConfig
	interfaceRegistry codectypes.InterfaceRegistry

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	RecordKeeper          recordkeeper.Keeper
	TokenKeeper           tokenkeeper.Keeper
	NftKeeper             nftkeeper.Keeper
	MtKeeper              mtkeeper.Keeper
	ServiceKeeper         servicekeeper.Keeper
	OracleKeeper          oraclekeeper.Keeper
	RandomKeeper          randomkeeper.Keeper
	IdentityKeeper        identitykeeper.Keeper
	NodeKeeper            *nodekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	EvmKeeper             *evmkeeper.Keeper
	FeeMarketKeeper       feemarketkeeper.Keeper

	// simulation manager
	sm *module.SimulationManager
}

// NewIritaAppV2 creates a new instance of the IritaAppV2 struct.
//
// Parameters:
// - logger: a log.Logger instance used for logging.
// - db: a dbm.DB instance used for the database.
// - traceStore: an io.Writer used for tracing.
// - loadLatest: a boolean indicating whether to load the latest state.
// - depInjectOptions: a DepinjectOptions instance.
// - appOpts: a servertypes.AppOptions instance.
// - baseAppOptions: a variadic list of functions used to configure the base app.
//
// Returns:
// - *IritaAppV2: a pointer to the newly created IritaAppV2 instance.
func NewIritaAppV2(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	depInjectOptions DepinjectOptions,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *IritaAppV2 {
	var (
		app        = &IritaAppV2{}
		appBuilder *runtime.AppBuilder

		providers = append(depInjectOptions.Providers,
			appOpts,
		)
		// merge the AppConfig and other configuration in one config
		appConfig = depinject.Configs(
			depInjectOptions.Config,
			depinject.Provide(
				wrapper.ProvideSlashingStakingKeeper,
				wrapper.ProvideEvidenceStakingKeeper,
				wrapper.ProvideEvmStakingKeeper,
				wrapper.ProvideEVMKeeper,
				wrapper.ProvideICS20Keeper,
				wrapper.ProvideEvmConstructor,
				wrapper.ProvideStakingHooks,
			),
			depinject.Supply(
				providers...,

			// ADVANCED CONFIGURATION

			//
			// AUTH
			//
			// For providing a custom function required in auth to generate custom account types
			// add it below. By default the auth module uses simulation.RandomGenesisAccounts.
			//
			// authtypes.RandomGenesisAccountsFn(simulation.RandomGenesisAccounts),

			// For providing a custom a base account type add it below.
			// By default the auth module uses authtypes.ProtoBaseAccount().
			//
			// func() authtypes.AccountI { return authtypes.ProtoBaseAccount() },

			//
			// MINT
			//

			// For providing a custom inflation function for x/mint add here your
			// custom function that implements the minttypes.InflationCalculationFn
			// interface.

			),
		)
	)
	hd.SetSupportedAlgorithms()

	consumer := append(depInjectOptions.Consumers,
		&appBuilder,
		&app.appCodec,
		&app.legacyAmino,
		&app.txConfig,
		&app.interfaceRegistry,
		&app.AccountKeeper,
		&app.BankKeeper,
		&app.CrisisKeeper,
		&app.CapabilityKeeper,
		&app.NodeKeeper,
		&app.SlashingKeeper,
		&app.CrisisKeeper,
		&app.UpgradeKeeper,
		&app.ParamsKeeper,
		&app.EvidenceKeeper,
		&app.FeeGrantKeeper,
		&app.RecordKeeper,
		&app.TokenKeeper,
		&app.MtKeeper,
		&app.NftKeeper,
		&app.ServiceKeeper,
		&app.OracleKeeper,
		&app.RandomKeeper,
		&app.IdentityKeeper,
		&app.ConsensusParamsKeeper,
		&app.EvmKeeper,
		&app.FeeMarketKeeper,
	)

	if err := depinject.Inject(appConfig, consumer...); err != nil {
		panic(err)
	}

	// Below we could construct and set an application specific mempool and
	// ABCI 1.0 PrepareProposal and ProcessProposal handlers. These defaults are
	// already set in the SDK's BaseApp, this shows an example of how to override
	// them.
	//
	// Example:
	//
	// app.App = appBuilder.Build(...)
	// nonceMempool := mempool.NewSenderNonceMempool()
	// abciPropHandler := NewDefaultProposalHandler(nonceMempool, app.App.BaseApp)
	//
	// app.App.BaseApp.SetMempool(nonceMempool)
	// app.App.BaseApp.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// app.App.BaseApp.SetProcessProposal(abciPropHandler.ProcessProposalHandler())
	//
	// Alternatively, you can construct BaseApp options, append those to
	// baseAppOptions and pass them to the appBuilder.
	//
	// Example:
	//
	// prepareOpt = func(app *baseapp.BaseApp) {
	// 	abciPropHandler := baseapp.NewDefaultProposalHandler(nonceMempool, app)
	// 	app.SetPrepareProposal(abciPropHandler.PrepareProposalHandler())
	// }
	// baseAppOptions = append(baseAppOptions, prepareOpt)

	app.App = appBuilder.Build(logger, db, traceStore, baseAppOptions...)

	// load state streaming if enabled
	if _, _, err := streaming.LoadStreamingServices(app.App.BaseApp, appOpts, app.appCodec, logger, app.kvStoreKeys()); err != nil {
		logger.Error("failed to load state streaming", "err", err)
		os.Exit(1)
	}

	// app.initParamsKeeper()

	/****  Module Options ****/

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)

	// RegisterUpgradeHandlers is used for registering any on-chain upgrades.
	// app.RegisterUpgradeHandlers()

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	overrideModules := map[string]module.AppModuleSimulation{
		authtypes.ModuleName: auth.NewAppModule(
			app.appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.getSubspace(authtypes.ModuleName),
		),
	}
	app.sm = module.NewSimulationManagerFromAppModules(app.ModuleManager.Modules, overrideModules)

	app.sm.RegisterStoreDecoders()
	app.SetInitChainer(app.InitChainer)

	// A custom InitChainer can be set if extra pre-init-genesis logic is required.
	// By default, when using app wiring enabled module, this is not required.
	// For instance, the upgrade module will set automatically the module version map in its init genesis thanks to app wiring.
	// However, when registering a module manually (i.e. that does not support app wiring), the module version map
	// must be set manually as follow. The upgrade module will de-duplicate the module version map.
	//
	// app.SetInitChainer(func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	// 	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	// 	return app.App.InitChainer(ctx, req)
	// })

	if err := app.Load(loadLatest); err != nil {
		panic(err)
	}

	return app
}

func (app *IritaAppV2) kvStoreKeys() map[string]*storetypes.KVStoreKey {
	keys := make(map[string]*storetypes.KVStoreKey)
	for _, k := range app.GetStoreKeys() {
		if kv, ok := k.(*storetypes.KVStoreKey); ok {
			keys[kv.Name()] = kv
		}
	}

	return keys
}

// getSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IritaAppV2) getSubspace(moduleName string) cosmosparamstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}