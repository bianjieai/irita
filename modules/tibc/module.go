package tibc

import (
	"github.com/bianjieai/irita/modules/tibc/client/cli"
	"github.com/bianjieai/irita/modules/tibc/keeper"

	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	packettypes "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	host "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	"github.com/bianjieai/tibc-go/modules/tibc/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule defines the basic application module used by the tibc module.
type AppModule struct {
	tibc.AppModule
	k *keeper.Keeper
}

// GetTxCmd returns the root tx command for the tibc module.
func (AppModule) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// Route returns the message routing key for the tibc module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(host.RouterKey, NewHandler(*am.k))
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	clienttypes.RegisterMsgServer(cfg.MsgServer(), am.k.Keeper)
	packettypes.RegisterMsgServer(cfg.MsgServer(), am.k.Keeper)
	types.RegisterQueryService(cfg.QueryServer(), am.k.Keeper)
}

// NewAppModule creates a new AppModule object
func NewAppModule(k *keeper.Keeper) AppModule {
	return AppModule{
		AppModule: tibc.NewAppModule(k.Keeper),
		k:         k,
	}
}
