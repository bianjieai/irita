package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/opb/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) ContractState(goCtx context.Context, request *types.QueryContractStateRequest) (*types.QueryContractStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	state, err := k.GetContractState(ctx, request.GetAddress())
	if err != nil {
		return nil, err
	}
	return &types.QueryContractStateResponse{Exist: state}, nil
}

func (k Keeper) ContractDenyList(goCtx context.Context, request *types.QueryContractDenyListRequest) (*types.QueryContractDenyListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	list, err := k.IteratorContractDanyList(ctx)
	if err != nil {
		return nil, err
	}
	return &types.QueryContractDenyListResponse{ContractAddress: list}, nil
}
