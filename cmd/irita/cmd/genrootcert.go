package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/bianjieai/iritamod/modules/genutil"
	"github.com/bianjieai/iritamod/modules/node"
)

// GenRootCert returns a command that sets the root cert.
func GenRootCert(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-root-cert [cert]",
		Short: "Add X.509 root certificate to verify the validator or node certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			cdc := clientCtx.Codec

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			cert, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			genFile := config.GenesisFile()
			appState, genDoc, err := genutil.GenesisStateFromGenFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to unmarshal genesis state: %w", err)
			}

			nodeGenState := node.GetGenesisStateFromAppState(cdc, appState)
			nodeGenState.RootCert = string(cert)

			nodeGenStateBz, err := cdc.MarshalJSON(&nodeGenState)
			if err != nil {
				return fmt.Errorf("failed to marshal node genesis state: %w", err)
			}

			appState[node.ModuleName] = nodeGenStateBz

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
