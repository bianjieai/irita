package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 2, "invalid amount")
	ErrInvalidDenom  = sdkerrors.Register(ModuleName, 3, "invalid denom")
	ErrUnauthorized  = sdkerrors.Register(ModuleName, 4, "unauthorized operation")
	// ErrInvalidContractAddress returns an error that the contract address is invalid
	ErrInvalidContractAddress = sdkerrors.Register(ModuleName, 5, "contract address is invalid")
	// ErrNotFound returns an error when not found contract from to deny list
	ErrNotFound = sdkerrors.Register(ModuleName, 6, "not found")
	// ErrContractAlreadyExist returns an error that the contract is already in ContractDenyList
	ErrContractAlreadyExist = sdkerrors.Register(ModuleName, 7, "contract already exist")
	// ErrAccountAlreadyExist returns an error that the account is already in AccountDenyList
	ErrAccountAlreadyExist = sdkerrors.Register(ModuleName, 8, "account already exist")
	ErrContractDisable     = sdkerrors.Register(ModuleName, 9, "contract is disable")
	ErrAccountDisable      = sdkerrors.Register(ModuleName, 10, "account is disable")
)
