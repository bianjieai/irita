package evm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"iritamod.bianjie.ai/modules/node"
)

type stakingKeeper struct {
	node.Keeper
}

// NewStakingKeeper creates a new instance of the stakingKeeper struct.
//
// It takes a `node.Keeper` as a parameter and returns a `evmtypes.StakingKeeper`.
func NewStakingKeeper(k node.Keeper) evmtypes.StakingKeeper {
	return stakingKeeper{k}
}

func (node stakingKeeper) GetHistoricalInfo(ctx sdk.Context, height int64) (stakingtypes.HistoricalInfo, bool) {
	return node.Keeper.GetHistoricalInfo(ctx, height)
}

func (node stakingKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found bool) {
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
