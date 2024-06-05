package ante

import (
	evmmoduleante "github.com/bianjieai/irita/modules/evm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	ethermintante "github.com/tharsis/ethermint/app/ante"
)

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethermintante.NewEthSetUpContextDecorator(options.EvmKeeper), // outermost AnteDecorator. SetUpContext must be called first

		evmmoduleante.NewEthSigVerificationDecorator(options.EvmKeeper, options.AccountKeeper, options.SignModeHandler),

		ethermintante.NewEthMempoolFeeDecorator(options.EvmKeeper), // Check eth effective gas price against minimal-gas-prices
		ethermintante.NewEthValidateBasicDecorator(options.EvmKeeper),
		ethermintante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.BankKeeper, options.EvmKeeper),
		ethermintante.NewEthGasConsumeDecorator(options.EvmKeeper),
		ethermintante.NewCanTransferDecorator(options.EvmKeeper),
		ethermintante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.

	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		tokenkeeper.NewValidateTokenFeeDecorator(options.TokenKeeper, options.BankKeeper),
	)
}

func newCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethermintante.RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),

		// perm check
		evmmoduleante.NewEthSigVerificationDecorator(options.EvmKeeper, options.AccountKeeper, options.SignModeHandler),

		// NOTE: extensions option decorator removed
		// ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		ethermintante.NewEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	)

}
