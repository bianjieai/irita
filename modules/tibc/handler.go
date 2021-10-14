package tibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/tibc/keeper"
	clienttypes "github.com/bianjieai/irita/modules/tibc/types"
	tibc "github.com/bianjieai/tibc-go/modules/tibc/core"
)

// NewHandler defines the TIBC handler
func NewHandler(k keeper.Keeper) sdk.Handler {
	tibcHandler := tibc.NewHandler(*k.Keeper)
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:
			res, err := k.CreateClient(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			res, err := tibcHandler(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		}
	}
}
