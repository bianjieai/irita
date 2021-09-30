package simapp

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cast"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	sdkupgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"

	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/irisnet/irismod/modules/nft"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	"github.com/irisnet/irismod/modules/oracle"
	oraclekeeper "github.com/irisnet/irismod/modules/oracle/keeper"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	"github.com/irisnet/irismod/modules/random"
	randomkeeper "github.com/irisnet/irismod/modules/random/keeper"
	randomtypes "github.com/irisnet/irismod/modules/random/types"
	"github.com/irisnet/irismod/modules/record"
	recordkeeper "github.com/irisnet/irismod/modules/record/keeper"
	recordtypes "github.com/irisnet/irismod/modules/record/types"
	"github.com/irisnet/irismod/modules/service"
	servicekeeper "github.com/irisnet/irismod/modules/service/keeper"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	"github.com/irisnet/irismod/modules/token"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/bianjieai/iritamod/modules/genutil"
	genutiltypes "github.com/bianjieai/iritamod/modules/genutil"
	"github.com/bianjieai/iritamod/modules/identity"
	identitykeeper "github.com/bianjieai/iritamod/modules/identity/keeper"
	identitytypes "github.com/bianjieai/iritamod/modules/identity/types"
	"github.com/bianjieai/iritamod/modules/node"
	nodekeeper "github.com/bianjieai/iritamod/modules/node/keeper"
	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	cparams "github.com/bianjieai/iritamod/modules/params"
	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
	permtypes "github.com/bianjieai/iritamod/modules/perm/types"
	cslashing "github.com/bianjieai/iritamod/modules/slashing"
	"github.com/bianjieai/iritamod/modules/upgrade"
	upgradekeeper "github.com/bianjieai/iritamod/modules/upgrade/keeper"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"

	"github.com/bianjieai/irita/lite"
	"github.com/bianjieai/irita/modules/opb"
	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	opbtypes "github.com/bianjieai/irita/modules/opb/types"

	tibcnfttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibcnfttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/keeper"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"

	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	tibckeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
	tibcmock "github.com/bianjieai/tibc-go/modules/tibc/testing/mock"
)

const appName = "SimApp"

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		params.AppModuleBasic{},
		cparams.AppModuleBasic{},
		crisis.AppModuleBasic{},
		cslashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		record.AppModuleBasic{},
		token.AppModuleBasic{},
		nft.AppModuleBasic{},
		service.AppModuleBasic{},
		oracle.AppModuleBasic{},
		random.AppModuleBasic{},
		perm.AppModuleBasic{},
		identity.AppModuleBasic{},
		node.AppModuleBasic{},
		opb.AppModuleBasic{},
		wasm.AppModuleBasic{},
		tibc.AppModuleBasic{},
		tibcnfttransfer.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:          nil,
		tokentypes.ModuleName:               {authtypes.Minter, authtypes.Burner},
		servicetypes.DepositAccName:         nil,
		servicetypes.RequestAccName:         nil,
		opbtypes.PointTokenFeeCollectorName: nil,
		tibcnfttypes.ModuleName:             nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{}
)

var _ simapp.App = (*SimApp)(nil)

// SimApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type SimApp struct {
	*baseapp.BaseApp
	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	CapabilityKeeper  *capabilitykeeper.Keeper
	AccountKeeper     authkeeper.AccountKeeper
	BankKeeper        bankkeeper.Keeper
	SlashingKeeper    slashingkeeper.Keeper
	CrisisKeeper      crisiskeeper.Keeper
	UpgradeKeeper     upgradekeeper.Keeper
	ParamsKeeper      paramskeeper.Keeper
	EvidenceKeeper    evidencekeeper.Keeper
	RecordKeeper      recordkeeper.Keeper
	TokenKeeper       tokenkeeper.Keeper
	NFTKeeper         nftkeeper.Keeper
	ServiceKeeper     servicekeeper.Keeper
	OracleKeeper      oraclekeeper.Keeper
	RandomKeeper      randomkeeper.Keeper
	PermKeeper        permkeeper.Keeper
	IdentityKeeper    identitykeeper.Keeper
	NodeKeeper        nodekeeper.Keeper
	OpbKeeper         opbkeeper.Keeper
	FeeGrantKeeper    feegrantkeeper.Keeper
	WasmKeeper        wasm.Keeper
	TIBCKeeper        *tibckeeper.Keeper // TIBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	NftTransferKeeper tibcnfttransferkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedTIBCKeeper     capabilitykeeper.ScopedKeeper
	ScopedTIBCMockKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".irita")
}

// NewSimApp returns a reference to an initialized NewSimApp.
func NewSimApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig simappparams.EncodingConfig, appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *SimApp {
	// TODO: Remove cdc in favor of appCodec once all modules are migrated.
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		slashingtypes.StoreKey,
		paramstypes.StoreKey,
		upgradetypes.StoreKey,
		feegrant.StoreKey,
		evidencetypes.StoreKey,
		recordtypes.StoreKey,
		tokentypes.StoreKey,
		nfttypes.StoreKey,
		servicetypes.StoreKey,
		oracletypes.StoreKey,
		randomtypes.StoreKey,
		permtypes.StoreKey,
		identitytypes.StoreKey,
		nodetypes.StoreKey,
		opbtypes.StoreKey,
		wasm.StoreKey,
		capabilitytypes.StoreKey,
		tibchost.StoreKey,
		tibcnfttypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &SimApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
	)
	app.NodeKeeper = node.NewKeeper(appCodec, keys[node.StoreKey], app.GetSubspace(node.ModuleName))
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &app.NodeKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)

	sdkUpgradeKeeper := sdkupgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(sdkUpgradeKeeper)

	// create evidence keeper with router
	EvidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.NodeKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *EvidenceKeeper

	app.TokenKeeper = tokenkeeper.NewKeeper(
		appCodec, keys[tokentypes.StoreKey], app.GetSubspace(tokentypes.ModuleName),
		app.BankKeeper, app.ModuleAccountAddrs(), opbtypes.PointTokenFeeCollectorName,
	)
	app.RecordKeeper = recordkeeper.NewKeeper(appCodec, keys[recordtypes.StoreKey])
	app.NFTKeeper = nftkeeper.NewKeeper(appCodec, keys[nfttypes.StoreKey])

	app.ServiceKeeper = servicekeeper.NewKeeper(
		appCodec, keys[servicetypes.StoreKey], app.AccountKeeper, app.BankKeeper,
		app.GetSubspace(servicetypes.ModuleName), app.ModuleAccountAddrs(), opbtypes.PointTokenFeeCollectorName,
	)

	app.OracleKeeper = oraclekeeper.NewKeeper(
		appCodec, keys[oracletypes.StoreKey], app.GetSubspace(oracletypes.ModuleName),
		app.ServiceKeeper,
	)

	app.RandomKeeper = randomkeeper.NewKeeper(appCodec, keys[randomtypes.StoreKey], app.BankKeeper, app.ServiceKeeper)

	app.NodeKeeper = *app.NodeKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.SlashingKeeper.Hooks()),
	)

	PermKeeper := permkeeper.NewKeeper(appCodec, keys[permtypes.StoreKey])
	app.PermKeeper = PermKeeper

	app.IdentityKeeper = identitykeeper.NewKeeper(appCodec, keys[identitytypes.StoreKey])

	app.OpbKeeper = opbkeeper.NewKeeper(
		appCodec, keys[opbtypes.StoreKey], app.AccountKeeper,
		app.BankKeeper, app.TokenKeeper, app.PermKeeper,
		app.GetSubspace(opbtypes.ModuleName),
	)
	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	scopedTIBCKeeper := app.CapabilityKeeper.ScopeToModule(tibchost.ModuleName)
	scopedTIBCMockKeeper := app.CapabilityKeeper.ScopeToModule(tibcmock.ModuleName)
	// Create Transfer Keepers
	app.NftTransferKeeper = tibcnfttransferkeeper.NewKeeper(
		appCodec, keys[tibcnfttypes.StoreKey], app.GetSubspace(tibcnfttypes.ModuleName),
		app.AccountKeeper, app.NFTKeeper,
		app.TIBCKeeper.PacketKeeper, app.TIBCKeeper.ClientKeeper,
	)
	nfttransferModule := tibcnfttransfer.NewAppModule(app.NftTransferKeeper)
	// Create TIBC Keeper
	app.TIBCKeeper = tibckeeper.NewKeeper(
		appCodec, keys[tibchost.StoreKey], app.GetSubspace(tibchost.ModuleName), stakingkeeper.Keeper{}, scopedTIBCKeeper,
	)
	tibcmockModule := tibcmock.NewAppModule()
	tibcRouter := tibcroutingtypes.NewRouter()
	tibcRouter.AddRoute(tibcnfttypes.ModuleName, nfttransferModule)
	tibcRouter.AddRoute(tibcmock.ModuleName, tibcmockModule)
	app.TIBCKeeper.SetRouter(tibcRouter)

	wasmDir := filepath.Join(homePath, "wasm")

	supportedFeatures := "stargate"
	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		stakingkeeper.Keeper{},
		distrkeeper.Keeper{},
		nil,
		nil,
		nil,
		nil,
		bApp.Router(),
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		wasmDir,
		wasm.DefaultWasmConfig(),
		supportedFeatures,
	)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.AccountKeeper, app.NodeKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, app.NodeKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		nft.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper),
		opb.NewAppModule(appCodec, app.OpbKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.NodeKeeper),
		tibc.NewAppModule(app.TIBCKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName, slashingtypes.ModuleName, evidencetypes.ModuleName,
		nodetypes.ModuleName, recordtypes.ModuleName, tokentypes.ModuleName,
		nfttypes.ModuleName, servicetypes.ModuleName, randomtypes.ModuleName,
		wasm.ModuleName, tibchost.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		node.ModuleName,
		servicetypes.ModuleName,
		wasm.ModuleName,
		tibchost.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		permtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		recordtypes.ModuleName,
		tokentypes.ModuleName,
		nfttypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		identitytypes.ModuleName,
		wasm.ModuleName,
		opb.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.SlashingKeeper, app.NodeKeeper), app.AccountKeeper, app.BankKeeper, app.NodeKeeper),
		params.NewAppModule(app.ParamsKeeper),
		cparams.NewAppModule(appCodec, app.ParamsKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		nft.NewAppModule(appCodec, app.NFTKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper),
		opb.NewAppModule(appCodec, app.OpbKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.NodeKeeper),
		tibc.NewAppModule(app.TIBCKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	anteHandler, err := ante.NewAnteHandler(
		ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:  app.FeeGrantKeeper,
			SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
		},
	)
	if err != nil {
		panic(err)
	}
	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	// set peer filter by node ID
	app.SetIDPeerFilter(app.NodeKeeper.FilterNodeByID)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		//ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		//app.capabilityKeeper.InitializeAndSeal(ctx)
	}
	app.ScopedTIBCKeeper = scopedTIBCKeeper
	app.ScopedTIBCMockKeeper = scopedTIBCMockKeeper
	return app
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// SimApp. It is useful for tests and clients who do not want to construct the
// full SimApp
func MakeCodecs() (codec.Codec, *codec.LegacyAmino) {
	config := MakeTestEncodingConfig()
	return config.Marshaler, config.Amino
}

// Name returns the name of the App
func (app *SimApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *SimApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *SimApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	// add system service at InitChainer, overwrite if it exists
	var serviceGenState servicetypes.GenesisState
	app.appCodec.MustUnmarshalJSON(genesisState[servicetypes.ModuleName], &serviceGenState)
	serviceGenState.Definitions = append(serviceGenState.Definitions, servicetypes.GenOraclePriceSvcDefinition())
	serviceGenState.Bindings = append(serviceGenState.Bindings, servicetypes.GenOraclePriceSvcBinding(tokentypes.GetNativeToken().MinUnit))
	serviceGenState.Definitions = append(serviceGenState.Definitions, randomtypes.GetSvcDefinition())
	genesisState[servicetypes.ModuleName] = app.appCodec.MustMarshalJSON(&serviceGenState)

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *SimApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns SimApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns SimApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *SimApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns SimApp's InterfaceRegistry
func (app *SimApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *SimApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *SimApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *SimApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *SimApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(apiSvr.ClientCtx, apiSvr.GRPCGatewayRouter)

	if apiConfig.Swagger {
		lite.RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *SimApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	ParamsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	ParamsKeeper.Subspace(authtypes.ModuleName)
	ParamsKeeper.Subspace(banktypes.ModuleName)
	ParamsKeeper.Subspace(nodetypes.ModuleName)
	ParamsKeeper.Subspace(slashingtypes.ModuleName)
	ParamsKeeper.Subspace(crisistypes.ModuleName)
	ParamsKeeper.Subspace(tokentypes.ModuleName)
	ParamsKeeper.Subspace(recordtypes.ModuleName)
	ParamsKeeper.Subspace(servicetypes.ModuleName)
	ParamsKeeper.Subspace(opbtypes.ModuleName)
	ParamsKeeper.Subspace(wasm.ModuleName)
	ParamsKeeper.Subspace(tibchost.ModuleName)

	return ParamsKeeper
}
