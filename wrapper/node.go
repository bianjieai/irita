package wrapper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node"
)

type NodeKeeper struct {
	nk node.Keeper
}

func NewNodeKeeper(nk node.Keeper) NodeKeeper {
	return NodeKeeper{nk}
}

func (k NodeKeeper) GetHistoricalInfo(
	ctx sdk.Context,
	height int64,
) (stakingtypes.HistoricalInfo, bool) {
	return k.nk.GetHistoricalInfo(ctx, height)
}

func (k NodeKeeper) GetValidatorByConsAddr(
	ctx sdk.Context,
	consAddr sdk.ConsAddress,
) (validator stakingtypes.Validator, found bool) {
	addr, found := k.nk.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return stakingtypes.Validator{}, false
	}
	validator.Jailed = addr.Jailed

	_, i, err := bech32.DecodeAndConvert(addr.Operator)
	if err != nil {
		return stakingtypes.Validator{}, false
	}
	validator.OperatorAddress = sdk.ValAddress(i).String()

	return validator, found
}
