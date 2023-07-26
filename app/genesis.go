package app

import (
	"encoding/json"
)

// GenesisState defines a type alias for the Iris genesis application state.
type GenesisState map[string]json.RawMessage
