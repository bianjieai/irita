package record

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

// NewHandler returns a handler for all "guardian" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgCreateRecord:
			return handleMsgCreateRecord(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

// handleMsgCreateRecord handles MsgCreateRecord
func handleMsgCreateRecord(ctx sdk.Context, k Keeper, msg MsgCreateRecord) (*sdk.Result, error) {
	record := NewRecord(tmhash.Sum(ctx.TxBytes()), msg.Contents, msg.Creator)
	recordId := k.AddRecord(ctx, record)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
		sdk.NewEvent(
			EventTypeCreateRecord,
			sdk.NewAttribute(AttributeKeyCreator, msg.Creator.String()),
			sdk.NewAttribute(AttributeKeyRecordID, hex.EncodeToString(recordId)),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
