package cmd

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/bianjieai/iritamod/modules/validator"

	"github.com/bianjieai/irita/modules/genutil"
)

// Genkey returns a command that generates the key from priv_validator_key.json.
// Will used to generate CA request.
func GenRootCert(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-root-cert [cert]",
		Short: "Add X.509 root certificate to verify the validator certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.JSONMarshaler

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			cert, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			validatorGenState := validator.GetGenesisStateFromAppState(cdc, appState)
			validatorGenState.RootCert = string(cert)

			validatorGenStateBz, err := cdc.MarshalJSON(&validatorGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal admin genesis state: %w", err)
			}

			appState[validator.ModuleName] = validatorGenStateBz

			appStateJSON, err := json.MarshalIndent(appState, "", "  ")
			if err != nil {
				return fmt.Errorf("failed to marshal application genesis state: %w", err)
			}

			genDoc.AppState = appStateJSON
			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")

	return cmd
}
