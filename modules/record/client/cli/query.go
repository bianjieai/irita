package cli

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/bianjieai/irita/modules/record/internal/types"
)

// GetQueryCmd returns the cli query commands for the record module.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the record module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(flags.GetCommands(
		GetCmdQueryRecord(cdc),
	)...)
	return txCmd
}

// GetCmdQueryRecord implements the query record command.
func GetCmdQueryRecord(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "record",
		Short:   "Query a record",
		Example: "iritacli query record [record-id]",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			recordID, err := hex.DecodeString(args[0])
			if err != nil {
				return errors.New("invalid record id, must be hex encoded string")
			}

			params := types.QueryRecordParams{
				RecordID: recordID,
			}

			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRecord), bz)
			if err != nil {
				return err
			}

			var record types.Record
			if err := cdc.UnmarshalJSON(res, &record); err != nil {
				return err
			}

			return cliCtx.PrintOutput(record)
		},
	}
	return cmd
}
