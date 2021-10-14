package tibc

import (
	"github.com/bianjieai/irita/modules/tibc/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/types/module"

	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
	"github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule defines the basic application module used by the tibc module.
type AppModule struct {
	tibc.AppModule
}

// GetTxCmd returns the root tx command for the tibc module.
func (AppModule) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// Route returns the message routing key for the tibc module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(host.RouterKey, NewHandler())
}

// NewAppModule creates a new AppModule object
func NewAppModule(k *keeper.Keeper) AppModule {
	return AppModule{
		tibc.NewAppModule(k),
	}
}
