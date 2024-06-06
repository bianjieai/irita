package cli

import (
	"io/ioutil"
	"strings"

	"github.com/bianjieai/tibc-go/modules/tibc/core/02-client/client/cli"
	packet "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet"
	"github.com/bianjieai/tibc-go/modules/tibc/core/exported"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	tibcrouting "github.com/bianjieai/irita/modules/tibc/routing/cli"
	"github.com/bianjieai/irita/modules/tibc/types"
)

func NewTxCmd() *cobra.Command {
	tibcTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "TIBC transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tibcClientCmd := &cobra.Command{
		Use:                        "client",
		Short:                      "TIBC client subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	tibcRoutingCmd := &cobra.Command{
		Use:                        "routing",
		Short:                      "TIBC routing subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	tibcRoutingCmd.AddCommand(
		tibcrouting.NewSetRoutingRulesCmd(),
	)
	tibcClientCmd.AddCommand(
		NewCreateClientCmd(),
		NewUpgradeClient(),
		NewRegisterRelayer(),
		cli.NewUpdateClientCmd(),
	)
	tibcTxCmd.AddCommand(
		packet.GetTxCmd(),
		tibcClientCmd,
		tibcRoutingCmd,
	)

	return tibcTxCmd
}

// NewCreateClientCmd implements the CreateClient command.
func NewCreateClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client-create [chain-name] [path/to/client_state.json] [path/to/consensus_state.json] [flags]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit a client create proposal",
		Long:  "create a new TIBC client with the specified client state and consensus state",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			chainName := args[0]

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			// attempt to unmarshal client state argument
			var clientState exported.ClientState
			clientStateBz, err := ioutil.ReadFile(args[1])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for client state were provided")
			}

			if err := cdc.UnmarshalInterfaceJSON(clientStateBz, &clientState); err != nil {
				return errors.Wrap(err, "error unmarshalling client state file")
			}

			var consensusState exported.ConsensusState
			consensusStateBz, err := ioutil.ReadFile(args[2])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for consensus state were provided")
			}

			if err := cdc.UnmarshalInterfaceJSON(consensusStateBz, &consensusState); err != nil {
				return errors.Wrap(err, "error unmarshalling consensus state file")
			}
			msg, err := types.NewMsgCreateClient(chainName, clientState, consensusState, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewUpgradeClient implements the UpgradeClient command.
func NewUpgradeClient() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client-upgrade [chain-name] [path/to/client_state.json] [path/to/consensus_state.json] [flags]",
		Args:  cobra.ExactArgs(3),
		Short: "Submit a client upgrade proposal",
		Long:  "upgrade a TIBC client with the specified client state and consensus state",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			chainName := args[0]

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)
			// attempt to unmarshal client state argument
			var clientState exported.ClientState
			clientStateBz, err := ioutil.ReadFile(args[1])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for client state were provided")
			}

			if err := cdc.UnmarshalInterfaceJSON(clientStateBz, &clientState); err != nil {
				return errors.Wrap(err, "error unmarshalling client state file")
			}

			var consensusState exported.ConsensusState
			consensusStateBz, err := ioutil.ReadFile(args[2])
			if err != nil {
				return errors.Wrap(err, "neither JSON input nor path to .json file for consensus state were provided")
			}

			if err := cdc.UnmarshalInterfaceJSON(consensusStateBz, &consensusState); err != nil {
				return errors.Wrap(err, "error unmarshalling consensus state file")
			}
			msg, err := types.NewMsgUpgradeClient(chainName, clientState, consensusState, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewRegisterRelayer implements the RegisterRelayer command.
func NewRegisterRelayer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "relayer-register [chain-name] [relayers-address] [flags]",
		Args:  cobra.ExactArgs(2),
		Short: "Submit a relayer register proposal",
		Long:  "Submit a relayer register proposal for the specified client",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			chainName := args[0]
			relayers := strings.Split(args[1], ",")
			msg, err := types.NewMsgRegisterRelayer(chainName, relayers, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
