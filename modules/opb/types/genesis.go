package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(params Params, contractAddress []string) *GenesisState {
	return &GenesisState{
		Params:              params,
		ContractDenyAddress: contractAddress,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}
	for _, contract := range data.ContractDenyAddress {
		if !common.IsHexAddress(contract) {
			return sdkerrors.Wrap(ErrInvalidContractAddress, "invalid from address")
		}
	}
	return nil
}
