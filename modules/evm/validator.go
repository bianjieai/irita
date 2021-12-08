package evm

import (
	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
)

type EthOpbValidator struct {
	opbKeeper   *opbkeeper.Keeper
	tokenKeeper *tokenkeeper.Keeper
}

func NewEthOpbValidator(opbKeeper *opbkeeper.Keeper, tokenKeeper *tokenkeeper.Keeper) *EthOpbValidator {
	return &EthOpbValidator{
		opbKeeper:   opbKeeper,
		tokenKeeper: tokenKeeper,
	}
}

func (ov EthOpbValidator) Authorization(ctx sdk.Context, denom string, addr string) bool {
	restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ctx)
	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		owner, err := ov.getOwner(ctx, denom)
		if err != nil {
			return false
		}
		if addr != owner {
			return false
		}
	}
	return false
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
