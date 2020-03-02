package record

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis stores genesis data
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, record := range data.Records {
		keeper.AddRecord(ctx, record)
	}
}

// ExportGenesis outputs genesis data
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	recordsIterator := k.RecordsIterator(ctx)
	defer recordsIterator.Close()

	var records []Record
	for ; recordsIterator.Valid(); recordsIterator.Next() {
		var record Record
		ModuleCdc.MustUnmarshalBinaryLengthPrefixed(recordsIterator.Value(), &record)
		records = append(records, record)
	}

	return NewGenesisState(records)
}
