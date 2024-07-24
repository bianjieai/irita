package wrapper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	tokentypes "mods.irisnet.org/modules/token/types"
)

type MockICS20 struct{}

// HasTrace implements types.ICS20Keeper.
func (t *MockICS20) HasTrace(ctx sdk.Context, denom string) bool {
	return false
}

func ProvideICS20Keeper() tokentypes.ICS20Keeper {
	return &MockICS20{}
}
