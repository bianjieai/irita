package wrapper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
)

type (
	TokenKeeper struct {
		k tokenkeeper.Keeper
	}

	ITokenKeeper interface {
		GetOwner(ctx sdk.Context, denom string) (sdk.AccAddress, error)
		GetSymbol(ctx sdk.Context, denom string) (string, error)
		MintToken(
			ctx sdk.Context,
			symbol string,
			amount uint64,
			recipient sdk.AccAddress,
			owner sdk.AccAddress,
		) error
	}
)

var _ ITokenKeeper = TokenKeeper{}

func NewTokenKeeper(k tokenkeeper.Keeper) TokenKeeper {
	return TokenKeeper{k}
}

func (tk TokenKeeper) GetOwner(ctx sdk.Context, denom string) (sdk.AccAddress, error) {
	return tk.k.GetOwner(ctx, denom)
}

func (tk TokenKeeper) GetSymbol(ctx sdk.Context, denom string) (string, error) {
	token, err := tk.k.GetToken(ctx, denom)
	if err != nil {
		return "", err
	}
	return token.GetSymbol(), nil
}

func (tk TokenKeeper) MintToken(
	ctx sdk.Context,
	symbol string,
	amount uint64,
	recipient sdk.AccAddress,
	owner sdk.AccAddress,
) error {
	return tk.k.MintToken(ctx, symbol, amount, recipient, owner)
}
