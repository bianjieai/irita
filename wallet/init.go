package wallet

import (
	"bufio"
	"errors"
	"io"

	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bianjieai/irita/wallet/keyring"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: `Generate a seed key, and all accounts in the wallet will generate a new private key based on the seed key and "hdpath".`,
		RunE:  runInitCmd,
	}

	cmd.PersistentFlags().String(FlagKeyAlgo, string(hd.Sm2Type), "Signature algorithm for generating keys")
	cmd.Flags().Bool(flagRecover, false, "Provide a seed to recover the wallet")

	_ = viper.BindPFlag(flagRecover, cmd.Flags().Lookup(flagRecover))
	_ = viper.BindPFlag(FlagKeyAlgo, cmd.PersistentFlags().Lookup(FlagKeyAlgo))

	return cmd
}

func getKeybase(buf io.Reader) (keyring.Keystore, error) {
	return keyring.New(viper.GetString(FlagHome), buf)
}

func runInitCmd(cmd *cobra.Command, args []string) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())
	outBuf := bufio.NewWriter(cmd.OutOrStdout())

	algo, err := cosmoskeyring.NewSigningAlgoFromString(viper.GetString(FlagKeyAlgo), keyring.SigningAlgoList)
	if err != nil {
		return err
	}

	kb, err := getKeybase(inBuf)
	if err != nil {
		return err
	}

	if kb.HasInit() {
		return errors.New("wallet already has been created")
	}

	var mnemonic string
	var info cosmoskeyring.Info
	if viper.GetBool(flagRecover) {
		mnemonic, err = input.GetString("Enter your mnemonic:", inBuf)
		if err != nil || len(mnemonic) == 0 {
			return errors.New("invalid mnemonic")
		}
		if info, err = kb.Recover(mnemonic, algo); err != nil {
			return err
		}
	} else {
		if info, err = kb.Init(algo, inBuf, outBuf); err != nil {
			return err
		}
	}

	cmd.PrintErrln()
	keyring.PrintInfo(cmd.OutOrStdout(), info)

	return nil
}
