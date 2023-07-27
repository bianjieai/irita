package app

import (
	"fmt"

	store "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

var router = &upgradeRouter{
	mu: make(map[string]Upgrade),
}

// RegisterUpgradePlans register a handler of upgrade plan
func (app *IritaApp) Register(us ...Upgrade) {
	router.registers(us...)
}

// RegisterUpgradePlans register a handler of upgrade plan
func (app *IritaApp) RegisterUpgradePlans() {
	app.setupUpgradeStoreLoaders()
	app.setupUpgradeHandlers()
}

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *IritaApp) setupUpgradeStoreLoaders() {
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
			router.upgradeInfo(upgradeInfo.Name).StoreUpgrades,
		),
	)
}

func (app *IritaApp) setupUpgradeHandlers() {
	for upgradeName, upgrade := range router.routers() {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgradeName,
			upgrade.UpgradeHandlerConstructor(
				app.mm,
				app.configurator,
				app,
			),
		)
	}
}

type upgradeRouter struct {
	mu map[string]Upgrade
}

func (r *upgradeRouter) registers(us ...Upgrade) {
	for _, u := range us {
		r.register(u)
	}
}

func (r *upgradeRouter) routers() map[string]Upgrade {
	return r.mu
}

func (r *upgradeRouter) upgradeInfo(planName string) Upgrade {
	return r.mu[planName]
}

func (r *upgradeRouter) register(u Upgrade) *upgradeRouter {
	if _, has := r.mu[u.UpgradeName]; has {
		panic(u.UpgradeName + " already registered")
	}
	r.mu[u.UpgradeName] = u
	return r
}

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// UpgradeHandlerConstructor defines the function that creates an upgrade handler
	UpgradeHandlerConstructor func(*module.Manager, module.Configurator, *IritaApp) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades *store.StoreUpgrades
}
