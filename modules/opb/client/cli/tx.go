package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/irita/modules/opb/types"
)

func NewTxCmd() *cobra.Command {
	opbTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "OPB transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	opbTxCmd.AddCommand(
		NewMintCmd(),
		NewReclaimCmd(),
	)

	return opbTxCmd
}

// NewMintCmd implements the mint command.
func NewMintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount] [to]",
		Short: "Mint the base native token",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Mint the base native token by the given amount in main unit",
		)),
		Example: fmt.Sprintf(
			"$ %s tx %s mint <amount> <to> --from mykey",
			version.AppName, types.ModuleName,
		),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			amount, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			var recipient sdk.AccAddress

			if len(args) > 1 {
				recipient, err = sdk.AccAddressFromBech32(args[1])
				if err != nil {
					return err
				}
			} else {
				recipient = operator
			}

			msg := types.NewMsgMint(amount, recipient, operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewReclaimCmd implements the reclaim command.
func NewReclaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reclaim [denom] [to]",
		Short: "Reclaim the native token of the specified denom",
		Long: strings.TrimSpace(fmt.Sprintf(
			"Reclaim the native token of the specified denom from the corresponding escrow account",
		)),
		Example: fmt.Sprintf(
			"$ %s tx %s reclaim <denom> <to> --from mykey",
			version.AppName, types.ModuleName,
		),
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			operator := clientCtx.GetFromAddress()

			var recipient sdk.AccAddress

			if len(args) > 1 {
				recipient, err = sdk.AccAddressFromBech32(args[1])
				if err != nil {
					return err
				}
			} else {
				recipient = operator
			}

			msg := types.NewMsgReclaim(args[0], recipient, operator)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
