package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewQuerier(k IKeeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}
