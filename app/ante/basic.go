package ante

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
)

// GasTx defines a Tx with a GetGas() method which is needed to use SetUpContextDecorator
type GasTx interface {
	sdk.Tx
	GetGas() uint64
}

type SetUpContextDecorator struct{}

func NewSetUpContextDecorator() SetUpContextDecorator {
	return SetUpContextDecorator{}
}

func (sucd SetUpContextDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	gasTx, ok := tx.(GasTx)
	if !ok {
		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
		// during runTx.
		newCtx = SetGasMeter(simulate, ctx, 0, false)
		return newCtx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be GasTx")
	}

	newCtx = SetGasMeter(simulate, ctx, gasTx.GetGas(), true)

	// Decorator will catch an OutOfGasPanic caused in the next antehandler
	// AnteHandlers must have their own defer/recover in order for the BaseApp
	// to know how much gas was used! This is because the GasMeter is created in
	// the AnteHandler, but if it panics the context won't be set properly in
	// runTx's recover call.
	defer func() {
		if r := recover(); r != nil {
			switch rType := r.(type) {
			case sdk.ErrorOutOfGas:
				log := fmt.Sprintf(
					"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
					rType.Descriptor, gasTx.GetGas(), newCtx.GasMeter().GasConsumed())

				err = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, log)
			default:
				panic(r)
			}
		}
	}()

	gasMeter, ok := newCtx.GasMeter().(*FixedGasMeter)
	if !ok {
		return next(newCtx, tx, simulate)
	}
	gasMeter.ConsumeGasWithMsgs(tx.GetMsgs())
	return next(newCtx, tx, simulate)
}

// SetGasMeter returns a new context with a gas meter set from a given context.
func SetGasMeter(simulate bool, ctx sdk.Context, gasLimit uint64, fixedGas bool) sdk.Context {
	// In various cases such as simulation and during the genesis block, we do not
	// meter any gas utilization.
	if ctx.BlockHeight() == 0 {
		return ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	}

	if !fixedGas {
		return ctx.WithGasMeter(sdk.NewGasMeter(gasLimit))
	}

	if simulate {
		gasLimit = DefaultSimulateGas
	}

	return ctx.WithGasMeter(NewFixedGasMeter(gasLimit))
}

var (
	_ sdk.GasMeter = &FixedGasMeter{}

	DefaultGasConfig = map[string]uint64{
		sdk.MsgTypeURL(&nfttypes.MsgIssueDenom{}): 400000,
		sdk.MsgTypeURL(&nfttypes.MsgMintNFT{}):    400000,
		sdk.MsgTypeURL(&mttypes.MsgIssueDenom{}):  400000,
		sdk.MsgTypeURL(&mttypes.MsgMintMT{}):      400000,
	}
	DefaultGas         = uint64(200000)
	DefaultSimulateGas = uint64(500000)
)

type FixedGasMeter struct {
	gasMeter  sdk.GasMeter
	gasConfig map[string]uint64
}

func NewFixedGasMeter(limit sdk.Gas) sdk.GasMeter {
	return &FixedGasMeter{
		gasMeter:  sdk.NewGasMeter(limit),
		gasConfig: DefaultGasConfig,
	}
}

// ConsumeGas implements types.GasMeter
func (fgm *FixedGasMeter) ConsumeGas(amount uint64, descriptor string) {
	//do nothing,overwrite sdk.GasMeter.ConsumeGas()
}

func (g *FixedGasMeter) ConsumeGasWithMsgs(msgs []sdk.Msg) {
	totalGas := uint64(0)
	for _, msg := range msgs {
		gasNeed, ok := g.gasConfig[sdk.MsgTypeURL(msg)]
		if !ok {
			gasNeed = DefaultGas
		}
		totalGas += gasNeed
	}
	g.gasMeter.ConsumeGas(totalGas, "tx_msg_consume")
}

// GasConsumed implements types.GasMeter
func (fgm *FixedGasMeter) GasConsumed() uint64 {
	return fgm.gasMeter.GasConsumed()
}

// GasRemaining returns the gas left in the GasMeter.
func (fgm *FixedGasMeter) GasRemaining() uint64 {
	if fgm.IsPastLimit() {
		return 0
	}
	return fgm.Limit() - fgm.gasMeter.GasConsumed()
}

// GasConsumedToLimit implements types.GasMeter
func (fgm *FixedGasMeter) GasConsumedToLimit() uint64 {
	return fgm.gasMeter.GasConsumedToLimit()
}

// IsOutOfGas d implements types.GasMeter
func (fgm *FixedGasMeter) IsOutOfGas() bool {
	return fgm.gasMeter.IsOutOfGas()
}

// IsPastLimit implements types.GasMeter
func (fgm *FixedGasMeter) IsPastLimit() bool {
	return fgm.gasMeter.IsPastLimit()
}

// Limit implements types.GasMeter
func (fgm *FixedGasMeter) Limit() uint64 {
	return fgm.gasMeter.Limit()
}

// RefundGas implements types.GasMeter
func (fgm *FixedGasMeter) RefundGas(amount uint64, descriptor string) {
	fgm.gasMeter.RefundGas(amount, descriptor)
}

// String implements types.GasMeter
func (fgm *FixedGasMeter) String() string {
	return fgm.gasMeter.String()
}
