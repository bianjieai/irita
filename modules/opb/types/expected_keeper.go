package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	tokentypes "github.com/irisnet/irismod/modules/token/types"
)

// TokenKeeper defines the expected token keeper (noalias)
type TokenKeeper interface {
	GetOwner(ctx sdk.Context, denom string) (sdk.AccAddress, error)
	GetToken(ctx sdk.Context, denom string) (tokentypes.TokenI, error)
	MintToken(ctx sdk.Context, symbol string, amount uint64, recipient sdk.AccAddress, owner sdk.AccAddress) error
}

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
}

// PermKeeper defines the expected perm keeper (noalias)
type PermKeeper interface {
	IsRootAdmin(ctx sdk.Context, address sdk.AccAddress) bool
	IsBaseM1Admin(ctx sdk.Context, address sdk.AccAddress) bool
	IsPlatformUser(ctx sdk.Context, address sdk.AccAddress) bool
}
