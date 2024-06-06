package keeper

import (
	"context"

	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/tibc/types"
)

func (k Keeper) CreateClient(ctx context.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return &types.MsgCreateClientResponse{}, err
	}
	consensusState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return &types.MsgCreateClientResponse{}, err
	}
	return &types.MsgCreateClientResponse{}, k.ClientKeeper.CreateClient(
		sdk.UnwrapSDKContext(ctx),
		msg.ChainName,
		clientState,
		consensusState,
	)
}

func (k Keeper) UpgradeClient(ctx context.Context, msg *types.MsgUpgradeClient) (*types.MsgUpgradeClientResponse, error) {
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return &types.MsgUpgradeClientResponse{}, err
	}

	consensusState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return &types.MsgUpgradeClientResponse{}, err
	}

	return &types.MsgUpgradeClientResponse{}, k.ClientKeeper.UpgradeClient(
		sdk.UnwrapSDKContext(ctx),
		msg.ChainName,
		clientState,
		consensusState,
	)
}

func (k Keeper) RegisterRelayer(ctx context.Context, msg *types.MsgRegisterRelayer) (*types.MsgRegisterRelayerResponse, error) {

	k.ClientKeeper.RegisterRelayers(
		sdk.UnwrapSDKContext(ctx),
		msg.ChainName,
		msg.Relayers,
	)
	return &types.MsgRegisterRelayerResponse{}, nil
}

func (k Keeper) SetRoutingRules(ctx context.Context, msg *types.MsgSetRoutingRules) (*types.MsgSetRoutingRulesResponse, error) {

	err := k.RoutingKeeper.SetRoutingRules(
		sdk.UnwrapSDKContext(ctx),
		msg.Rules,
	)
	if err != nil {
		return nil, err
	}
	return &types.MsgSetRoutingRulesResponse{}, nil
}
