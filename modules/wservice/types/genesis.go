package types

func NewGenesisState(sequence []RequestSequence) *GenesisState {
	return &GenesisState{
		ReqSequence: sequence,
	}
}

// ValidateGenesis validates the provided genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	//todo
	return nil
}
