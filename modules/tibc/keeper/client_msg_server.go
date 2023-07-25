package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	packettypes "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
	corekeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

type clientMsgServer struct {
	k Keeper
}

// ClientMsgServer return a client MsgServer for overiding tibc client MsgServer
func (k Keeper) ClientMsgServer() clientMsgServer {
	return clientMsgServer{k}
}

// PacketMsgServer return a packet MsgServer
func (k Keeper) PacketMsgServer() packettypes.MsgServer {
	return corekeeper.NewMsgServerImpl(*k.Keeper)
}

// CreateClient should not be implemented
func (cms clientMsgServer) CreateClient(
	goCtx context.Context,
	req *clienttypes.MsgCreateClient,
) (*clienttypes.MsgCreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}

// UpdateClient implement the method UpdateClient of the tibc client MsgServer
func (cms clientMsgServer) UpdateClient(
	goCtx context.Context,
	req *clienttypes.MsgUpdateClient,
) (*clienttypes.MsgUpdateClientResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	header, err := clienttypes.UnpackHeader(req.Header)
	if err != nil {
		return nil, err
	}

	// Verify that the account has permission to update the client
	if !cms.k.ClientKeeper.AuthRelayer(ctx, req.ChainName, req.Signer) {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"relayer: %s",
			req.Signer,
		)
	}

	if err = cms.k.ClientKeeper.UpdateClient(ctx, req.ChainName, header); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, clienttypes.AttributeValueCategory),
		),
	)

	return &clienttypes.MsgUpdateClientResponse{}, nil
}

// UpgradeClient should not be implemented
func (cms clientMsgServer) UpgradeClient(
	goCtx context.Context,
	req *clienttypes.MsgUpgradeClient,
) (*clienttypes.MsgUpgradeClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpgradeClient not implemented")
}

// RegisterRelayer should not be implemented
func (cms clientMsgServer) RegisterRelayer(
	goCtx context.Context,
	req *clienttypes.MsgRegisterRelayer,
) (*clienttypes.MsgRegisterRelayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRelayer not implemented")
}
