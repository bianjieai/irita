package app

import (
	"cosmossdk.io/depinject"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	appante "github.com/bianjieai/irita/app/ante"
)

type AddModuleFun func(app *IritaApp, mm *module.Manager, keys map[string]*storetypes.KVStoreKey)
type AnteHandlerFun func(app *IritaApp, handlerOptions appante.HandlerOptions) sdk.AnteHandler
type RegisterUpgradePlanFun func(app *IritaApp, configurator module.Configurator, mm *module.Manager)

type IritaAppOptions struct {
	addModule   AddModuleFun
	anteHandler AnteHandlerFun
	upgradePlan RegisterUpgradePlanFun
}

func NewAppOptions(addModule AddModuleFun, anteHandler AnteHandlerFun, registerUpgradePlan RegisterUpgradePlanFun, modules ...module.AppModuleBasic) {
	for _, moduleBasic := range modules {
		ModuleBasics[moduleBasic.Name()] = moduleBasic
	}
	appOptions = IritaAppOptions{
		addModule:   addModule,
		anteHandler: anteHandler,
		upgradePlan: registerUpgradePlan,
	}
}

// DepinjectOptions are passed to the app on creation 
type DepinjectOptions struct {
	Config    depinject.Config
	Providers []interface{}
	Consumers []interface{}
}
