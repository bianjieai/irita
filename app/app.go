package app

import (
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/irisnet/irismod/modules/mt"

	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/bianjieai/irita/modules/evm/crypto"
	appevmtypes "github.com/bianjieai/irita/modules/evm/types"
	evmutils "github.com/bianjieai/irita/modules/evm/utils"
	tibcclienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"

	wservicekeeper "github.com/bianjieai/irita/modules/wservice/keeper"
	wservicetypes "github.com/bianjieai/irita/modules/wservice/types"

	"github.com/spf13/cast"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"

	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
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
	sdkupgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	mtkeeper "github.com/irisnet/irismod/modules/mt/keeper"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
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

	"github.com/bianjieai/irita/address"
	"github.com/bianjieai/irita/lite"
	appkeeper "github.com/bianjieai/irita/modules/evm"
	"github.com/bianjieai/irita/modules/opb"
	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	opbtypes "github.com/bianjieai/irita/modules/opb/types"

	tibc "github.com/bianjieai/irita/modules/tibc"
	tibckeeper "github.com/bianjieai/irita/modules/tibc/keeper"

	tibcnfttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibcnfttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/keeper"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	tibccorekeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"

	"github.com/tharsis/ethermint/app/ante"
	srvflags "github.com/tharsis/ethermint/server/flags"
	ethermint "github.com/tharsis/ethermint/types"
	"github.com/tharsis/ethermint/x/evm"
	evmrest "github.com/tharsis/ethermint/x/evm/client/rest"
	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	"github.com/tharsis/ethermint/x/feemarket"
	feemarketkeeper "github.com/tharsis/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"
)

const appName = "IritaApp"

// DefaultNodeHome default home directories for the application daemon
var DefaultNodeHome string
var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
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
		mt.AppModuleBasic{},
		service.AppModuleBasic{},
		oracle.AppModuleBasic{},
		random.AppModuleBasic{},
		perm.AppModuleBasic{},
		identity.AppModuleBasic{},
		node.AppModuleBasic{},
		opb.AppModuleBasic{},
		tibc.AppModule{},
		tibcnfttransfer.AppModuleBasic{},
		wasm.AppModuleBasic{},

		// evm
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
	)
	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:          nil,
		tokentypes.ModuleName:               {authtypes.Minter, authtypes.Burner},
		servicetypes.DepositAccName:         nil,
		servicetypes.RequestAccName:         nil,
		opbtypes.PointTokenFeeCollectorName: nil,
		tibcnfttypes.ModuleName:             nil,

		// evm
		evmtypes.ModuleName: {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
	}
	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{}
)

// Verify app interface at compile time
var _ simapp.App = (*IritaApp)(nil)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".irita")

	address.ConfigureBech32Prefix()
	tokentypes.SetNativeToken(
		"irita",
		"Irita base native token",
		"uirita",
		6,
		1000000000,
		math.MaxUint64,
		true,
		sdk.AccAddress{},
	)
}

// IritaApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type IritaApp struct {
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
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	slashingKeeper   slashingkeeper.Keeper
	crisisKeeper     crisiskeeper.Keeper
	upgradeKeeper    upgradekeeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	evidenceKeeper   evidencekeeper.Keeper
	recordKeeper     recordkeeper.Keeper
	tokenKeeper      tokenkeeper.Keeper
	nftKeeper        nftkeeper.Keeper
	mtKeeper         mtkeeper.Keeper
	serviceKeeper    servicekeeper.Keeper
	oracleKeeper     oraclekeeper.Keeper
	randomKeeper     randomkeeper.Keeper
	permKeeper       permkeeper.Keeper
	identityKeeper   identitykeeper.Keeper
	nodeKeeper       nodekeeper.Keeper
	opbKeeper        opbkeeper.Keeper
	wservicekeeper   wservicekeeper.IKeeper
	feeGrantKeeper   feegrantkeeper.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	wasmKeeper       wasm.Keeper
	// tibc
	scopedTIBCKeeper     capabilitykeeper.ScopedKeeper
	scopedTIBCMockKeeper capabilitykeeper.ScopedKeeper
	tibcKeeper           *tibckeeper.Keeper
	nftTransferKeeper    tibcnfttransferkeeper.Keeper

	// Ethermint keepers
	EvmKeeper       *evmkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

// NewIritaApp returns a reference to an initialized IritaApp.
func NewIritaApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig simappparams.EncodingConfig, appOpts servertypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *IritaApp {
	// TODO: Remove cdc in favor of appCodec once all modules are migrated.

	evmutils.SetEthermintSupportedAlgorithms()

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
		mttypes.StoreKey,
		servicetypes.StoreKey,
		oracletypes.StoreKey,
		randomtypes.StoreKey,
		permtypes.StoreKey,
		identitytypes.StoreKey,
		nodetypes.StoreKey,
		opbtypes.StoreKey,
		tibchost.StoreKey,
		tibcnfttypes.StoreKey,
		wasm.StoreKey,

		// evm
		evmtypes.StoreKey, feemarkettypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &IritaApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.paramsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add keepers
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.accountKeeper, app.GetSubspace(banktypes.ModuleName), app.ModuleAccountAddrs(),
	)
	app.nodeKeeper = node.NewKeeper(appCodec, keys[nodetypes.StoreKey], app.GetSubspace(node.ModuleName))
	app.slashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &app.nodeKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.crisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.bankKeeper, authtypes.FeeCollectorName,
	)
	app.feeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.accountKeeper)

	sdkUpgradeKeeper := sdkupgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, app.BaseApp)
	app.upgradeKeeper = upgradekeeper.NewKeeper(sdkUpgradeKeeper)

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.nodeKeeper, app.slashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.evidenceKeeper = *evidenceKeeper

	app.tokenKeeper = tokenkeeper.NewKeeper(
		appCodec, keys[tokentypes.StoreKey], app.GetSubspace(tokentypes.ModuleName),
		app.bankKeeper, app.ModuleAccountAddrs(), opbtypes.PointTokenFeeCollectorName,
	)

	app.recordKeeper = recordkeeper.NewKeeper(appCodec, keys[recordtypes.StoreKey])
	app.nftKeeper = nftkeeper.NewKeeper(appCodec, keys[nfttypes.StoreKey])
	app.mtKeeper = mtkeeper.NewKeeper(appCodec, keys[mttypes.StoreKey])

	app.serviceKeeper = servicekeeper.NewKeeper(
		appCodec, keys[servicetypes.StoreKey], app.accountKeeper, app.bankKeeper,
		app.GetSubspace(servicetypes.ModuleName), app.ModuleAccountAddrs(), opbtypes.PointTokenFeeCollectorName,
	)

	app.oracleKeeper = oraclekeeper.NewKeeper(
		appCodec, keys[oracletypes.StoreKey], app.GetSubspace(oracletypes.ModuleName),
		app.serviceKeeper,
	)

	app.randomKeeper = randomkeeper.NewKeeper(appCodec, keys[randomtypes.StoreKey], app.bankKeeper, app.serviceKeeper)

	app.nodeKeeper = *app.nodeKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.slashingKeeper.Hooks()),
	)

	permKeeper := permkeeper.NewKeeper(appCodec, keys[permtypes.StoreKey])
	app.permKeeper = RegisterAccessControl(permKeeper)

	app.identityKeeper = identitykeeper.NewKeeper(appCodec, keys[identitytypes.StoreKey])

	// evm
	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))

	// Create Ethermint  keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec, keys[feemarkettypes.StoreKey], app.GetSubspace(feemarkettypes.ModuleName),
	)

	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec, keys[evmtypes.StoreKey], tkeys[evmtypes.TransientKey], app.GetSubspace(evmtypes.ModuleName),
		app.accountKeeper, app.bankKeeper, appkeeper.WNodeKeeper{Keeper: app.nodeKeeper}, app.FeeMarketKeeper,
		tracer, // debug EVM based on Baseapp options
	)

	app.opbKeeper = opbkeeper.NewKeeper(
		appCodec, keys[opbtypes.StoreKey], app.accountKeeper,
		app.bankKeeper, app.tokenKeeper, app.permKeeper,
		app.GetSubspace(opbtypes.ModuleName),
	)

	ethOpbV := appkeeper.NewEthOpbValidator(&app.opbKeeper, &app.tokenKeeper, app.EvmKeeper, logger)
	app.EvmKeeper.TransferFunc = ethOpbV.Transfer
	app.EvmKeeper.AccStoreKey = keys[authtypes.StoreKey]

	// register the proposal types
	tibccorekeeper := tibccorekeeper.NewKeeper(
		appCodec, keys[tibchost.StoreKey], app.GetSubspace(tibchost.ModuleName), stakingkeeper.Keeper{},
	)
	app.tibcKeeper = tibckeeper.NewKeeper(tibccorekeeper)
	app.nftTransferKeeper = tibcnfttransferkeeper.NewKeeper(
		appCodec, keys[tibcnfttypes.StoreKey], app.GetSubspace(tibcnfttypes.ModuleName),
		app.accountKeeper, tibckeeper.WrapNftKeeper(app.nftKeeper),
		app.tibcKeeper.PacketKeeper, app.tibcKeeper.ClientKeeper,
	)
	nfttransferModule := tibcnfttransfer.NewAppModule(app.nftTransferKeeper)
	tibcRouter := tibcroutingtypes.NewRouter()
	tibcRouter.AddRoute(tibcnfttypes.ModuleName, nfttransferModule)
	app.tibcKeeper.SetRouter(tibcRouter)

	app.wservicekeeper = wservicekeeper.NewKeeper(appCodec, keys[wservicetypes.StoreKey], app.serviceKeeper)

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	supportedFeatures := "stargate"
	app.wasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
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
		wasmConfig,
		supportedFeatures,
	)

	/****  Module Options ****/
	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)

	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.nodeKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper, skipGenesisInvariants),
		feegrantmodule.NewAppModule(appCodec, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.slashingKeeper, app.nodeKeeper), app.accountKeeper, app.bankKeeper, app.nodeKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		params.NewAppModule(app.paramsKeeper),
		cparams.NewAppModule(appCodec, app.paramsKeeper),
		token.NewAppModule(appCodec, app.tokenKeeper, app.accountKeeper, app.bankKeeper),
		nft.NewAppModule(appCodec, app.nftKeeper, app.accountKeeper, app.bankKeeper),
		mt.NewAppModule(appCodec, app.mtKeeper, app.accountKeeper, app.bankKeeper),
		service.NewAppModule(appCodec, app.serviceKeeper, app.accountKeeper, app.bankKeeper),
		oracle.NewAppModule(appCodec, app.oracleKeeper, app.accountKeeper, app.bankKeeper),
		random.NewAppModule(appCodec, app.randomKeeper, app.accountKeeper, app.bankKeeper),
		perm.NewAppModule(appCodec, app.permKeeper),
		identity.NewAppModule(app.identityKeeper),
		record.NewAppModule(appCodec, app.recordKeeper, app.accountKeeper, app.bankKeeper),
		node.NewAppModule(appCodec, app.nodeKeeper),
		opb.NewAppModule(appCodec, app.opbKeeper),
		tibc.NewAppModule(app.tibcKeeper),
		wasm.NewAppModule(appCodec, &app.wasmKeeper, app.nodeKeeper),
		nfttransferModule,

		// evm
		evm.NewAppModule(app.EvmKeeper, app.accountKeeper),
		feemarket.NewAppModule(app.FeeMarketKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradetypes.ModuleName, slashingtypes.ModuleName, evidencetypes.ModuleName,
		nodetypes.ModuleName, recordtypes.ModuleName, tokentypes.ModuleName,
		nfttypes.ModuleName, mttypes.ModuleName, servicetypes.ModuleName, randomtypes.ModuleName,
		tibchost.ModuleName, evmtypes.ModuleName, wasm.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		nodetypes.ModuleName,
		servicetypes.ModuleName,
		tibchost.ModuleName,
		wasm.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		permtypes.ModuleName,
		authtypes.ModuleName,
		nodetypes.ModuleName,
		banktypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		evidencetypes.ModuleName,
		recordtypes.ModuleName,
		tokentypes.ModuleName,
		nfttypes.ModuleName,
		mttypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		identitytypes.ModuleName,
		opb.ModuleName,
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
		wasm.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		feegrantmodule.NewAppModule(appCodec, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.slashingKeeper, app.nodeKeeper), app.accountKeeper, app.bankKeeper, app.nodeKeeper),
		params.NewAppModule(app.paramsKeeper),
		cparams.NewAppModule(appCodec, app.paramsKeeper),
		record.NewAppModule(appCodec, app.recordKeeper, app.accountKeeper, app.bankKeeper),
		token.NewAppModule(appCodec, app.tokenKeeper, app.accountKeeper, app.bankKeeper),
		nft.NewAppModule(appCodec, app.nftKeeper, app.accountKeeper, app.bankKeeper),
		mt.NewAppModule(appCodec, app.mtKeeper, app.accountKeeper, app.bankKeeper),
		service.NewAppModule(appCodec, app.serviceKeeper, app.accountKeeper, app.bankKeeper),
		oracle.NewAppModule(appCodec, app.oracleKeeper, app.accountKeeper, app.bankKeeper),
		random.NewAppModule(appCodec, app.randomKeeper, app.accountKeeper, app.bankKeeper),
		perm.NewAppModule(appCodec, app.permKeeper),
		identity.NewAppModule(app.identityKeeper),
		node.NewAppModule(appCodec, app.nodeKeeper),
		opb.NewAppModule(appCodec, app.opbKeeper),
		tibc.NewAppModule(app.tibcKeeper),
		wasm.NewAppModule(appCodec, &app.wasmKeeper, app.nodeKeeper),
		nfttransferModule,

		// evm
		evm.NewAppModule(app.EvmKeeper, app.accountKeeper),
		feemarket.NewAppModule(app.FeeMarketKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	anteHandler := NewAnteHandler(
		HandlerOptions{
			permKeeper:      app.permKeeper,
			accountKeeper:   app.accountKeeper,
			bankKeeper:      app.bankKeeper,
			tokenKeeper:     app.tokenKeeper,
			opbKeeper:       app.opbKeeper,
			signModeHandler: encodingConfig.TxConfig.SignModeHandler(),
			feegrantKeeper:  app.feeGrantKeeper,
			wserviceKeeper:  app.wservicekeeper,
			sigGasConsumer:  ante.DefaultSigVerificationGasConsumer,

			// evm
			evmFeeMarketKeeper: app.FeeMarketKeeper,
			evmKeeper:          app.EvmKeeper,
		},
	)
	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	// Set software upgrade execution logic
	// app.RegisterUpgradePlan("add-record-module",
	// 	store.StoreUpgrades{
	// 		Added: []string{recordtypes.StoreKey},
	// 	},
	// 	func(ctx sdk.Context, plan sdkupgrade.Plan) {},
	// )

	app.RegisterUpgradePlan(
		"v2.2-wenchangchain", store.StoreUpgrades{
			Added: []string{feegrant.StoreKey, tibchost.StoreKey, tibcnfttypes.StoreKey},
		},
		func(ctx sdk.Context, plan sdkupgrade.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			tibcclienttypes.SetDefaultGenesisState(tibcclienttypes.GenesisState{
				NativeChainName: "wenchangchain-mainnet",
			})
			fromVM[authtypes.ModuleName] = auth.AppModule{}.ConsensusVersion()
			fromVM[banktypes.ModuleName] = 1
			fromVM[stakingtypes.ModuleName] = 1
			fromVM[opbtypes.ModuleName] = 1
			fromVM[wservicetypes.ModuleName] = 1
			fromVM[identitytypes.ModuleName] = 1
			fromVM[cslashing.ModuleName] = cslashing.AppModule{}.ConsensusVersion()
			fromVM[capabilitytypes.ModuleName] = capability.AppModule{}.ConsensusVersion()
			fromVM[nodetypes.ModuleName] = node.AppModule{}.ConsensusVersion()
			fromVM[genutiltypes.ModuleName] = genutil.AppModule{}.ConsensusVersion()
			fromVM[paramstypes.ModuleName] = cparams.AppModule{}.ConsensusVersion()
			fromVM[crisistypes.ModuleName] = crisis.AppModule{}.ConsensusVersion()
			fromVM[upgradetypes.ModuleName] = crisis.AppModule{}.ConsensusVersion()
			fromVM[evidencetypes.ModuleName] = evidence.AppModule{}.ConsensusVersion()
			fromVM[feegrant.ModuleName] = feegrantmodule.AppModule{}.ConsensusVersion()
			fromVM[tokentypes.ModuleName] = token.AppModule{}.ConsensusVersion()
			fromVM[recordtypes.ModuleName] = record.AppModule{}.ConsensusVersion()
			fromVM[nfttypes.ModuleName] = nft.AppModule{}.ConsensusVersion()
			fromVM[servicetypes.ModuleName] = service.AppModule{}.ConsensusVersion()
			fromVM[oracletypes.ModuleName] = oracle.AppModule{}.ConsensusVersion()
			fromVM[randomtypes.ModuleName] = random.AppModule{}.ConsensusVersion()
			fromVM[permtypes.ModuleName] = perm.AppModule{}.ConsensusVersion()
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.RegisterUpgradePlan(
		"v3.0.0-wenchangchain", store.StoreUpgrades{
			Added: []string{evmtypes.StoreKey, feemarkettypes.StoreKey},
		},
		func(ctx sdk.Context, plan sdkupgrade.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			newParams := evmtypes.NewParams(appevmtypes.DefaultEvmDenom, true, true, evmtypes.DefaultChainConfig())
			evmtypes.SetDefaultGenesisState(newParams, []evmtypes.GenesisAccount{})
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	app.RegisterUpgradePlan(
		"v3.1.0-wenchangchain", store.StoreUpgrades{
			Added: []string{mttypes.StoreKey},
		},
		func(ctx sdk.Context, plan sdkupgrade.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			opbParams := app.opbKeeper.GetParams(ctx)
			gasOwner, err := sdk.AccAddressFromBech32(opbParams.BaseTokenManager)
			if err != nil {
				return nil, err
			}
			err = app.tokenKeeper.IssueToken(
				ctx,
				"gas",
				"IRITA Fee Token",
				"ugas",
				18,
				1000000000,
				math.MaxUint64,
				true,
				gasOwner,
			)
			if err != nil {
				return nil, err
			}
			evmParams := app.EvmKeeper.GetParams(ctx)
			evmParams.EvmDenom = "ugas"
			app.EvmKeeper.SetParams(ctx, evmParams)
			fMtParams := app.FeeMarketKeeper.GetParams(ctx)
			fMtParams.NoBaseFee = true
			app.FeeMarketKeeper.SetParams(ctx, fMtParams)
			return app.mm.RunMigrations(ctx, app.configurator, fromVM)
		},
	)

	// set peer filter by node ID
	app.SetIDPeerFilter(app.nodeKeeper.FilterNodeByID)

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
	return app
}

// Name returns the name of the App
func (app *IritaApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *IritaApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	chainID, _ := ethermint.ParseChainID(req.GetHeader().ChainID)
	if app.EvmKeeper.Signer == nil {
		app.EvmKeeper.Signer = crypto.NewSm2Signer(chainID)
	}
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *IritaApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *IritaApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	// add system service at InitChainer, overwrite if it exists
	var serviceGenState servicetypes.GenesisState
	app.appCodec.MustUnmarshalJSON(genesisState[servicetypes.ModuleName], &serviceGenState)
	//req.ChainId

	chainID, _ := ethermint.ParseChainID(req.ChainId)
	app.EvmKeeper.Signer = crypto.NewSm2Signer(chainID)
	serviceGenState.Definitions = append(serviceGenState.Definitions, servicetypes.GenOraclePriceSvcDefinition())
	serviceGenState.Bindings = append(serviceGenState.Bindings, servicetypes.GenOraclePriceSvcBinding(tokentypes.GetNativeToken().MinUnit))
	serviceGenState.Definitions = append(serviceGenState.Definitions, randomtypes.GetSvcDefinition())
	genesisState[servicetypes.ModuleName] = app.appCodec.MustMarshalJSON(&serviceGenState)

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *IritaApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *IritaApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *IritaApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns IritaApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *IritaApp) LegacyAmino() *codec.LegacyAmino {
	return app.cdc
}

// AppCodec returns IritaApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *IritaApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns IritaApp's InterfaceRegistry
func (app *IritaApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IritaApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IritaApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *IritaApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IritaApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.paramsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *IritaApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *IritaApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)

	// evm
	evmrest.RegisterTxRoutes(clientCtx, apiSvr.Router)

	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	if apiConfig.Swagger {
		lite.RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *IritaApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterUpgradePlan implements the upgrade execution logic of the upgrade module
func (app *IritaApp) RegisterUpgradePlan(planName string,
	upgrades store.StoreUpgrades, upgradeHandler sdkupgrade.UpgradeHandler) {
	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		app.Logger().Info("not found upgrade plan", "planName", planName, "err", err.Error())
		return
	}

	if upgradeInfo.Name == planName && !app.upgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// this configures a no-op upgrade handler for the planName upgrade
		app.upgradeKeeper.SetUpgradeHandler(planName, upgradeHandler)
		// configure store loader that checks if version+1 == upgradeHeight and applies store upgrades
		app.SetStoreLoader(sdkupgrade.UpgradeStoreLoader(upgradeInfo.Height, &upgrades))
	}
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
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(nodetypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(tokentypes.ModuleName)
	paramsKeeper.Subspace(recordtypes.ModuleName)
	paramsKeeper.Subspace(servicetypes.ModuleName)
	paramsKeeper.Subspace(opbtypes.ModuleName)
	paramsKeeper.Subspace(tibchost.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)

	// evm
	paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(feemarkettypes.ModuleName)

	return paramsKeeper
}
