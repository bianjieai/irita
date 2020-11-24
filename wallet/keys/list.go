package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/bianjieai/irita/wallet/keyring"
)

func ListCmd(generator KeybaseGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all keys.",
		Long:  `Return a list of all public keys stored by this key manager along with their associated name and address.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListCmd(cmd, args, generator)
		},
	}
	return cmd
}

func runListCmd(cmd *cobra.Command, args []string, generator KeybaseGenerator) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())

	kb, err := generator(inBuf)
	if err != nil {
		return err
	}

	infos, err := kb.List()
	if err != nil {
		return err
	}

	cmd.SetOut(cmd.OutOrStdout())
	keyring.PrintInfo(cmd.OutOrStdout(), infos...)

	return nil
}
