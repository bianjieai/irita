package app

import (
	"io"
	"math"
	"os"
	"path/filepath"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/bianjieai/irita/address"
	appante "github.com/bianjieai/irita/app/ante"
	"github.com/bianjieai/irita/crypto/hd"
	"github.com/bianjieai/irita/lite"
	appevm "github.com/bianjieai/irita/modules/evm"
	tibc "github.com/bianjieai/irita/modules/tibc"
	tibckeeper "github.com/bianjieai/irita/modules/tibc/keeper"
	"github.com/bianjieai/irita/wrapper"
	tibcmttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer"
	tibcmttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/keeper"
	tibcmttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/types"
	tibcnfttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibcnfttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/keeper"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	tibccorekeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	grpcnode "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	store "github.com/cosmos/cosmos-sdk/store/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
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
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	sdkupgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ethermintante "github.com/evmos/ethermint/app/ante"
	srvflags "github.com/evmos/ethermint/server/flags"
	"github.com/evmos/ethermint/x/evm"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/evmos/ethermint/x/evm/vm/geth"
	"github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	"github.com/spf13/cast"
	"iritamod.bianjie.ai/modules/genutil"
	genutiltypes "iritamod.bianjie.ai/modules/genutil"
	"iritamod.bianjie.ai/modules/identity"
	identitykeeper "iritamod.bianjie.ai/modules/identity/keeper"
	identitytypes "iritamod.bianjie.ai/modules/identity/types"
	"iritamod.bianjie.ai/modules/node"
	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	nodetypes "iritamod.bianjie.ai/modules/node/types"
	cparams "iritamod.bianjie.ai/modules/params"
	cslashing "iritamod.bianjie.ai/modules/slashing"
	"iritamod.bianjie.ai/modules/upgrade"
	upgradekeeper "iritamod.bianjie.ai/modules/upgrade/keeper"
	upgradetypes "iritamod.bianjie.ai/modules/upgrade/types"
	"mods.irisnet.org/modules/mt"
	mtkeeper "mods.irisnet.org/modules/mt/keeper"
	mttypes "mods.irisnet.org/modules/mt/types"
	"mods.irisnet.org/modules/nft"
	nftkeeper "mods.irisnet.org/modules/nft/keeper"
	nfttypes "mods.irisnet.org/modules/nft/types"
	"mods.irisnet.org/modules/oracle"
	oraclekeeper "mods.irisnet.org/modules/oracle/keeper"
	oracletypes "mods.irisnet.org/modules/oracle/types"
	"mods.irisnet.org/modules/random"
	randomkeeper "mods.irisnet.org/modules/random/keeper"
	randomtypes "mods.irisnet.org/modules/random/types"
	"mods.irisnet.org/modules/record"
	recordkeeper "mods.irisnet.org/modules/record/keeper"
	recordtypes "mods.irisnet.org/modules/record/types"
	"mods.irisnet.org/modules/service"
	servicekeeper "mods.irisnet.org/modules/service/keeper"
	servicetypes "mods.irisnet.org/modules/service/types"
	"mods.irisnet.org/modules/token"
	tokenkeeper "mods.irisnet.org/modules/token/keeper"
	tokentypes "mods.irisnet.org/modules/token/types"
	tokentypesv1 "mods.irisnet.org/modules/token/types/v1"
)

const (
	appName = "IritaApp"
)

var storeKeys = []string{
	authtypes.StoreKey,
	banktypes.StoreKey,
	slashingtypes.StoreKey,
	crisistypes.StoreKey,
	paramstypes.StoreKey,
	consensustypes.StoreKey,
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
	identitytypes.StoreKey,
	nodetypes.StoreKey,
	tibchost.StoreKey,
	tibcnfttypes.StoreKey,
	tibcmttypes.StoreKey,

	// evm
	evmtypes.StoreKey, feemarkettypes.StoreKey,
}

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
		consensus.AppModuleBasic{},
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
		identity.AppModuleBasic{},
		node.AppModuleBasic{},
		tibc.AppModule{},
		tibcnfttransfer.AppModuleBasic{},
		tibcmttransfer.AppModuleBasic{},

		// evm
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
	)
	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:  nil,
		tokentypes.ModuleName:       {authtypes.Minter, authtypes.Burner},
		servicetypes.DepositAccName: nil,
		servicetypes.RequestAccName: nil,
		tibcnfttypes.ModuleName:     nil,
		tibcmttypes.ModuleName:      nil,
		nfttypes.ModuleName:         nil,

		// evm
		evmtypes.ModuleName: {authtypes.Minter, authtypes.Burner}, // used for secure addition and subtraction of balance using module account
	}
	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{}
	// app options
	appOptions = IritaAppOptions{}
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".irita")

	address.ConfigureBech32Prefix()
	tokentypesv1.SetNativeToken(
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

var _ servertypes.Application = (*IritaApp)(nil)

// IritaApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type IritaApp struct {
	*baseapp.BaseApp
	cdc               *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	// keys to access the substores
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	accountKeeper         authkeeper.AccountKeeper
	bankKeeper            bankkeeper.Keeper
	slashingKeeper        slashingkeeper.Keeper
	crisisKeeper          *crisiskeeper.Keeper
	upgradeKeeper         upgradekeeper.Keeper
	paramsKeeper          paramskeeper.Keeper
	evidenceKeeper        evidencekeeper.Keeper
	recordKeeper          recordkeeper.Keeper
	tokenKeeper           tokenkeeper.Keeper
	nftKeeper             nftkeeper.Keeper
	mtKeeper              mtkeeper.Keeper
	serviceKeeper         servicekeeper.Keeper
	oracleKeeper          oraclekeeper.Keeper
	randomKeeper          randomkeeper.Keeper
	identityKeeper        identitykeeper.Keeper
	nodeKeeper            *nodekeeper.Keeper
	feeGrantKeeper        feegrantkeeper.Keeper
	capabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	// tibc
	scopedTIBCKeeper     capabilitykeeper.ScopedKeeper
	scopedTIBCMockKeeper capabilitykeeper.ScopedKeeper
	tibcKeeper           *tibckeeper.Keeper
	nftTransferKeeper    tibcnfttransferkeeper.Keeper
	mtTransferKeeper     tibcmttransferkeeper.Keeper

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
	logger log.Logger, 
	db dbm.DB, 
	traceStore io.Writer, 
	loadLatest bool, 
	encodingConfig simappparams.EncodingConfig, 
	appOpts servertypes.AppOptions, 
	baseAppOptions ...func(*baseapp.BaseApp),
) *IritaApp {
	// TODO: Remove cdc in favor of appCodec once all modules are migrated.

	hd.SetSupportedAlgorithms()

	appCodec := encodingConfig.Codec
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(storeKeys...)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey, evmtypes.TransientKey, feemarkettypes.TransientKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &IritaApp{
		BaseApp:           bApp,
		cdc:               cdc,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	authority := cast.ToString(appOpts.Get(flagAuthority))
	authorityAddr := sdk.MustAccAddressFromBech32(authority)

	app.paramsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])
	app.ConsensusParamsKeeper = consensuskeeper.NewKeeper(
		appCodec,
		app.keys[consensustypes.StoreKey],
		authority,
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add keepers
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.Bech32PrefixAccAddr,
		authority,
	)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.accountKeeper,
		app.ModuleAccountAddrs(),
		authority,
	)
	app.nodeKeeper = node.NewKeeper(appCodec, keys[nodetypes.StoreKey], app.GetSubspace(node.ModuleName))

	stakingKeeper := wrapper.NewStakingKeeper(app.nodeKeeper)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		cdc,
		keys[slashingtypes.StoreKey],
		stakingKeeper,
		authority,
	)
	app.crisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		app.bankKeeper,
		authtypes.FeeCollectorName,
		authority,
	)
	app.feeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.accountKeeper)

	// get skipUpgradeHeights from the app options
	skipUpgradeHeights := map[int64]bool{}
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}
	app.upgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		app.BaseApp,
		authority,
	)

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		stakingKeeper,
		app.slashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.evidenceKeeper = *evidenceKeeper

	app.tokenKeeper = tokenkeeper.NewKeeper(
		appCodec,
		keys[tokentypes.StoreKey],
		app.bankKeeper,
		app.accountKeeper,
		wrapper.NewEVMKeeper(app.EvmKeeper),
		nil,
		authtypes.FeeCollectorName,
		authority,
	)

	app.recordKeeper = recordkeeper.NewKeeper(appCodec, keys[recordtypes.StoreKey])
	app.nftKeeper = nftkeeper.NewKeeper(
		appCodec,
		keys[nfttypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
	)
	app.mtKeeper = mtkeeper.NewKeeper(appCodec, keys[mttypes.StoreKey])

	app.serviceKeeper = servicekeeper.NewKeeper(
		appCodec,
		keys[servicetypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		servicetypes.FeeCollectorName,
		authority,
	)

	app.oracleKeeper = oraclekeeper.NewKeeper(
		appCodec,
		keys[oracletypes.StoreKey],
		app.serviceKeeper,
	)

	app.randomKeeper = randomkeeper.NewKeeper(appCodec, keys[randomtypes.StoreKey], app.bankKeeper, app.serviceKeeper)

	app.nodeKeeper = app.nodeKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.slashingKeeper.Hooks()),
	)

	app.identityKeeper = identitykeeper.NewKeeper(appCodec, keys[identitytypes.StoreKey])

	// evm
	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))

	// Create Ethermint  keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec,
		authorityAddr,
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.TransientKey],
		app.GetSubspace(feemarkettypes.ModuleName),
	)
	app.EvmKeeper = evmkeeper.NewKeeper(
		appCodec,
		keys[evmtypes.StoreKey],
		tkeys[evmtypes.TransientKey],
		authorityAddr,
		app.accountKeeper,
		app.bankKeeper,
		appevm.NewStakingKeeper(*app.nodeKeeper),
		app.FeeMarketKeeper,
		nil,
		geth.NewEVM,
		tracer, // debug EVM based on Baseapp options
		app.GetSubspace(evmtypes.ModuleName),
	)

	// register the proposal types
	tibccorekeeper := tibccorekeeper.NewKeeper(
		appCodec,
		keys[tibchost.StoreKey],
		app.nodeKeeper,
		authority,
	)
	app.tibcKeeper = tibckeeper.NewKeeper(tibccorekeeper)
	app.nftTransferKeeper = tibcnfttransferkeeper.NewKeeper(
		appCodec, keys[tibcnfttypes.StoreKey],
		app.accountKeeper, nftkeeper.NewLegacyKeeper(app.nftKeeper),
		app.tibcKeeper.PacketKeeper, app.tibcKeeper.ClientKeeper,
	)
	app.mtTransferKeeper = tibcmttransferkeeper.NewKeeper(
		appCodec, keys[tibcmttypes.StoreKey],
		app.accountKeeper, app.mtKeeper,
		app.tibcKeeper.PacketKeeper, app.tibcKeeper.ClientKeeper,
	)
	nfttransferModule := tibcnfttransfer.NewAppModule(app.nftTransferKeeper)
	mttransferModule := tibcmttransfer.NewAppModule(app.mtTransferKeeper)
	tibcRouter := tibcroutingtypes.NewRouter()
	tibcRouter.AddRoute(tibcnfttypes.ModuleName, nfttransferModule)
	tibcRouter.AddRoute(tibcmttypes.ModuleName, mttransferModule)
	app.tibcKeeper.SetRouter(tibcRouter)

	/****  Module Options ****/
	skipGenesisInvariants := false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)

	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.nodeKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper, app.GetSubspace(banktypes.ModuleName)),
		crisis.NewAppModule(app.crisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.slashingKeeper, app.nodeKeeper), app.accountKeeper, app.bankKeeper, stakingKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		params.NewAppModule(app.paramsKeeper),
		cparams.NewAppModule(appCodec, app.paramsKeeper),
		token.NewAppModule(appCodec, app.tokenKeeper, app.accountKeeper, app.bankKeeper, app.GetSubspace(tokentypes.ModuleName)),
		nft.NewAppModule(appCodec, app.nftKeeper, app.accountKeeper, app.bankKeeper),
		mt.NewAppModule(appCodec, app.mtKeeper, app.accountKeeper, app.bankKeeper),
		service.NewAppModule(appCodec, app.serviceKeeper, app.accountKeeper, app.bankKeeper, app.GetSubspace(servicetypes.ModuleName)),
		oracle.NewAppModule(appCodec, app.oracleKeeper, app.accountKeeper, app.bankKeeper),
		random.NewAppModule(appCodec, app.randomKeeper, app.accountKeeper, app.bankKeeper),
		identity.NewAppModule(app.identityKeeper),
		record.NewAppModule(appCodec, app.recordKeeper, app.accountKeeper, app.bankKeeper),
		node.NewAppModule(appCodec, app.nodeKeeper),
		tibc.NewAppModule(app.tibcKeeper),
		nfttransferModule,
		mttransferModule,
		// evm
		evm.NewAppModule(app.EvmKeeper, app.accountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	// extend Modules
	if appOptions.addModule != nil {
		appOptions.addModule(app, app.mm, app.keys)
	}

	app.mm.SetOrderMigrations(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		genutiltypes.ModuleName,
		feegrant.ModuleName,
		tibchost.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	app.mm.RegisterInvariants(app.crisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.accountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper, app.GetSubspace(banktypes.ModuleName)),
		feegrantmodule.NewAppModule(appCodec, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		cslashing.NewAppModule(appCodec, cslashing.NewKeeper(app.slashingKeeper, app.nodeKeeper), app.accountKeeper, app.bankKeeper, stakingKeeper),
		params.NewAppModule(app.paramsKeeper),
		cparams.NewAppModule(appCodec, app.paramsKeeper),
		record.NewAppModule(appCodec, app.recordKeeper, app.accountKeeper, app.bankKeeper),
		token.NewAppModule(appCodec, app.tokenKeeper, app.accountKeeper, app.bankKeeper, app.GetSubspace(tokentypes.ModuleName)),
		nft.NewAppModule(appCodec, app.nftKeeper, app.accountKeeper, app.bankKeeper),
		mt.NewAppModule(appCodec, app.mtKeeper, app.accountKeeper, app.bankKeeper),
		service.NewAppModule(appCodec, app.serviceKeeper, app.accountKeeper, app.bankKeeper, app.GetSubspace(servicetypes.ModuleName)),
		oracle.NewAppModule(appCodec, app.oracleKeeper, app.accountKeeper, app.bankKeeper),
		random.NewAppModule(appCodec, app.randomKeeper, app.accountKeeper, app.bankKeeper),
		identity.NewAppModule(app.identityKeeper),
		node.NewAppModule(appCodec, app.nodeKeeper),
		tibc.NewAppModule(app.tibcKeeper),
		// TODO
		// nfttransferModule,
		// mttransferModule,
		// evm
		evm.NewAppModule(app.EvmKeeper, app.accountKeeper, app.GetSubspace(evmtypes.ModuleName)),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(app.BuildAnteHandler(encodingConfig))
	app.SetEndBlocker(app.EndBlocker)

	// Set software upgrade execution logic
	// app.RegisterUpgradePlan("add-record-module",
	// 	store.StoreUpgrades{
	// 		Added: []string{recordtypes.StoreKey},
	// 	},
	// 	func(ctx sdk.Context, plan sdkupgrade.Plan) {},
	// )
	if appOptions.upgradePlan != nil {
		appOptions.upgradePlan(app, app.configurator, app.mm)
	}

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
		// ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		// app.capabilityKeeper.InitializeAndSeal(ctx)
	}
	return app
}

// Name returns the name of the App
func (app *IritaApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *IritaApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
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
	serviceGenState.Definitions = append(serviceGenState.Definitions, servicetypes.GenOraclePriceSvcDefinition())
	serviceGenState.Bindings = append(serviceGenState.Bindings, servicetypes.GenOraclePriceSvcBinding(tokentypesv1.GetNativeToken().MinUnit))
	serviceGenState.Definitions = append(serviceGenState.Definitions, servicetypes.GetRandomSvcDefinition())
	genesisState[servicetypes.ModuleName] = app.appCodec.MustMarshalJSON(&serviceGenState)

	app.upgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *IritaApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *IritaApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		app.Query,
	)
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
func (app *IritaApp) GetKey(storeKey string) *storetypes.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IritaApp) GetTKey(storeKey string) *storetypes.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *IritaApp) GetMemKey(storeKey string) *storetypes.MemoryStoreKey {
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

	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	grpcnode.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	if apiConfig.Swagger {
		lite.RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *IritaApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterNodeService implements types.Application.
func (app *IritaApp) RegisterNodeService(clientCtx client.Context) {
	grpcnode.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
}

// RegisterUpgradePlan implements the upgrade execution logic of the upgrade module
func (app *IritaApp) RegisterUpgradePlan(planName string,
	upgrades store.StoreUpgrades, upgradeHandler sdkupgrade.UpgradeHandler,
) {
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

// BuildAnteHandler constructs the ante handler for App
func (app *IritaApp) BuildAnteHandler(encodingConfig simappparams.EncodingConfig) sdk.AnteHandler {
	handlerOptions := appante.HandlerOptions{
		AccountKeeper:   app.accountKeeper,
		BankKeeper:      app.bankKeeper,
		TokenKeeper:     app.tokenKeeper,
		FeegrantKeeper:  app.feeGrantKeeper,
		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:  ethermintante.DefaultSigVerificationGasConsumer,

		// evm
		FeeMarketKeeper: app.FeeMarketKeeper,
		EvmKeeper:          app.EvmKeeper,
	}

	if appOptions.anteHandler != nil {
		return appOptions.anteHandler(app, handlerOptions)
	}
	return appante.NewAnteHandler(handlerOptions)
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// GetStoreKeys return app store keys list
func GetStoreKeys() []string {
	return storeKeys
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey storetypes.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(nodetypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(tokentypes.ModuleName)
	paramsKeeper.Subspace(recordtypes.ModuleName)
	paramsKeeper.Subspace(servicetypes.ModuleName)
	paramsKeeper.Subspace(tibchost.ModuleName)

	// evm
	paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(feemarkettypes.ModuleName)

	return paramsKeeper
}
