package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrDuplicatedRequestSequence = sdkerrors.Register(ModuleName, 2, "request sequence exist")
	ErrTimeoutRequestSequence    = sdkerrors.Register(ModuleName, 3, "request sequence timeout")
	ErrInvalidInput              = sdkerrors.Register(ModuleName, 4, "invalid input")
	ErrInvalidOutput             = sdkerrors.Register(ModuleName, 5, "invalid output")
	ErrInvalidProvider           = sdkerrors.Register(ModuleName, 6, "invalid provider")
	ErrInvalidProviderAddr       = sdkerrors.Register(ModuleName, 7, "invalid provider addr")
	ErrInvalidConsumerAddr       = sdkerrors.Register(ModuleName, 8, "invalid consumer addr")
	ErrInvalidProviderLength     = sdkerrors.Register(ModuleName, 9, "invalid providers")
	ErrRequestNotFound           = sdkerrors.Register(ModuleName, 10, "request not found")
)
