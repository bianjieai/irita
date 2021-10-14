package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/tibc/types"
	clienttypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
)

func (k Keeper) CreateClient(ctx sdk.Context, msg *types.MsgCreateClient) (*types.MsgCreateClientResponse, error) {
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return &types.MsgCreateClientResponse{}, err
	}

	consensusState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return &types.MsgCreateClientResponse{}, err
	}
	return &types.MsgCreateClientResponse{}, k.ClientKeeper.CreateClient(ctx, msg.ChainName, clientState, consensusState)
}

func (k Keeper) UpgradeClient(ctx sdk.Context, msg *types.MsgUpgradeClient) (*types.MsgUpgradeClientResponse, error) {
	clientState, err := clienttypes.UnpackClientState(msg.ClientState)
	if err != nil {
		return &types.MsgUpgradeClientResponse{}, err
	}

	consensusState, err := clienttypes.UnpackConsensusState(msg.ConsensusState)
	if err != nil {
		return &types.MsgUpgradeClientResponse{}, err
	}

	return &types.MsgUpgradeClientResponse{}, k.ClientKeeper.CreateClient(ctx, msg.ChainName, clientState, consensusState)
}
func (k Keeper) RegisterRelayer(ctx sdk.Context, msg *types.MsgRegisterRelayer) (*types.MsgRegisterRelayerResponse, error) {
	k.ClientKeeper.RegisterRelayers(ctx, msg.ChainName, msg.Relayers)
	return &types.MsgRegisterRelayerResponse{}, nil
}
