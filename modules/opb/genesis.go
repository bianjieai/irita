package opb

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis stores the genesis state
func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) (res []abci.ValidatorUpdate) {
	if err := ValidateGenesis(data); err != nil {
		panic(err.Error())
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
