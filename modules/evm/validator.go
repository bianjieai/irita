package evm

import (
	"fmt"
	"math/big"

	"github.com/tendermint/tendermint/libs/log"

	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
)

var IritaCoefficient = new(big.Int).SetUint64(1000000000000)

const feeDenom = "uirita"

type EthOpbValidator struct {
	opbKeeper   *opbkeeper.Keeper
	tokenKeeper *tokenkeeper.Keeper
	evmKeeper   *evmkeeper.Keeper
	logger      log.Logger
}

func NewEthOpbValidator(opbKeeper *opbkeeper.Keeper, tokenKeeper *tokenkeeper.Keeper, evmKeeper *evmkeeper.Keeper, logger log.Logger) *EthOpbValidator {
	return &EthOpbValidator{
		opbKeeper:   opbKeeper,
		tokenKeeper: tokenKeeper,
		evmKeeper:   evmKeeper,
		logger:      logger,
	}
}

func (ov EthOpbValidator) Transfer(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
	ctx := ov.evmKeeper.Ctx()
	senderCosmosAddr := sdk.AccAddress(sender.Bytes())
	recipientCosmosAddr := sdk.AccAddress(recipient.Bytes())

	params := ov.evmKeeper.GetParams(ctx)
	restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ctx)
	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		owner, err := ov.getOwner(ctx, params.EvmDenom)
		if err != nil {
			//ov.logger.Error("unauthorized operation", "err_msg", err.Error())
			ov.evmKeeper.Logger(ctx).Error(
				"unauthorized operation",
				"err_msg", err.Error(),
				"amount", amount.Int64(),
			)
			return
		}
		if senderCosmosAddr.String() != owner && recipientCosmosAddr.String() != owner {
			errMsg := fmt.Sprintf("either the sender or recipient must be the owner %s for token %s", owner, params.EvmDenom)
			//ov.logger.Error("unauthorized operation", "err_msg", errMsg)
			ov.evmKeeper.Logger(ctx).Error(
				"unauthorized operation",
				"err_msg", errMsg,
				"amount", amount.Int64(),
			)
			return
		}
	}
	// go-ethereum
	amount.Quo(amount, IritaCoefficient)
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

func (ov EthOpbValidator) getOwner(ctx sdk.Context, denom string) (owner string, err error) {
	baseTokenDenom := ov.opbKeeper.BaseTokenDenom(ctx)

	if denom == baseTokenDenom {
		owner = ov.opbKeeper.BaseTokenManager(ctx)
	} else {
		ownerAddr, err := ov.tokenKeeper.GetOwner(ctx, denom)
		if err == nil {
			owner = ownerAddr.String()
		}
	}

	return
}
