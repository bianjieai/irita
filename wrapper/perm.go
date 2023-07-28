package wrapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type PermKeeper struct {
	cdc  codec.Codec
	perm permkeeper.Keeper
}

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

	if err := k.perm.Access(ctx, addr, perm.RoleLayer2User.Auth()); err != nil {
		return false
	}
	return true
}
