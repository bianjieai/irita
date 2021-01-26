package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/irita/modules/opb/types"
)

// ParamKeyTable for the OPB module
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&types.Params{})
}

// BaseTokenDenom returns the base token denom
func (k Keeper) BaseTokenDenom(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyBaseTokenDenom, &res)
	return
}

// PointTokenDenom returns the point token denom
func (k Keeper) PointTokenDenom(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyPointTokenDenom, &res)
	return
}

// BaseTokenManager returns the base token manager
func (k Keeper) BaseTokenManager(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyBaseTokenManager, &res)
	return
}

// UnrestrictedTokenTransfer returns the boolean value which indicates if the token transfer is restricted
func (k Keeper) UnrestrictedTokenTransfer(ctx sdk.Context) (res bool) {
	k.paramSpace.Get(ctx, types.KeyUnrestrictedTokenTransfer, &res)
	return
}

// GetParams gets all parameters
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var p types.Params
	k.paramSpace.GetParamSet(ctx, &p)

	return p
}

// SetParams sets the params to the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
