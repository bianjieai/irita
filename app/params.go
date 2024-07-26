package app

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
)

const (
	flagAuthority    = "authority"
	defaultAuthority = "gov"
)

// AddStartFlags defines flags for the start of the application
func AddStartFlags(startCmd *cobra.Command) {
	authority := authtypes.NewModuleAddress(defaultAuthority).String()
	startCmd.Flags().String(flagAuthority, authority, "authority of the module")
}
