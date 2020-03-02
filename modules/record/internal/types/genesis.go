package types

// GenesisState - all record state that must be provided at genesis
type GenesisState struct {
	Records []Record `json:"records"`
}

// NewGenesisState constructs a GenesisState
func NewGenesisState(records []Record) GenesisState {
	return GenesisState{
		Records: records,
	}
}

// DefaultGenesisState gets raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}
