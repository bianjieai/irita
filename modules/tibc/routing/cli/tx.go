package cli

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	tibctypes "github.com/bianjieai/irita/modules/tibc/types"
)

// NewSetRoutingRulesCmd implements a command handler for submitting a setting rules proposal transaction.
func NewSetRoutingRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-rules [path/to/routing_rules.json] [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "set routing rules",
		Long:  "set routing rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			routingRulesBz, err := ioutil.ReadFile(args[0])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for routing rules were provided")
			}

			var rules []string
			if err := json.Unmarshal(routingRulesBz, &rules); err != nil {
				return errors.Wrap(err, "error unmarshalling rules file")
			}

			msg, err := tibctypes.NewMsgSetRoutingRules(rules, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
