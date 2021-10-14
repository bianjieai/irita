package tibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/bianjieai/irita/modules/tibc/types"
	packettypes "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	"github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

// NewHandler defines the TIBC handler
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *clienttypes.MsgCreateClient:

			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized TIBC message type: %T", msg)
		}
	}
}
