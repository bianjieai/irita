package app

import (
	appante "github.com/bianjieai/irita/app/ante"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
)

type AddModuleFun func(app *IritaApp, mm *module.Manager, keys map[string]*sdk.KVStoreKey)
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

func (app *IritaApp) GetAccountKeeper() authkeeper.AccountKeeper {
	return app.accountKeeper
}

func (app *IritaApp) GetBankKeeper() bankkeeper.Keeper {
	return app.bankKeeper
}

func (app *IritaApp) GetTokenKeeper() tokenkeeper.Keeper {
	return app.tokenKeeper
}

func (app *IritaApp) GetParamsKeeper() paramskeeper.Keeper {
	return app.paramsKeeper
}
