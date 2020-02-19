package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// service module sentinel errors
var (
	ErrUnknownSvcDef            = sdkerrors.Register(ModuleName, 1, "unknown service definition")
	ErrUnknownSvcBinding        = sdkerrors.Register(ModuleName, 2, "unknown service binding")
	ErrInvalidServiceName       = sdkerrors.Register(ModuleName, 3, "invalid service name, must contain alphanumeric characters, _ and - only，length greater than 0 and less than or equal to 128")
	ErrInvalidOutputPrivacyEnum = sdkerrors.Register(ModuleName, 4, "invalid output privacy enum")
	ErrInvalidOutputCachedEnum  = sdkerrors.Register(ModuleName, 5, "invalid output cached enum")
	ErrSvcBindingExists         = sdkerrors.Register(ModuleName, 6, "service binding already exists")
	ErrLtMinProviderDeposit     = sdkerrors.Register(ModuleName, 7, "deposit amount must be equal or greater than min provider deposit")
	ErrUnknownMethod            = sdkerrors.Register(ModuleName, 8, "unknown service method")
	ErrUnavailable              = sdkerrors.Register(ModuleName, 9, "service binding is unavailable")
	ErrAvailable                = sdkerrors.Register(ModuleName, 10, "service binding is available")
	ErrRefundDeposit            = sdkerrors.Register(ModuleName, 11, "can't refund deposit")
	ErrIntOverflow              = sdkerrors.Register(ModuleName, 12, "Int overflow")
	ErrUnknownProfiler          = sdkerrors.Register(ModuleName, 13, "unknown profiler")
	ErrUnknownTrustee           = sdkerrors.Register(ModuleName, 14, "unknown trustee")
	ErrLtServiceFee             = sdkerrors.Register(ModuleName, 15, "service fee amount must be equal or greater than price")
	ErrUnknownActiveRequest     = sdkerrors.Register(ModuleName, 16, "unknown active request")
	ErrNotMatchingProvider      = sdkerrors.Register(ModuleName, 17, "not a matching provider")
	ErrNotMatchingReqChainID    = sdkerrors.Register(ModuleName, 18, "not a matching request chain-id")
	ErrUnknownReturnFee         = sdkerrors.Register(ModuleName, 19, "no service refund fees")
	ErrUnknownWithdrawFee       = sdkerrors.Register(ModuleName, 20, "no service withdraw fees")
	ErrUnknownResponse          = sdkerrors.Register(ModuleName, 21, "unknown response")
	ErrInvalidPriceCount        = sdkerrors.Register(ModuleName, 22, "invalid price count")
)
