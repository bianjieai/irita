package wevm

import (
	wevmkeeper "github.com/bianjieai/iritamod/modules/wevm/keeper"
	"github.com/bianjieai/iritamod/modules/wevm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

type EthCanCallDecorator struct {
	WevmKeeper wevmkeeper.Keeper
}

func NewEthCanCallDecorator(wevmKeeper wevmkeeper.Keeper) EthCanCallDecorator {
	return EthCanCallDecorator{WevmKeeper: wevmKeeper}
}

func (e EthCanCallDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
		}
		ethTx := msgEthTx.AsTransaction()
		if ethTx.To() != nil {
			state, _ := e.WevmKeeper.GetContractState(ctx, ethTx.To().String())
			if state {
				return ctx, sdkerrors.Wrapf(types.ErrContractDisable, "the contract %s is in contract deny list ! ", ethTx.To())
			}
		}
	}
	return next(ctx, tx, simulate)
}
