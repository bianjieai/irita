package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irismod/service"
	"github.com/irismod/token"
)

type TokenAdapter struct {
	TokenKeeper token.Keeper
}

func NewTokenAdapter(tokenKeeper token.Keeper) TokenAdapter {
	return TokenAdapter{
		TokenKeeper: tokenKeeper,
	}
}

func (tk TokenAdapter) GetToken(ctx sdk.Context, denom string) (service.TokenI, error) {
	return tk.TokenKeeper.GetToken(ctx, denom)
}
