package app

import (
	"fmt"
	types "github.com/bianjieai/irita/app/upgrades"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var (
	UpgradeRouter = types.NewUpgradeRouter()
	//Register(v302.Upgrade).
	//Register(v302_rc.Upgrade)
)

// RegisterUpgradePlans register a handler of upgrade plan
func (app *IritaAppV2) RegisterUpgradePlans() {
	app.setupUpgradeStoreLoaders()
	app.setupUpgradeHandlers()
}

func (app *IritaAppV2) appKeepers() types.AppKeepers {
	return types.AppKeepers{
		AppCodec:              app.appCodec,
		AccountKeeper:         app.AccountKeeper,
		BankKeeper:            app.BankKeeper,
		SlashingKeeper:        app.SlashingKeeper,
		CrisisKeeper:          app.CrisisKeeper,
		UpgradeKeeper:         app.UpgradeKeeper,
		ParamsKeeper:          app.ParamsKeeper,
		EvidenceKeeper:        app.EvidenceKeeper,
		RecordKeeper:          app.RecordKeeper,
		TokenKeeper:           app.TokenKeeper,
		NftKeeper:             app.NftKeeper,
		MtKeeper:              app.MtKeeper,
		ServiceKeeper:         app.ServiceKeeper,
		OracleKeeper:          app.OracleKeeper,
		RandomKeeper:          app.RandomKeeper,
		IdentityKeeper:        app.IdentityKeeper,
		NodeKeeper:            app.NodeKeeper,
		FeeGrantKeeper:        app.FeeGrantKeeper,
		CapabilityKeeper:      app.CapabilityKeeper,
		ConsensusParamsKeeper: app.ConsensusParamsKeeper,
		EvmKeeper:             app.EvmKeeper,
		FeeMarketKeeper:       app.FeeMarketKeeper,
	}
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *IritaAppV2) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	app.SetStoreLoader(
		upgradetypes.UpgradeStoreLoader(
			upgradeInfo.Height,
			UpgradeRouter.UpgradeInfo(upgradeInfo.Name).StoreUpgrades,
		),
	)
}

func (app *IritaAppV2) setupUpgradeHandlers() {
	for upgradeName, upgrade := range UpgradeRouter.Routers() {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgradeName,
			upgrade.UpgradeHandlerConstructor(
				app.ModuleManager,
				app.Configurator(),
				app.appKeepers(),
			),
		)
	}
}
