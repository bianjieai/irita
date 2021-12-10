package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/bianjieai/irita/modules/opb/types"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd() *cobra.Command {
	opbQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the OPB module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	ContractDenyListQueryCmd := &cobra.Command{
		Use:                        types.ContractDenyListName,
		Short:                      "Querying commands for the contract-deny-list",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	AccountDenyListQueryCmd := &cobra.Command{
		Use:                        types.AccountDenyListName,
		Short:                      "Querying commands for the account-deny-list",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	ContractDenyListQueryCmd.AddCommand(
		NewGetContractDenyList(),
		NewGetContractState(),
	)
	AccountDenyListQueryCmd.AddCommand(
		NewGetAccountDenyList(),
		NewGetAccountState(),
	)
	opbQueryCmd.AddCommand(
		GetCmdQueryParams(),
		ContractDenyListQueryCmd,
		AccountDenyListQueryCmd,
	)

	return opbQueryCmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Query the OPB parameters",
		Long:    "Query the current OPB parameter set",
		Example: fmt.Sprintf("$ %s query %s params", version.AppName, types.ModuleName),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func NewGetContractDenyList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny-list  [flags]",
		Short: "get contract deny list state",
		Long:  "get contract deny list state",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryContractDenyListRequest{}
			res, err := queryClient.ContractDenyList(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
func NewGetContractState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract [contractAddress] [flags]",
		Short: "get contract state",
		Long:  "get contract state by contract address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			contractAddr := args[0]
			req := &types.QueryContractStateRequest{
				Address: contractAddr,
			}
			res, err := queryClient.ContractState(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func NewGetAccountDenyList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deny-list  [flags]",
		Short: "get all account deny list",
		Long:  "get all account deny list",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAccountDenyListRequest{}
			res, err := queryClient.AccountDenyList(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
func NewGetAccountState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account [accountAddress] [flags]",
		Short: "get account state",
		Long:  "get account state by account address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			addr := args[0]
			req := &types.QueryAccountStateRequest{
				Address: addr,
			}
			res, err := queryClient.AccountState(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
