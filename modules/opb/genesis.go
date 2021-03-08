package opb

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) (res []abci.ValidatorUpdate) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
	}

	params := data.Params

	if !k.HasToken(ctx, params.BaseTokenDenom) {
		panic(fmt.Sprintf("token %s does not exist", params.BaseTokenDenom))
	}

	if !k.HasToken(ctx, params.PointTokenDenom) {
		panic(fmt.Sprintf("token %s does not exist", params.PointTokenDenom))
	}

	if !params.UnrestrictedTokenTransfer && len(params.BaseTokenManager) == 0 {
		panic("base token manager must be specified when the token transfer restriction enabled")
	}

	k.SetParams(ctx, data.Params)

	return nil
}

// ExportGenesis outputs the genesis state
func ExportGenesis(ctx sdk.Context, k Keeper) *GenesisState {
	return NewGenesisState(
		k.GetParams(ctx),
	)
}
