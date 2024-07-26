package tibc

import (
	"slices"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"

	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	clientkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/keeper"
	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	packetkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/keeper"
	routingkeeper "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/keeper"
	routingtypes "github.com/bianjieai/tibc-go/modules/tibc/core/26-routing/types"

	tibmodulev1 "github.com/bianjieai/irita/api/irita/tibc/module/v1"
	tibckeeper "github.com/bianjieai/irita/modules/tibc/keeper"
	"github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

// App Wiring Setup
func init() {
	appmodule.Register(&tibmodulev1.Module{},
		appmodule.Provide(ProvideModule),
		appmodule.Invoke(InvokeAddRoutes),
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

// InvokeAddRoutes adds routes to the TIBC router
func InvokeAddRoutes(keeper *tibckeeper.Keeper, routes []routingtypes.Route) {
	if keeper == nil || routes == nil {
		return
	}

	// Default route order is a lexical sort by RouteKey.
	// Explicit ordering can be added to the module config if required.
	slices.SortFunc(routes, func(a, b routingtypes.Route) int {
		if a.Port < b.Port {
			return -1
		}
		return 1
	})

	router := routingtypes.NewRouter()
	for _, r := range routes {
		router.Add(r)
	}
	keeper.SetRouter(router)
}
