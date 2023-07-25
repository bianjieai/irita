package app

import (
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/spf13/cast"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	tmos "github.com/cometbft/cometbft/libs/os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	grpcnode "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
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
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
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
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	sdkupgrade "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/irisnet/irismod/modules/mt"
	mtkeeper "github.com/irisnet/irismod/modules/mt/keeper"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
	nft "github.com/irisnet/irismod/modules/nft/module"
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

	appante "github.com/bianjieai/irita/app/ante"
	sidechainmodule "github.com/bianjieai/irita/modules/side-chain"
	"github.com/bianjieai/iritamod/modules/genutil"
	genutiltypes "github.com/bianjieai/iritamod/modules/genutil"
	"github.com/bianjieai/iritamod/modules/identity"
	identitykeeper "github.com/bianjieai/iritamod/modules/identity/keeper"
	identitytypes "github.com/bianjieai/iritamod/modules/identity/types"
	"github.com/bianjieai/iritamod/modules/node"
	nodekeeper "github.com/bianjieai/iritamod/modules/node/keeper"
	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	cparams "github.com/bianjieai/iritamod/modules/params"
	cparamskeeper "github.com/bianjieai/iritamod/modules/params/keeper"
	cparamstypes "github.com/bianjieai/iritamod/modules/params/types"
	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
	permtypes "github.com/bianjieai/iritamod/modules/perm/types"
	sidechain "github.com/bianjieai/iritamod/modules/side-chain"
	sidechainkeeper "github.com/bianjieai/iritamod/modules/side-chain/keeper"
	sidechaintypes "github.com/bianjieai/iritamod/modules/side-chain/types"
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

	tibcmttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer"
	tibcmttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/keeper"
	tibcmttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/types"
	tibcnfttransfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibcnfttransferkeeper "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/keeper"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	tibccorekeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"

	ethermintante "github.com/evmos/ethermint/app/ante"
	srvflags "github.com/evmos/ethermint/server/flags"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/evmos/ethermint/x/evm"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	"github.com/evmos/ethermint/x/feemarket"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
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
		tibcmttransfer.AppModuleBasic{},
		sidechain.AppModuleBasic{},

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
		tibcmttypes.ModuleName:              nil,
		sidechaintypes.ModuleName:           nil,

		// evm
		evmtypes.ModuleName: {
			authtypes.Minter,
			authtypes.Burner,
		}, // used for secure addition and subtraction of balance using module account
	}
	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc   = map[string]bool{}
	allowUpdateByCparamsMsgs = []string{
		"/cosmos.auth.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/iritamod.slashing.MsgUpdateParams",
		"/iritamod.node.MsgUpdateParams",
	}
)

// Verify app interface at compile time
// var _ simapp.App = (*IritaApp)(nil)

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
	keys    map[string]*storetypes.KVStoreKey
	tkeys   map[string]*storetypes.TransientStoreKey
	memKeys map[string]*storetypes.MemoryStoreKey

	// keepers
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	EvidenceKeeper        *evidencekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensusparamkeeper.Keeper

	SlashingKeeper       cslashing.Keeper
	UpgradeKeeper        *upgradekeeper.Keeper
	RecordKeeper         recordkeeper.Keeper
	TokenKeeper          tokenkeeper.Keeper
	NftKeeper            nftkeeper.Keeper
	MtKeeper             mtkeeper.Keeper
	ServiceKeeper        servicekeeper.Keeper
	OracleKeeper         oraclekeeper.Keeper
	RandomKeeper         randomkeeper.Keeper
	PermKeeper           permkeeper.Keeper
	IdentityKeeper       identitykeeper.Keeper
	NodeKeeper           nodekeeper.Keeper
	OpbKeeper            opbkeeper.Keeper
	SidechainKeeper      sidechainkeeper.Keeper
	TibcKeeper           *tibckeeper.Keeper
	NftTransferKeeper    tibcnfttransferkeeper.Keeper
	MtTransferKeeper     tibcmttransferkeeper.Keeper
	ScopedTIBCKeeper     capabilitykeeper.ScopedKeeper
	ScopedTIBCMockKeeper capabilitykeeper.ScopedKeeper

	EvmKeeper       *appkeeper.Keeper
	FeeMarketKeeper feemarketkeeper.Keeper

	// Ethermint keepers

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
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *IritaApp {
	// TODO: Remove cdc in favor of appCodec once all modules are migrated.
	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(
		appName,
		logger,
		db,
		encodingConfig.TxConfig.TxDecoder(),
		baseAppOptions...)
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
		tibcmttypes.StoreKey,
		sidechaintypes.StoreKey,
		consensusparamtypes.StoreKey,
		// evm
		evmtypes.StoreKey, feemarkettypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(
		paramstypes.TStoreKey,
		evmtypes.TransientKey,
		feemarkettypes.StoreKey,
	)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &IritaApp{
		BaseApp:           bApp,
		cdc:               legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	app.ConsensusParamsKeeper = consensusparamkeeper.NewKeeper(
		appCodec,
		keys[upgradetypes.StoreKey],
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(&app.ConsensusParamsKeeper)

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		ethermint.ProtoAccount, //Note: replace authtypes.ProtoBaseAccount
		maccPerms,
		sdk.Bech32MainPrefix,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.ModuleAccountAddrs(),
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	app.NodeKeeper = node.NewKeeper(
		appCodec,
		keys[nodetypes.StoreKey],
	)

	app.SlashingKeeper = cslashing.NewKeeper(
		appCodec,
		legacyAmino,
		keys[slashingtypes.StoreKey],
		&app.NodeKeeper,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		appCodec,
		keys[crisistypes.StoreKey],
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	)

	// create evidence keeper with router
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.NodeKeeper, app.SlashingKeeper,
	)

	app.TokenKeeper = tokenkeeper.NewKeeper(
		appCodec,
		keys[tokentypes.StoreKey],
		app.GetSubspace(tokentypes.ModuleName),
		app.BankKeeper,
		app.ModuleAccountAddrs(),
		opbtypes.PointTokenFeeCollectorName,
	)

	app.RecordKeeper = recordkeeper.NewKeeper(appCodec, keys[recordtypes.StoreKey])
	app.NftKeeper = nftkeeper.NewKeeper(appCodec, keys[nfttypes.StoreKey])
	app.MtKeeper = mtkeeper.NewKeeper(appCodec, keys[mttypes.StoreKey])

	app.ServiceKeeper = servicekeeper.NewKeeper(
		appCodec,
		keys[servicetypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(
			servicetypes.ModuleName,
		),
		app.ModuleAccountAddrs(),
		opbtypes.PointTokenFeeCollectorName,
	)

	app.OracleKeeper = oraclekeeper.NewKeeper(
		appCodec,
		keys[oracletypes.StoreKey],
		app.GetSubspace(oracletypes.ModuleName),
		app.ServiceKeeper,
	)

	app.RandomKeeper = randomkeeper.NewKeeper(
		appCodec,
		keys[randomtypes.StoreKey],
		app.BankKeeper,
		app.ServiceKeeper,
	)

	app.NodeKeeper = *app.NodeKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.SlashingKeeper.Hooks()),
	)

	permKeeper := permkeeper.NewKeeper(appCodec, keys[permtypes.StoreKey])
	app.PermKeeper = appante.RegisterAccessControl(permKeeper)

	app.IdentityKeeper = identitykeeper.NewKeeper(appCodec, keys[identitytypes.StoreKey])

	app.OpbKeeper = opbkeeper.NewKeeper(
		appCodec, keys[opbtypes.StoreKey], app.AccountKeeper,
		app.BankKeeper, app.TokenKeeper, app.PermKeeper,
		app.GetSubspace(opbtypes.ModuleName),
	)

	sidechainPermKeeper := sidechainmodule.NewPermKeeper(appCodec, app.PermKeeper)
	app.SidechainKeeper = sidechainkeeper.NewKeeper(
		appCodec,
		keys[sidechaintypes.StoreKey],
		app.AccountKeeper,
	)

	// evm
	tracer := cast.ToString(appOpts.Get(srvflags.EVMTracer))

	// Create Ethermint  keepers
	app.FeeMarketKeeper = feemarketkeeper.NewKeeper(
		appCodec,
		authtypes.NewModuleAddress(cparamstypes.ModuleName),
		keys[feemarkettypes.StoreKey],
		tkeys[feemarkettypes.StoreKey],
		app.GetSubspace(feemarkettypes.ModuleName),
	)
	app.EvmKeeper = appkeeper.NewKeeper(
		appCodec,
		keys[evmtypes.StoreKey],
		tkeys[evmtypes.TransientKey],
		authtypes.NewModuleAddress(cparamstypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		appkeeper.WNodeKeeper{Keeper: app.NodeKeeper},
		app.FeeMarketKeeper,
		nil,
		nil,
		tracer,
		app.GetSubspace(evmtypes.ModuleName),
	)

	opbValidator := appkeeper.NewEthOpbValidator(
		app.OpbKeeper,
		app.TokenKeeper,
		app.EvmKeeper,
		app.PermKeeper,
	)

	app.EvmKeeper.SetValidator(
		func(ctx sdk.Context) vm.CanTransferFunc {
			opbValidator.With(ctx)
			return opbValidator.CanTransfer
		},
		func(ctx sdk.Context) vm.TransferFunc {
			opbValidator.With(ctx)
			return opbValidator.Transfer
		},
	)

	app.TibcKeeper = tibckeeper.NewKeeper(tibccorekeeper.NewKeeper(
		appCodec,
		keys[tibchost.StoreKey],
		app.NodeKeeper,
		authtypes.NewModuleAddress(cparamstypes.ModuleName).String(),
	))

	app.NftTransferKeeper = tibcnfttransferkeeper.NewKeeper(
		appCodec, keys[tibcnfttypes.StoreKey], app.GetSubspace(tibcnfttypes.ModuleName),
		app.AccountKeeper, tibckeeper.WrapNftKeeper(app.NftKeeper),
		app.TibcKeeper.PacketKeeper, app.TibcKeeper.ClientKeeper,
	)

	app.MtTransferKeeper = tibcmttransferkeeper.NewKeeper(
		appCodec, keys[tibcmttypes.StoreKey], app.GetSubspace(tibcmttypes.ModuleName),
		app.AccountKeeper, app.MtKeeper,
		app.TibcKeeper.PacketKeeper, app.TibcKeeper.ClientKeeper,
	)

	cparamsKeeper := cparamskeeper.NewKeeper(
		app.AccountKeeper, app.MsgServiceRouter(), allowUpdateByCparamsMsgs)

	nfttransferModule := tibcnfttransfer.NewAppModule(app.NftTransferKeeper)
	mttransferModule := tibcmttransfer.NewAppModule(app.MtTransferKeeper)
	tibcRouter := tibcroutingtypes.NewRouter()
	tibcRouter.AddRoute(tibcnfttypes.ModuleName, nfttransferModule)
	tibcRouter.AddRoute(tibcmttypes.ModuleName, mttransferModule)
	app.TibcKeeper.SetRouter(tibcRouter)
	/****  Module Options ****/
	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)

	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.NodeKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
		bank.NewAppModule(
			appCodec,
			app.BankKeeper,
			app.AccountKeeper,
			app.GetSubspace(banktypes.ModuleName),
		),
		crisis.NewAppModule(
			app.CrisisKeeper,
			skipGenesisInvariants,
			app.GetSubspace(crisistypes.ModuleName),
		),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		cslashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			&app.NodeKeeper,
			app.GetSubspace(cslashing.ModuleName),
		),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(*app.EvidenceKeeper),
		params.NewAppModule(app.ParamsKeeper),
		cparams.NewAppModule(appCodec, cparamsKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		nft.NewAppModule(appCodec, app.NftKeeper, app.AccountKeeper, app.BankKeeper),
		mt.NewAppModule(appCodec, app.MtKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper, app.GetSubspace(node.ModuleName)),
		opb.NewAppModule(appCodec, app.OpbKeeper),
		tibc.NewAppModule(app.TibcKeeper),
		sidechain.NewAppModule(appCodec, app.SidechainKeeper),
		nfttransferModule,
		mttransferModule,
		// evm
		appkeeper.NewAppModule(
			app.EvmKeeper,
			app.AccountKeeper,
			app.GetSubspace(evmtypes.ModuleName),
		),
		feemarket.NewAppModule(app.FeeMarketKeeper, app.GetSubspace(feemarkettypes.ModuleName)),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		sidechaintypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		sidechaintypes.ModuleName,

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
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		sidechaintypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	app.mm.SetOrderMigrations(
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
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
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		sidechaintypes.ModuleName,

		// evm
		evmtypes.ModuleName, feemarkettypes.ModuleName,
	)

	app.mm.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(
		app.appCodec,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
	)
	app.mm.RegisterServices(app.configurator)

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(
			appCodec,
			app.AccountKeeper,
			authsims.RandomGenesisAccounts,
			app.GetSubspace(authtypes.ModuleName),
		),
		bank.NewAppModule(
			appCodec,
			app.BankKeeper,
			app.AccountKeeper,
			app.GetSubspace(banktypes.ModuleName),
		),
		feegrantmodule.NewAppModule(
			appCodec,
			app.AccountKeeper,
			app.BankKeeper,
			app.FeeGrantKeeper,
			app.interfaceRegistry,
		),
		cslashing.NewAppModule(
			appCodec,
			app.SlashingKeeper,
			&app.NodeKeeper,
			app.GetSubspace(cslashing.ModuleName),
		),
		params.NewAppModule(app.ParamsKeeper),
		cparams.NewAppModule(appCodec, cparamsKeeper),
		record.NewAppModule(appCodec, app.RecordKeeper, app.AccountKeeper, app.BankKeeper),
		token.NewAppModule(appCodec, app.TokenKeeper, app.AccountKeeper, app.BankKeeper),
		nft.NewAppModule(appCodec, app.NftKeeper, app.AccountKeeper, app.BankKeeper),
		mt.NewAppModule(appCodec, app.MtKeeper, app.AccountKeeper, app.BankKeeper),
		service.NewAppModule(appCodec, app.ServiceKeeper, app.AccountKeeper, app.BankKeeper),
		oracle.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper),
		random.NewAppModule(appCodec, app.RandomKeeper, app.AccountKeeper, app.BankKeeper),
		perm.NewAppModule(appCodec, app.PermKeeper),
		identity.NewAppModule(app.IdentityKeeper),
		node.NewAppModule(appCodec, app.NodeKeeper, app.GetSubspace(node.ModuleName)),
		opb.NewAppModule(appCodec, app.OpbKeeper),
		tibc.NewAppModule(app.TibcKeeper),
		// evm
		appkeeper.NewAppModule(
			app.EvmKeeper,
			app.AccountKeeper,
			app.GetSubspace(evmtypes.ModuleName),
		),
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
	anteHandler := appante.DefaultAnteHandler(
		appante.HandlerOptions{
			PermKeeper:          app.PermKeeper,
			AccountKeeper:       app.AccountKeeper,
			BankKeeper:          app.BankKeeper,
			TokenKeeper:         app.TokenKeeper,
			OpbKeeper:           app.OpbKeeper,
			SignModeHandler:     encodingConfig.TxConfig.SignModeHandler(),
			FeegrantKeeper:      app.FeeGrantKeeper,
			SigGasConsumer:      ethermintante.DefaultSigVerificationGasConsumer,
			SideChainKeeper:     app.SidechainKeeper,
			SideChainPermKeeper: sidechainPermKeeper,

			// evm
			EvmFeeMarketKeeper: app.FeeMarketKeeper,
			EvmKeeper:          app.EvmKeeper,
		},
	)
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
	return app
}

// Name returns the name of the App
func (app *IritaApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *IritaApp) BeginBlocker(
	ctx sdk.Context,
	req abci.RequestBeginBlock,
) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *IritaApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *IritaApp) InitChainer(
	ctx sdk.Context,
	req abci.RequestInitChain,
) abci.ResponseInitChain {
	var genesisState GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	// add system service at InitChainer, overwrite if it exists
	var serviceGenState servicetypes.GenesisState
	app.appCodec.MustUnmarshalJSON(genesisState[servicetypes.ModuleName], &serviceGenState)

	serviceGenState.Definitions = append(
		serviceGenState.Definitions,
		servicetypes.GenOraclePriceSvcDefinition(),
	)
	serviceGenState.Bindings = append(
		serviceGenState.Bindings,
		servicetypes.GenOraclePriceSvcBinding(tokentypes.GetNativeToken().MinUnit),
	)
	serviceGenState.Definitions = append(
		serviceGenState.Definitions,
		randomtypes.GetSvcDefinition(),
	)
	genesisState[servicetypes.ModuleName] = app.appCodec.MustMarshalJSON(&serviceGenState)

	app.UpgradeKeeper.UpgradeKeeper().SetModuleVersionMap(ctx, app.mm.GetVersionMap())
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

// RegisterNodeService registers the node gRPC service on the provided
// application gRPC query router.
func (app *IritaApp) RegisterNodeService(clientCtx client.Context) {
	grpcnode.RegisterNodeService(clientCtx, app.GRPCQueryRouter())
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
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
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
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register node gRPC service for grpc-gateway.
	grpcnode.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		lite.RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *IritaApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(
		app.BaseApp.GRPCQueryRouter(),
		clientCtx,
		app.BaseApp.Simulate,
		app.interfaceRegistry,
	)
}

// RegisterUpgradePlan implements the upgrade execution logic of the upgrade module
func (app *IritaApp) RegisterUpgradePlan(planName string,
	upgrades storetypes.StoreUpgrades, upgradeHandler sdkupgrade.UpgradeHandler) {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		app.Logger().Info("not found upgrade plan", "planName", planName, "err", err.Error())
		return
	}

	if upgradeInfo.Name == planName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		// this configures a no-op upgrade handler for the planName upgrade
		app.UpgradeKeeper.SetUpgradeHandler(planName, upgradeHandler)
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
func initParamsKeeper(
	appCodec codec.BinaryCodec,
	legacyAmino *codec.LegacyAmino,
	key, tkey storetypes.StoreKey,
) paramskeeper.Keeper {
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

	// evm
	paramsKeeper.Subspace(evmtypes.ModuleName)
	paramsKeeper.Subspace(feemarkettypes.ModuleName)

	return paramsKeeper
}
