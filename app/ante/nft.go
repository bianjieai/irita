package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type MultipleMsgValidatorDecorator struct {
	ak authkeeper.AccountKeeper
}

// NewMultipleMsgValidatorDecorator valid nft msg
func NewMultipleMsgValidatorDecorator(AccountKeeper authkeeper.AccountKeeper) MultipleMsgValidatorDecorator {
	return MultipleMsgValidatorDecorator{
		ak: AccountKeeper,
	}
}

func (mmvd MultipleMsgValidatorDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgLen := uint64(len(tx.GetMsgs()))
	params := mmvd.ak.GetParams(ctx)
	var gasAmount uint64
	if msgLen > 1 && msgLen <= 1000 {
		gasAmount = params.TxSizeCostPerByte * sdk.Gas(len(ctx.TxBytes())) * msgLen * 9 / 10
	} else if msgLen > 1000 {
		gasAmount = params.TxSizeCostPerByte * sdk.Gas(len(ctx.TxBytes())) * msgLen * 11 / 10
	}

	ctx.GasMeter().ConsumeGas(gasAmount, "txSize")

	return next(ctx, tx, simulate)
}
