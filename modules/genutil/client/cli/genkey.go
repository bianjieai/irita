package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/tempfile"

	"github.com/cosmos/cosmos-sdk/server"

	"github.com/bianjieai/irita/modules/genutil"
)

var (
	FlagOutFile = "out-file"
)

// Genkey returns a command that generates the key from priv_validator_key.json.
// Will used to generate CA request.
func GenKey(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "genkey",
		Short: "generate key from validator key file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			_, filePv, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			privKey := filePv.Key.PrivKey
			key, err := genutil.Genkey(privKey)
			if err != nil {
				return err
			}

			return tempfile.WriteFileAtomic(viper.GetString(FlagOutFile), key, 0600)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")
	cmd.Flags().String(FlagOutFile, "priv.pem", "private key file path")

	return cmd
}
