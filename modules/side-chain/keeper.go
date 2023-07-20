package sidechain

import (
	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewPermKeeper(cdc codec.Codec, pk permkeeper.Keeper) PermKeeper {
	return PermKeeper{
		cdc:  cdc,
		perm: pk,
	}
}

func (k PermKeeper) HasSideChainUserRole(ctx sdk.Context, addr sdk.AccAddress) bool {
	if k.perm.IsRootAdmin(ctx, addr) {
		return true
	}

	if err := k.perm.Access(ctx, addr, perm.RoleSideChainUser.Auth()); err != nil {
		return false
	}
	return true
}
