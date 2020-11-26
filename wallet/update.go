package wallet

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/input"
)

func UpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update root passphrase.",
		RunE:  runUpdateCmd,
	}

	cmd.Flags().String(flagType, TypePassphrase, "A way to reset passphrase, should be 'passphrase' or `mnemonic`")

	_ = viper.BindPFlag(flagType, cmd.Flags().Lookup(flagType))
	return cmd
}

func runUpdateCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())

	kb, err := getKeybase(inBuf)
	if err != nil {
		return err
	}

	if !kb.HasInit() {
		return errors.New("wallet has not been created, please run `iritawallet init` first")
	}

	typ := viper.GetString(flagType)
	if typ != TypePassphrase && typ != TypeMnemonic {
		return errors.New("invalid way to reset passphrase, should be `passphrase` or `mnemonic`")
	}

	var mnemonic string
	if typ == TypeMnemonic {
		mnemonic, err = input.GetString("Enter your mnemonic:", inBuf)
		if err != nil || len(mnemonic) == 0 {
			return errors.New("invalid mnemonic")
		}
	}

	newRootPassphrase, err := input.GetPassword("Enter your new root passphrase:", inBuf)
	if err != nil || len(newRootPassphrase) == 0 {
		return errors.New("invalid new root passphrases")
	}

	//input again
	p2, err := input.GetPassword("Re-enter the new root passphrase:", inBuf)
	if err != nil {
		return err
	}

	if newRootPassphrase != p2 {
		return errors.New("two passphrases inputs do not match")
	}

	if err := kb.UpdateRoot(mnemonic, newRootPassphrase); err != nil {
		return fmt.Errorf("update root passphrase failed:%s", err.Error())
	}
	fmt.Println("update root passphrase success")

	return nil
}
