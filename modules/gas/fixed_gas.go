package gas

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
)

var (
	_ sdk.GasMeter = &FixedGasMeter{}

	DefaultGasConfig = map[string]uint64{
		sdk.MsgTypeURL(&nfttypes.MsgIssueDenom{}): 400000,
		sdk.MsgTypeURL(&nfttypes.MsgMintNFT{}):    400000,
		sdk.MsgTypeURL(&mttypes.MsgIssueDenom{}):  400000,
		sdk.MsgTypeURL(&mttypes.MsgMintMT{}):      400000,
	}
	DefaultGas = uint64(200000)
)

type FixedGasMeter struct {
	gasMeter  sdk.GasMeter
	gasConfig map[string]uint64
}

func NewFixedGasMeter(limit sdk.Gas, simulate bool) sdk.GasMeter {
	if simulate {
		limit = math.MaxUint64
	}
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

// GasConsumedToLimit implements types.GasMeter
func (fgm *FixedGasMeter) GasConsumedToLimit() uint64 {
	return fgm.gasMeter.GasConsumedToLimit()
}

// IsOutOfGas implements types.GasMeter
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
