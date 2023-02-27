package evm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bianjieai/iritamod/modules/node"
)

type WNodeKeeper struct {
	node.Keeper
}

func (node WNodeKeeper) GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool) {
	return node.Keeper.GetHistoricalInfo(ctx, height)
}

func (node WNodeKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
	addr, found := node.Keeper.GetValidatorByConsAddr(ctx, consAddr)
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
