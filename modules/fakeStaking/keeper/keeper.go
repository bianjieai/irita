package keeper

import (
	"github.com/bianjieai/iritamod/modules/node"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type FakeStakingKeeper struct {
	node.Keeper
}

func (node FakeStakingKeeper) GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool) {
	return node.Keeper.GetHistoricalInfo(ctx, height)
}

func (node FakeStakingKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	addr, f := node.Keeper.GetValidatorByConsAddr(ctx, consAddr)
	validator.Jailed = addr.Jailed

	_, i, err := bech32.DecodeAndConvert(addr.Operator)
	if err != nil {
		return stakingtypes.Validator{}, false
	}
	validator.OperatorAddress = sdk.ValAddress(i).String()

	return validator, f
}
