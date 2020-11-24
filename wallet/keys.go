package wallet

import (
	"github.com/spf13/cobra"

	"github.com/bianjieai/irita/wallet/keys"
)

func KeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Manage all keys in the wallet.",
	}

	cmd.AddCommand(
		keys.CreateCmd(getKeybase),
		keys.ShowCmd(getKeybase),
		keys.ListCmd(getKeybase),
		keys.ExportCmd(getKeybase),
	)
	return cmd
}
