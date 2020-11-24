package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/bianjieai/iritamod/modules/validator"

	"github.com/bianjieai/irita/modules/genutil"
	cautil "github.com/bianjieai/irita/utils/ca"
)

// ValidatorMsgBuildingHelpers helpers for message building gen-tx command
type ValidatorMsgBuildingHelpers interface {
	CreateValidatorMsgHelpers(ipDefault string) (fs *flag.FlagSet, certFlag, powerFlag, defaultsDesc string)
	PrepareFlagsForTxCreateValidator(config *cfg.Config, nodeID, chainID string, cert string)
	BuildCreateValidatorMsg(cliCtx client.Context, txBldr tx.Factory) (tx.Factory, sdk.Msg, error)
}

// AddGenesisValidatorCmd returns add-genesis-validator cobra Command.
func AddGenesisValidatorCmd(
	mbm module.BasicManager, smbh ValidatorMsgBuildingHelpers, defaultNodeHome, defaultCLIHome string,
) *cobra.Command {
	ipDefault, _ := server.ExternalIP()
	fsCreateValidator, flagCert, _, defaultsDesc := smbh.CreateValidatorMsgHelpers(ipDefault)

	cmd := &cobra.Command{
		Use:   "add-genesis-validator",
		Short: "Generate a genesis tx to create a validator",
		Args:  cobra.NoArgs,
		Long: fmt.Sprintf(
			"This command is an alias of the 'tx validator create' command'.\n\n"+
				"It creates a genesis transaction to create a validator.\n"+
				"The following default parameters are included:\n	%s",
			defaultsDesc,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			depCdc := clientCtx.JSONMarshaler
			cdc := depCdc.(codec.Marshaler)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			nodeID, _, err := genutil.InitializeNodeValidatorFiles(config)
			if err != nil {
				return errors.Wrap(err, "failed to initialize node validator files")
			}

			genDoc, err := tmtypes.GenesisDocFromFile(config.GenesisFile())
			if err != nil {
				return errors.Wrapf(err, "failed to read genesis doc file %s", config.GenesisFile())
			}

			var genesisState map[string]json.RawMessage
			if err = json.Unmarshal(genDoc.AppState, &genesisState); err != nil {
				return errors.Wrap(err, "failed to unmarshal genesis state")
			}

			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, genesisState); err != nil {
				return errors.Wrap(err, "failed to validate genesis state")
			}

			// Set flags for creating gentx
			viper.Set(flags.FlagHome, viper.GetString(flagClientHome))
			smbh.PrepareFlagsForTxCreateValidator(config, nodeID, genDoc.ChainID, viper.GetString(flagCert))

			//txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))
			//cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			// Set the generate-only flag here after the CLI context has
			// been created. This allows the from name/key to be correctly populated.
			//
			// TODO: Consider removing the manual setting of generate-only in
			// favor of a 'gentx' flag in the create-validator command.
			viper.Set(flags.FlagGenerateOnly, true)

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).
				WithTxConfig(clientCtx.TxConfig).
				WithAccountRetriever(clientCtx.AccountRetriever)

			// create a 'create-validator' message
			_, msg, err := smbh.BuildCreateValidatorMsg(clientCtx, txf)
			if err != nil {
				return errors.Wrap(err, "failed to build create-validator message")
			}

			if msg, ok := msg.(*validator.MsgCreateValidator); ok {
				validatorGenState := validator.GetGenesisStateFromAppState(cdc, genesisState)

				cert, err := cautil.ReadCertificateFromMem([]byte(msg.Certificate))
				if err != nil {
					return errors.Wrap(err, "failed to convert certificate")
				}

				rootCert, err := cautil.ReadCertificateFromMem([]byte(validatorGenState.RootCert))
				if err != nil {
					return errors.Wrap(err, "failed to convert root certificate")
				}

				if err = cert.VerifyCertFromRoot(rootCert); err != nil {
					return errors.Wrap(err, "invalid certificate, cannot be verified by root certificate")
				}

				pk, err := cautil.GetPubkeyFromCert(cert)
				if err != nil {
					return err
				}

				operator, err := sdk.AccAddressFromBech32(msg.Operator)
				if err != nil {
					return err
				}

				validatorGenState.Validators = append(
					validatorGenState.Validators,
					validator.NewValidator(
						tmhash.Sum(msg.GetSignBytes()),
						msg.Name, msg.Description, pk,
						msg.Certificate, msg.Power, operator,
					),
				)

				validatorGenStateBz, err := cdc.MarshalJSON(&validatorGenState)
				if err != nil {
					return fmt.Errorf("failed to marshal validator genesis state: %w", err)
				}

				genesisState[validator.ModuleName] = validatorGenStateBz
			}

			if err = mbm.ValidateGenesis(cdc, clientCtx.TxConfig, genesisState); err != nil {
				return errors.Wrap(err, "failed to validate genesis state")
			}

			appState, err := json.MarshalIndent(genesisState, "", "  ")
			if err != nil {
				return err
			}

			genDoc.AppState = appState
			return genutil.ExportGenesisFile(genDoc, config.GenesisFile())
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, defaultCLIHome, "client's home directory")
	cmd.Flags().String(flags.FlagOutputDocument, "", "write the genesis transaction JSON document to the given file instead of the default location")
	cmd.Flags().AddFlagSet(fsCreateValidator)
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	cmd.Flags().String(flags.FlagKeyringBackend, flags.DefaultKeyringBackend, "Select keyring's backend (os|file|test)")
	_ = viper.BindPFlag(flags.FlagKeyringBackend, cmd.Flags().Lookup(flags.FlagKeyringBackend))

	_ = cmd.MarkFlagRequired(flags.FlagName)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
