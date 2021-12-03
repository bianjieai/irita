package wevm

import (
	wevmkeeper "github.com/bianjieai/iritamod/modules/wevm/keeper"
	"github.com/bianjieai/iritamod/modules/wevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/palantir/stacktrace"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

type EthCanCallDecorator struct {
	WevmKeeper wevmkeeper.Keeper
}

func NewEthCanCallDecorator(wevmKeeper wevmkeeper.Keeper) EthCanCallDecorator {
	return EthCanCallDecorator{WevmKeeper: wevmKeeper}
}

func (e EthCanCallDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for i, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, stacktrace.Propagate(
				sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil)),
				"failed to cast transaction %d", i,
			)
		}
		ethTx := msgEthTx.AsTransaction()
		if ethTx.To() != nil {
			state, _ := e.WevmKeeper.GetContractState(ctx, ethTx.To().String())
			if state {
				return ctx, stacktrace.Propagate(
					sdkerrors.Wrapf(types.ErrContractDisable, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil)),
					"failed to run transaction %d", i,
				)
			}
		}
	}
	return next(ctx, tx, simulate)
}
