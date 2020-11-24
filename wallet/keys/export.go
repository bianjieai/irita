package keys

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagOutput = "output"
)

func ExportCmd(generator KeybaseGenerator) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export <name|address>",
		Short: "Export private keys.",
		Long:  `Export a private key from the local keybase in ASCII-armored encrypted format.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runExportCmd(cmd, args, generator)
		},
	}
	cmd.Flags().StringP(flagOutput, "o", "", "the file name of the private key exported")

	_ = viper.BindPFlag(flagOutput, cmd.Flags().Lookup(flagOutput))

	return cmd
}

func runExportCmd(cmd *cobra.Command, args []string, generator KeybaseGenerator) error {
	inBuf := bufio.NewReader(cmd.InOrStdin())

	kb, err := generator(inBuf)
	if err != nil {
		return err
	}

	armor, passphrase, err := kb.Export(args[0])
	if err != nil {
		return err
	}

	var filename = viper.GetString(flagOutput)
	if len(filename) == 0 {
		filename = fmt.Sprintf("%s.key", args[0])
		filename = path.Join("./", filename)
	}

	if err = ioutil.WriteFile(filename, []byte(armor), 0444); err != nil {
		return err
	}

	_, _ = fmt.Fprintf(
		cmd.OutOrStdout(),
		"The private key file was saved in [%s], encrypted password: %s\n",
		filename, passphrase,
	)

	return nil
}
