package app

import (
	runtimev1alpha1 "cosmossdk.io/api/cosmos/app/runtime/v1alpha1"
	appv1alpha1 "cosmossdk.io/api/cosmos/app/v1alpha1"
	authmodulev1 "cosmossdk.io/api/cosmos/auth/module/v1"
	bankmodulev1 "cosmossdk.io/api/cosmos/bank/module/v1"
	consensusmodulev1 "cosmossdk.io/api/cosmos/consensus/module/v1"
	crisismodulev1 "cosmossdk.io/api/cosmos/crisis/module/v1"
	evidencemodulev1 "cosmossdk.io/api/cosmos/evidence/module/v1"
	feegrantmodulev1 "cosmossdk.io/api/cosmos/feegrant/module/v1"
	txconfigv1 "cosmossdk.io/api/cosmos/tx/config/v1"
	"cosmossdk.io/core/appconfig"

	tibmodulev1 "github.com/bianjieai/irita/api/irita/tibc/module/v1"
	_ "github.com/bianjieai/irita/modules/tibc"               // import for side-effects
	tibctypes "github.com/bianjieai/irita/modules/tibc/types" // import for side-effects
	tibcmtmodulev1 "github.com/bianjieai/tibc-go/api/tibc/apps/mt_transfer/module/v1"
	tibcnftmodulev1 "github.com/bianjieai/tibc-go/api/tibc/apps/nft_transfer/module/v1"
	_ "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer" // import for side-effects
	tibcmttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/types"
	_ "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer" // import for side-effects
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/group"
	evmmodulev1 "github.com/evmos/ethermint/api/ethermint/evm/module/v1"
	feemarketmodulev1 "github.com/evmos/ethermint/api/ethermint/feemarket/module/v1"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	feemarkettypes "github.com/evmos/ethermint/x/feemarket/types"
	genutilmodulev1 "iritamod.bianjie.ai/api/iritamod/genutil/module/v1"
	identitymodulev1 "iritamod.bianjie.ai/api/iritamod/identity/module/v1"
	nodemodulev1 "iritamod.bianjie.ai/api/iritamod/node/module/v1"
	paramsmodulev1 "iritamod.bianjie.ai/api/iritamod/params/module/v1"
	slashingmodulev1 "iritamod.bianjie.ai/api/iritamod/slashing/module/v1"
	upgrademodulev1 "iritamod.bianjie.ai/api/iritamod/upgrade/module/v1"
	_ "iritamod.bianjie.ai/modules/genutil" // import for side-effects
	genutiltypes "iritamod.bianjie.ai/modules/genutil"
	_ "iritamod.bianjie.ai/modules/identity" // import for side-effects
	identitytypes "iritamod.bianjie.ai/modules/identity/types"
	_ "iritamod.bianjie.ai/modules/node" // import for side-effects
	nodetypes "iritamod.bianjie.ai/modules/node/types"
	_ "iritamod.bianjie.ai/modules/params" // import for side-effects
	paramstypes "iritamod.bianjie.ai/modules/params/types"
	_ "iritamod.bianjie.ai/modules/slashing" // import for side-effects
	slashingtypes "iritamod.bianjie.ai/modules/slashing/types"
	_ "iritamod.bianjie.ai/modules/upgrade" // import for side-effects
	upgradetypes "iritamod.bianjie.ai/modules/upgrade/types"
	mtmodule "mods.irisnet.org/api/irismod/mt/module/v1"
	nftmodule "mods.irisnet.org/api/irismod/nft/module/v1"
	oraclemodule "mods.irisnet.org/api/irismod/oracle/module/v1"
	randommodule "mods.irisnet.org/api/irismod/random/module/v1"
	recordmodule "mods.irisnet.org/api/irismod/record/module/v1"
	servicemodule "mods.irisnet.org/api/irismod/service/module/v1"
	tokenmodule "mods.irisnet.org/api/irismod/token/module/v1"
	_ "mods.irisnet.org/modules/mt" // import for side-effects
	mttypes "mods.irisnet.org/modules/mt/types"
	nfttypes "mods.irisnet.org/modules/nft/types"
	_ "mods.irisnet.org/modules/oracle" // import for side-effects
	oracletypes "mods.irisnet.org/modules/oracle/types"
	_ "mods.irisnet.org/modules/random" // import for side-effects
	randomtypes "mods.irisnet.org/modules/random/types"
	_ "mods.irisnet.org/modules/record" // import for side-effects
	recordtypes "mods.irisnet.org/modules/record/types"
	_ "mods.irisnet.org/modules/service" // import for side-effects
	servicetypes "mods.irisnet.org/modules/service/types"
	_ "mods.irisnet.org/modules/token" // import for side-effects
	tokentypes "mods.irisnet.org/modules/token/types"
)

var (

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	genesisModuleOrder = []string{
		authtypes.ModuleName,
		banktypes.ModuleName,
		nodetypes.ModuleName,
		slashingtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		consensustypes.ModuleName,
		mttypes.ModuleName,
		nfttypes.ModuleName,
		servicetypes.ModuleName,
		oracletypes.ModuleName,
		randomtypes.ModuleName,
		recordtypes.ModuleName,
		identitytypes.ModuleName,
		tokentypes.ModuleName,
		tibctypes.ModuleName,
		tibcnfttypes.ModuleName,
		tibcmttypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
	}

	// module account permissions
	moduleAccPerms = []*authmodulev1.ModuleAccountPermission{
		{Account: authtypes.FeeCollectorName},
		{Account: nfttypes.ModuleName},
		{Account: mttypes.ModuleName},
		{Account: servicetypes.DepositAccName, Permissions: []string{authtypes.Burner}},
		{Account: servicetypes.RequestAccName},
		{Account: servicetypes.FeeCollectorName, Permissions: []string{authtypes.Burner}},
		{Account: tokentypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: evmtypes.ModuleName, Permissions: []string{authtypes.Minter, authtypes.Burner}},
		{Account: tibcnfttypes.ModuleName},
		{Account: tibcmttypes.ModuleName},
	}

	// blocked account addresses
	blockAccAddrs = []string{
		authtypes.FeeCollectorName,
		servicetypes.DepositAccName,
		servicetypes.RequestAccName,
		servicetypes.FeeCollectorName,
		tokentypes.ModuleName,
		evmtypes.ModuleName,
		feemarkettypes.ModuleName,
		// We allow the following module accounts to receive funds:
		// govtypes.ModuleName
	}

	// AppConfig application configuration (used by depinject)
	AppConfig = appconfig.Compose(&appv1alpha1.Config{
		Modules: []*appv1alpha1.ModuleConfig{
			{
				Name: "runtime",
				Config: appconfig.WrapAny(&runtimev1alpha1.Module{
					AppName: "SimApp",
					// During begin block slashing happens after distr.BeginBlocker so that
					// there is nothing left over in the validator fee pool, so as to keep the
					// CanWithdrawInvariant invariant.
					// NOTE: staking module is required if HistoricalEntries param > 0
					// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
					BeginBlockers: []string{
						upgradetypes.ModuleName,
						slashingtypes.ModuleName,
						evidencetypes.ModuleName,
						nodetypes.ModuleName,
						authtypes.ModuleName,
						banktypes.ModuleName,
						crisistypes.ModuleName,
						genutiltypes.ModuleName,
						feegrant.ModuleName,
						paramstypes.ModuleName,
						identitytypes.ModuleName,
						mttypes.ModuleName,
						nfttypes.ModuleName,
						servicetypes.ModuleName,
						oracletypes.ModuleName,
						randomtypes.ModuleName,
						recordtypes.ModuleName,
						tokentypes.ModuleName,
						consensustypes.ModuleName,
						evmtypes.ModuleName,
						feemarkettypes.ModuleName,
						tibctypes.ModuleName,
						tibcnfttypes.ModuleName,
						tibcmttypes.ModuleName,
					},
					EndBlockers: []string{
						crisistypes.ModuleName,
						nodetypes.ModuleName,
						authtypes.ModuleName,
						banktypes.ModuleName,
						slashingtypes.ModuleName,
						genutiltypes.ModuleName,
						evidencetypes.ModuleName,
						authz.ModuleName,
						feegrant.ModuleName,
						group.ModuleName,
						paramstypes.ModuleName,
						upgradetypes.ModuleName,
						consensustypes.ModuleName,
						identitytypes.ModuleName,
						mttypes.ModuleName,
						nfttypes.ModuleName,
						servicetypes.ModuleName,
						oracletypes.ModuleName,
						randomtypes.ModuleName,
						recordtypes.ModuleName,
						tokentypes.ModuleName,
						evmtypes.ModuleName,
						feemarkettypes.ModuleName,
						tibctypes.ModuleName,
						tibcnfttypes.ModuleName,
						tibcmttypes.ModuleName,
					},
					OverrideStoreKeys: []*runtimev1alpha1.StoreKeyConfig{
						{
							ModuleName: authtypes.ModuleName,
							KvStoreKey: "acc",
						},
					},
					InitGenesis: genesisModuleOrder,
					// When ExportGenesis is not specified, the export genesis module order
					// is equal to the init genesis order
					// ExportGenesis: genesisModuleOrder,
					// Uncomment if you want to set a custom migration order here.
					// OrderMigrations: nil,
				}),
			},
			{
				Name: authtypes.ModuleName,
				Config: appconfig.WrapAny(&authmodulev1.Module{
					Bech32Prefix:             "iaa",
					ModuleAccountPermissions: moduleAccPerms,
				}),
			},
			{
				Name: banktypes.ModuleName,
				Config: appconfig.WrapAny(&bankmodulev1.Module{
					BlockedModuleAccountsOverride: blockAccAddrs,
				}),
			},
			{
				Name:   nodetypes.ModuleName,
				Config: appconfig.WrapAny(&nodemodulev1.Module{}),
			},
			{
				Name:   slashingtypes.ModuleName,
				Config: appconfig.WrapAny(&slashingmodulev1.Module{}),
			},
			{
				Name:   paramstypes.ModuleName,
				Config: appconfig.WrapAny(&paramsmodulev1.Module{}),
			},
			{
				Name:   "tx",
				Config: appconfig.WrapAny(&txconfigv1.Config{}),
			},
			{
				Name:   genutiltypes.ModuleName,
				Config: appconfig.WrapAny(&genutilmodulev1.Module{}),
			},
			{
				Name:   upgradetypes.ModuleName,
				Config: appconfig.WrapAny(&upgrademodulev1.Module{}),
			},
			{
				Name:   evidencetypes.ModuleName,
				Config: appconfig.WrapAny(&evidencemodulev1.Module{}),
			},
			{
				Name:   feegrant.ModuleName,
				Config: appconfig.WrapAny(&feegrantmodulev1.Module{}),
			},
			{
				Name:   crisistypes.ModuleName,
				Config: appconfig.WrapAny(&crisismodulev1.Module{}),
			},
			{
				Name:   consensustypes.ModuleName,
				Config: appconfig.WrapAny(&consensusmodulev1.Module{}),
			},
			{
				Name:   identitytypes.ModuleName,
				Config: appconfig.WrapAny(&identitymodulev1.Module{}),
			},
			{
				Name:   nfttypes.ModuleName,
				Config: appconfig.WrapAny(&nftmodule.Module{}),
			},
			{
				Name:   mttypes.ModuleName,
				Config: appconfig.WrapAny(&mtmodule.Module{}),
			},
			{
				Name:   oracletypes.ModuleName,
				Config: appconfig.WrapAny(&oraclemodule.Module{}),
			},
			{
				Name: servicetypes.ModuleName,
				Config: appconfig.WrapAny(&servicemodule.Module{
					FeeCollectorName: servicetypes.FeeCollectorName,
				}),
			},
			{
				Name:   randomtypes.ModuleName,
				Config: appconfig.WrapAny(&randommodule.Module{}),
			},
			{
				Name:   recordtypes.ModuleName,
				Config: appconfig.WrapAny(&recordmodule.Module{}),
			},
			{
				Name: tokentypes.ModuleName,
				Config: appconfig.WrapAny(&tokenmodule.Module{
					FeeCollectorName: authtypes.FeeCollectorName,
				}),
			},
			{
				Name:   evmtypes.ModuleName,
				Config: appconfig.WrapAny(&evmmodulev1.Module{
					// Tracer:             "mockTracer",
				}),
			},
			{
				Name:   feemarkettypes.ModuleName,
				Config: appconfig.WrapAny(&feemarketmodulev1.Module{}),
			},
			{
				Name:   tibctypes.ModuleName,
				Config: appconfig.WrapAny(&tibmodulev1.Module{}),
			},
			{
				Name:   tibcnfttypes.ModuleName,
				Config: appconfig.WrapAny(&tibcnftmodulev1.Module{}),
			},
			{
				Name:   tibcmttypes.ModuleName,
				Config: appconfig.WrapAny(&tibcmtmodulev1.Module{}),
			},
		},
	})
)

// DefaultDepinjectOptions returns the default depinject options
func DefaultDepinjectOptions() DepinjectOptions {
	return DepinjectOptions{
		Config:    AppConfig,
		Providers: []interface{}{},
		Consumers: []interface{}{},
	}
}
