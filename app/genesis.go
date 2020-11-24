package app

import (
	"encoding/json"
)

// GenesisState defines a type alias for the Iris genesis application state.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	encCfg := MakeEncodingConfig()
	return ModuleBasics.DefaultGenesis(encCfg.Marshaler)
}
