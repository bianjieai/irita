package tibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"

	"github.com/bianjieai/irita/modules/tibc/keeper"
	clienttypes "github.com/bianjieai/irita/modules/tibc/types"
)

// NewHandler defines the TIBC handler
func NewHandler(k keeper.Keeper) sdk.Handler {
	tibcHandler := tibc.NewHandler(*k.Keeper)
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:
			res, err := k.CreateClient(ctx.Context(), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *clienttypes.MsgUpgradeClient:
			res, err := k.UpgradeClient(ctx.Context(), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *clienttypes.MsgRegisterRelayer:
			res, err := k.RegisterRelayer(ctx.Context(), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *clienttypes.MsgSetRoutingRules:
			res, err := k.SetRoutingRules(ctx.Context(), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			res, err := tibcHandler(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		}
	}
}
