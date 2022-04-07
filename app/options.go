package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
)

func SetName(name string) func(app *baseapp.BaseApp) {
	return func(bap *baseapp.BaseApp) { bap.SetName(name) }
}
