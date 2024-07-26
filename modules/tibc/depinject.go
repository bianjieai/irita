package tibc

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	clientkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/keeper"
	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	packetkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/keeper"
	routingkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/keeper"
	tibcroutingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"
	"github.com/bianjieai/tibc-go/modules/tibc/core/keeper"

	tibmodulev1 "github.com/bianjieai/irita/api/irita/tibc/module/v1"
	tibckeeper "github.com/bianjieai/irita/modules/tibc/keeper"
)

// App Wiring Setup
func init() {
	appmodule.Register(&tibmodulev1.Module{},
		appmodule.Provide(ProvideModule),
		appmodule.Invoke(InvokeTibcRouter),
	)
}

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// Inputs define the module inputs for the depinject.
type Inputs struct {
	depinject.In

	Config *tibmodulev1.Module
	Cdc    codec.Codec
	Key    *store.KVStoreKey

	StakingKeeper clienttypes.StakingKeeper
}

// Outputs define the module outputs for the depinject.
type Outputs struct {
	depinject.Out

	TibcKeeper    *tibckeeper.Keeper
	ClientKeeper  clientkeeper.Keeper
	PacketKeeper  packetkeeper.Keeper
	RoutingKeeper routingkeeper.Keeper
	Module        appmodule.AppModule
}

// ProvideModule defines a function that provides the TIBC module with necessary inputs and returns the outputs.
//
// Inputs: Inputs struct containing configuration, codec, store key, and staking keeper.
// Outputs: Outputs struct with TIBC keeper, module, client keeper, packet keeper, and routing keeper.
func ProvideModule(in Inputs) Outputs {
	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}

	coreKeeper := keeper.NewKeeper(
		in.Cdc,
		in.Key,
		in.StakingKeeper,
		authority.String(),
	)

	keeper := tibckeeper.NewKeeper(coreKeeper)
	m := NewAppModule(keeper)
	return Outputs{
		TibcKeeper:    keeper,
		ClientKeeper:  keeper.ClientKeeper,
		PacketKeeper:  keeper.PacketKeeper,
		RoutingKeeper: keeper.RoutingKeeper,
		Module:        m,
	}
}

// InvokeTibcRouter invokes the TIBC router with the TIBC modules.
func InvokeTibcRouter(tibcKeeper *tibckeeper.Keeper, tibcModules map[string]tibcroutingtypes.TIBCModule) {
	if tibcModules == nil || len(tibcModules) == 0 {
		return
	}

	modules := maps.Keys(tibcModules)
	slices.Sort(modules)

	tibcRouter := tibcroutingtypes.NewRouter()
	for _, moduleName := range modules {
		tibcRouter.AddRoute(tibcroutingtypes.Port(moduleName), tibcModules[moduleName])
	}
	tibcKeeper.SetRouter(tibcRouter)
}
