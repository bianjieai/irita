package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 2, "invalid amount")
	ErrInvalidDenom  = sdkerrors.Register(ModuleName, 3, "invalid denom")
	ErrUnauthorized  = sdkerrors.Register(ModuleName, 4, "unauthorized operation")
)
