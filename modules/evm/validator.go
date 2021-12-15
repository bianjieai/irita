package evm

import (
	"math/big"

	evmkeeper "github.com/tharsis/ethermint/x/evm/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
)

type EthOpbValidator struct {
	opbKeeper   *opbkeeper.Keeper
	tokenKeeper *tokenkeeper.Keeper
	evmKeeper   *evmkeeper.Keeper
}

func NewEthOpbValidator(opbKeeper *opbkeeper.Keeper, tokenKeeper *tokenkeeper.Keeper, evmKeeper *evmkeeper.Keeper) *EthOpbValidator {
	return &EthOpbValidator{
		opbKeeper:   opbKeeper,
		tokenKeeper: tokenKeeper,
		evmKeeper:   evmKeeper,
	}
}

func (ov EthOpbValidator) CanTransfer(db vm.StateDB, addr common.Address, amount *big.Int) bool {
	ctx := ov.evmKeeper.Ctx()
	cosmosAddr := sdk.AccAddress(addr.Bytes())
	params := ov.evmKeeper.GetParams(ctx)
	restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ctx)
	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		owner, err := ov.getOwner(ctx, params.EvmDenom)
		if err != nil {
			return false
		}
		if cosmosAddr.String() != owner {
			return false
		}
	}
	return db.GetBalance(addr).Cmp(amount) >= 0
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
