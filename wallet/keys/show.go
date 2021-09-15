package keys

import (
	"bufio"

	"github.com/spf13/cobra"

	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/wallet/keyring"
)

func ShowCmd(generator KeybaseGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show <name|address>",
		Short: "Retrieve key information by name or address.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShowCmd(cmd, args, generator)
		},
	}
	return cmd
}

func runShowCmd(cmd *cobra.Command, args []string, generator KeybaseGenerator) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())

	kb, err := generator(inBuf)
	if err != nil {
		return err
	}

	var info cosmoskeyring.Info
	if addr, err := sdk.AccAddressFromBech32(args[0]); err == nil {
		if info, err = kb.KeyByAddress(addr); err != nil {
			return err
		}
	} else {
		if info, err = kb.Key(args[0]); err != nil {
			return err
		}
	}
	keyring.PrintInfo(cmd.OutOrStdout(), info)

	return nil
}
