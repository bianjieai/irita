package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/opb/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the OPB MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Mint(ctx, msg.Amount, recipient, operator); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%v", msg.Amount)),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgMintResponse{}, nil
}

func (m msgServer) Reclaim(goCtx context.Context, msg *types.MsgReclaim) (*types.MsgReclaimResponse, error) {
	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return nil, err
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.Keeper.Reclaim(ctx, msg.Denom, recipient, operator); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeReclaim,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRecipient, msg.Recipient),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Operator),
		),
	})

	return &types.MsgReclaimResponse{}, nil
}

func (m msgServer) AddToContractDenyList(goCtx context.Context, msg *types.MsgAddToContractDenyList) (*types.MsgAddToContractDenyListResponse, error) {
	if !common.IsHexAddress(msg.ContractAddress) {
		return &types.MsgAddToContractDenyListResponse{},
			errors.Wrapf(types.ErrInvalidContractAddress, "contract Address %s is invalid", msg.ContractAddress)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	err := m.Keeper.AddToContractDenyList(ctx, msg.ContractAddress)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeContractAdd, msg.ContractAddress),
		),
	})
	return &types.MsgAddToContractDenyListResponse{}, nil
}

func (m msgServer) RemoveFromContractDenyList(goCtx context.Context, msg *types.MsgRemoveFromContractDenyList) (*types.MsgRemoveFromContractDenyListResponse, error) {
	if !common.IsHexAddress(msg.ContractAddress) {
		return &types.MsgRemoveFromContractDenyListResponse{},
			errors.Wrapf(types.ErrInvalidContractAddress, "contract Address %s is invalid", msg.ContractAddress)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := m.Keeper.RemoveFromContractDenyList(ctx, msg.ContractAddress)
	if err != nil {
		return &types.MsgRemoveFromContractDenyListResponse{}, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.EventTypeContractRemove, msg.ContractAddress),
		),
	})
	return &types.MsgRemoveFromContractDenyListResponse{}, nil
}
