package tibc

import (
	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	packettypes "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	tibcorekeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
	"github.com/bianjieai/tibc-go/modules/tibc/core/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/spf13/cobra"

	"github.com/bianjieai/irita/modules/tibc/client/cli"
	"github.com/bianjieai/irita/modules/tibc/keeper"
	tibctypes "github.com/bianjieai/irita/modules/tibc/types"
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

// RegisterInterfaces registers module concrete types into protobuf Any.
func (am AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
	tibctypes.RegisterInterfaces(registry)
}

// RegisterServices registers module services.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	tibctypes.RegisterMsgServer(cfg.MsgServer(), am.k)
	types.RegisterQueryService(cfg.QueryServer(), am.k.Keeper)

	msgserver := tibcorekeeper.NewMsgServerImpl(*am.k.Keeper)
	clienttypes.RegisterMsgServer(cfg.MsgServer(), msgserver)
	packettypes.RegisterMsgServer(cfg.MsgServer(), msgserver)
}

// NewAppModule creates a new AppModule object
func NewAppModule(k *keeper.Keeper) AppModule {
	return AppModule{
		AppModule: tibc.NewAppModule(k.Keeper),
		k:         k,
	}
}
