package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
)

// GetBroadcastCommand returns the tx broadcast command.
func GetBroadcastCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "broadcast [file_path]",
		Short: "Broadcast transactions generated offline",
		Long: strings.TrimSpace(`Broadcast transactions created with the --generate-only
flag and signed with the sign command. Read a transaction from [file_path] and
broadcast it to a node. If you supply a dash (-) argument in place of an input
filename, the command reads from standard input.
$ <appd> tx broadcast ./mytxn.json
`),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			if offline, _ := cmd.Flags().GetBool(flags.FlagOffline); offline {
				return errors.New("cannot broadcast tx during offline mode")
			}

			stdTx, err := authclient.ReadTxFromFile(clientCtx, args[0])
			if err != nil {
				return err
			}

			txBytes, err := clientCtx.TxConfig.TxEncoder()(stdTx)
			if err != nil {
				return err
			}

			if clientCtx.Simulate {
				txSvcClient := tx.NewServiceClient(clientCtx)
				simRes, err := txSvcClient.Simulate(context.Background(), &tx.SimulateRequest{
					TxBytes: txBytes,
				})

				if err != nil {
					return err
				}

				_, _ = fmt.Fprintf(os.Stderr, "%d\n", simRes.GasInfo.GasUsed)
				return nil
			}

			res, err := clientCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
