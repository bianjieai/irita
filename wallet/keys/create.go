package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/bianjieai/irita/wallet/keyring"
)

func CreateCmd(generator KeybaseGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Add an encrypted private key (either newly generated or recovered), encrypt it and save to disk.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateCmd(cmd, args, generator)
		},
	}
	return cmd
}

func runCreateCmd(cmd *cobra.Command, args []string, generator KeybaseGenerator) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	kb, err := generator(inBuf)
	if err != nil {
		return err
	}

	info, err := kb.NewKey(args[0])
	if err != nil {
		return err
	}
	keyring.PrintInfo(cmd.OutOrStdout(), info)

	return nil
}
