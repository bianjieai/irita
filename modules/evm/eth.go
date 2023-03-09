package evm

import (
	"errors"
	"math/big"

	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	permtypes "github.com/bianjieai/iritamod/modules/perm/types"

	"github.com/tharsis/ethermint/crypto/ethsecp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	evmcrypto "github.com/bianjieai/irita/modules/evm/crypto"

	ethermint "github.com/tharsis/ethermint/types"
	"github.com/tharsis/ethermint/x/evm/statedb"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// EthSigVerificationDecorator validates an ethereum signatures
type EthSigVerificationDecorator struct {
	accountKeeper   authante.AccountKeeper
	signModeHandler authsigning.SignModeHandler
	evmKeeper       EVMKeeper
}

// NewEthSigVerificationDecorator creates a new EthSigVerificationDecorator
func NewEthSigVerificationDecorator(ek EVMKeeper, ak authante.AccountKeeper, signModeHandler authsigning.SignModeHandler) EthSigVerificationDecorator {
	return EthSigVerificationDecorator{
		evmKeeper:       ek,
		accountKeeper:   ak,
		signModeHandler: signModeHandler,
	}
}

// AnteHandle validates checks that the registered chain id is the same as the one on the message, and
// that the signer address matches the one defined on the message.
// It's not skipped for RecheckTx, because it set `From` address which is critical from other ante handler to work.
// Failure in RecheckTx will prevent tx to be included into block, especially when CheckTx succeed, in which case user
// won't see the error message.
func (esvd EthSigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	chainID := esvd.evmKeeper.ChainID()

	params := esvd.evmKeeper.GetParams(ctx)

	ethCfg := params.ChainConfig.EthereumConfig(chainID)
	blockNum := big.NewInt(ctx.BlockHeight())
	signer := ethtypes.MakeSigner(ethCfg, blockNum)

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}
		if msgEthTx.From == "" {
			if err := esvd.anteHandle(msgEthTx, signer); err != nil {
				return ctx, err
			}
		} else {
			ethAddr := common.HexToAddress(msgEthTx.From)
			cosmosAddr := sdk.AccAddress(ethAddr.Bytes())
			account := esvd.accountKeeper.GetAccount(ctx, cosmosAddr)

			pubKey := account.GetPubKey()
			if pubKey != nil && pubKey.Type() == ethsecp256k1.KeyType {
				if err := esvd.anteHandle(msgEthTx, signer); err != nil {
					return ctx, err
				}
			} else {
				if err := esvd.anteHandleSm2(ctx, msgEthTx, tx, simulate); err != nil {
					return ctx, err
				}
			}
		}

	}
	return next(ctx, tx, simulate)
}

func (esvd EthSigVerificationDecorator) anteHandleSm2(ctx sdk.Context, msgEthTx *evmtypes.MsgEthereumTx, tx sdk.Tx, simulate bool) error {

	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid unpack transaction data")
	}
	_, r, s := txData.GetRawSignatureValues()
	R := r.Bytes()
	S := s.Bytes()
	sig := make([]byte, 64)
	copy(sig[32-len(R):32], R[:])
	copy(sig[64-len(S):64], S[:])

	signerAddrs := sigTx.GetSigners()
	acc, err := authante.GetSignerAcc(ctx, esvd.accountKeeper, signerAddrs[0])
	if err != nil {
		return err
	}
	//todo

	msgEthTx.From = common.BytesToAddress(acc.GetAddress()).String()
	// retrieve pubkey
	pubKey := acc.GetPubKey()
	if !simulate && pubKey == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
	}

	// Check account sequence number.
	if txData.GetNonce() != acc.GetSequence() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrWrongSequence,
			"account sequence mismatch, expected %d, got %d", acc.GetSequence(), txData.GetNonce(),
		)
	}

	chainID, err := ethermint.ParseChainID(ctx.ChainID())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidChainID, "chainID is invalid %s", chainID)
	}

	signer := evmcrypto.NewSm2Signer(chainID)
	ethTx := msgEthTx.AsTransaction()
	txHash := signer.Hash(ethTx)
	if !simulate {
		if !pubKey.VerifySignature(txHash.Bytes(), sig) {
			return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "unable to verify single signer signature")
		}
	}

	return nil
}

func (esvd EthSigVerificationDecorator) anteHandle(msgEthTx *evmtypes.MsgEthereumTx, signer ethtypes.Signer) error {

	sender, err := signer.Sender(msgEthTx.AsTransaction())
	if err != nil {
		return sdkerrors.Wrapf(
			sdkerrors.ErrorInvalidSigner,
			"couldn't retrieve sender address ('%s') from the ethereum transaction: %s",
			msgEthTx.From,
			err.Error(),
		)
	}

	// set up the sender to the transaction field if not already
	msgEthTx.From = sender.Hex()
	return nil
}

// EthValidateBasicDecorator is adapted from ValidateBasicDecorator from cosmos-sdk, it ignores ErrNoSignatures
type EthValidateBasicDecorator struct {
	evmKeeper EVMKeeper
}

// NewEthValidateBasicDecorator creates a new EthValidateBasicDecorator
func NewEthValidateBasicDecorator(ek EVMKeeper) EthValidateBasicDecorator {
	return EthValidateBasicDecorator{
		evmKeeper: ek,
	}
}

// AnteHandle handles basic validation of tx
func (vbd EthValidateBasicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// no need to validate basic on recheck tx, call next antehandler
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}

	err := tx.ValidateBasic()
	// ErrNoSignatures is fine with eth tx
	if err != nil && !errors.Is(err, sdkerrors.ErrNoSignatures) {
		return ctx, sdkerrors.Wrap(err, "tx basic validation failed")
	}

	// For eth type cosmos tx, some fields should be veified as zero values,
	// since we will only verify the signature against the hash of the MsgEthereumTx.Data
	if wrapperTx, ok := tx.(protoTxProvider); ok {
		protoTx := wrapperTx.GetProtoTx()
		body := protoTx.Body
		if body.Memo != "" || body.TimeoutHeight != uint64(0) || len(body.NonCriticalExtensionOptions) > 0 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"for eth tx body Memo TimeoutHeight NonCriticalExtensionOptions should be empty")
		}

		if len(body.ExtensionOptions) != 1 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx length of ExtensionOptions should be 1")
		}

		if len(protoTx.GetMsgs()) != 1 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "only 1 ethereum msg supported per tx")
		}
		msg := protoTx.GetMsgs()[0]
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
		}
		ethGasLimit := msgEthTx.GetGas()

		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, sdkerrors.Wrap(err, "failed to unpack MsgEthereumTx Data")
		}
		params := vbd.evmKeeper.GetParams(ctx)

		if !ctx.MinGasPrices().IsZero() {
			amount := ctx.MinGasPrices().AmountOf(vbd.evmKeeper.GetParams(ctx).EvmDenom)
			if !amount.IsZero() {

				// consistent with ethermint
				// https://github.com/bianjieai/ethermint/blob/5df518e6293679271fb7ec866bddaade4c946099/types/coin.go?_pjax=%23js-repo-pjax-container%2C%20div%5Bitemtype%3D%22http%3A%2F%2Fschema.org%2FSoftwareSourceCode%22%5D%20main%2C%20%5Bdata-pjax-container%5D#L24
				var defaultAmont int64 = ethermint.DefaultGasPrice
				minGasFee := amount.RoundInt64()

				if minGasFee != 0 {
					defaultAmont = minGasFee
				}
				txGasPrice := txData.GetGasPrice()
				gasPrice := new(big.Int).SetInt64(defaultAmont)
				if txGasPrice.Cmp(gasPrice) == -1 {
					return ctx, sdkerrors.New(ethermint.RootCodespace, 101, "invalid gas price")
				}
			}
		}

		ethFeeAmount := sdk.Coins{sdk.NewCoin(params.EvmDenom, sdk.NewIntFromBigInt(txData.Fee()))}

		authInfo := protoTx.AuthInfo
		if len(authInfo.SignerInfos) > 0 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx AuthInfo SignerInfos should be empty")
		}

		if authInfo.Fee.Payer != "" || authInfo.Fee.Granter != "" {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx AuthInfo Fee payer and granter should be empty")
		}

		if !authInfo.Fee.Amount.IsEqual(ethFeeAmount) {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid eth tx AuthInfo Fee Amount")
		}

		if authInfo.Fee.GasLimit != ethGasLimit {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid eth tx AuthInfo Fee GasLimit")
		}

		sigs := protoTx.Signatures
		if len(sigs) > 0 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "for eth tx Signatures should be empty")
		}
	}

	return next(ctx, tx, simulate)
}

type ContractCallable interface {
	GetBlockContract(sdk.Context, []byte) bool
}

type EthContractCallableDecorator struct {
	contractCallable ContractCallable
}

func NewEthContractCallableDecorator(contractCallable ContractCallable) EthContractCallableDecorator {
	return EthContractCallableDecorator{contractCallable: contractCallable}
}

func (e EthContractCallableDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
		}
		ethTx := msgEthTx.AsTransaction()
		if ethTx.To() != nil {
			state := e.contractCallable.GetBlockContract(ctx, ethTx.To().Bytes())
			if state {
				return ctx, sdkerrors.Wrapf(permtypes.ErrContractDisable, "the contract %s is in contract deny list ! ", ethTx.To())
			}
		}
	}
	return next(ctx, tx, simulate)
}

// CanTransferDecorator checks if the sender is allowed to transfer funds according to the EVM block
// context rules.
type CanTransferDecorator struct {
	evmKeeper   EVMKeeper
	opbKeeper   opbkeeper.Keeper
	tokenKeeper tokenkeeper.Keeper
	permKeeper  permkeeper.Keeper
}

// NewCanTransferDecorator creates a new CanTransferDecorator instance.
func NewCanTransferDecorator(evmKeeper EVMKeeper, opbKeeper opbkeeper.Keeper, tokenKeeper tokenkeeper.Keeper, permKeeper permkeeper.Keeper) CanTransferDecorator {
	return CanTransferDecorator{
		evmKeeper:   evmKeeper,
		opbKeeper:   opbKeeper,
		tokenKeeper: tokenKeeper,
		permKeeper:  permKeeper,
	}
}

// AnteHandle creates an EVM from the message and calls the BlockContext CanTransfer function to
// see if the address can execute the transaction.
func (ctd CanTransferDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	params := ctd.evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(ctd.evmKeeper.ChainID())
	signer := ethtypes.MakeSigner(ethCfg, big.NewInt(ctx.BlockHeight()))

	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid message type %T, expected %T", msg, (*evmtypes.MsgEthereumTx)(nil))
		}

		baseFee := ctd.evmKeeper.BaseFee(ctx, ethCfg)

		coreMsg, err := msgEthTx.AsMessage(signer, baseFee)
		if err != nil {
			return ctx, sdkerrors.Wrapf(
				err,
				"failed to create an ethereum core.Message from signer %T", signer,
			)
		}

		// NOTE: pass in an empty coinbase address and nil tracer as we don't need them for the check below
		cfg := &evmtypes.EVMConfig{
			ChainConfig: ethCfg,
			Params:      params,
			CoinBase:    common.Address{},
			BaseFee:     baseFee,
		}
		stateDB := statedb.New(ctx, ctd.evmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(ctx.HeaderHash().Bytes())))
		evm := ctd.evmKeeper.NewEVM(ctx, coreMsg, cfg, evmtypes.NewNoOpTracer(), stateDB)

		// author: sheldon@bianjie.ai
		// Whether to allow transfers

		validator := NewEthOpbValidator(ctx, ctd.opbKeeper, ctd.tokenKeeper, ctd.evmKeeper, ctd.permKeeper)
		evm.Context.CanTransfer = validator.CanTransfer

		if coreMsg.To() != nil {

			if coreMsg.Value().Sign() > 0 && !evm.Context.CanTransfer(stateDB, *coreMsg.To(), coreMsg.Value()) {
				return ctx, sdkerrors.Wrapf(
					sdkerrors.ErrInsufficientFunds,
					"failed to transfer %s from address %s using the EVM block context transfer function",
					coreMsg.Value(),
					coreMsg.From(),
				)
			}

		}

		// check that caller has enough balance to cover asset transfer for **topmost** call
		// NOTE: here the gas consumed is from the context with the infinite gas meter
		if coreMsg.Value().Sign() > 0 && !evm.Context.CanTransfer(stateDB, coreMsg.From(), coreMsg.Value()) {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrInsufficientFunds,
				"failed to transfer %s from address %s using the EVM block context transfer function",
				coreMsg.Value(),
				coreMsg.From(),
			)
		}

		if evmtypes.IsLondon(ethCfg, ctx.BlockHeight()) {
			if baseFee == nil {
				return ctx, sdkerrors.Wrap(
					evmtypes.ErrInvalidBaseFee,
					"base fee is supported but evm block context value is nil",
				)
			}
			if coreMsg.GasFeeCap().Cmp(baseFee) < 0 {
				return ctx, sdkerrors.Wrapf(
					sdkerrors.ErrInsufficientFee,
					"max fee per gas less than block base fee (%s < %s)",
					coreMsg.GasFeeCap(), baseFee,
				)
			}
		}
	}

	return next(ctx, tx, simulate)
}

type EthFeeGrantValidator struct {
	feegrantKeeper FeegrantKeeper
	evmKeeper      EVMKeeper
}

// NewEthFeeGrantValidator creates a new EthFeeGrantValidator
func NewEthFeeGrantValidator(evmKeeper EVMKeeper, fk FeegrantKeeper) EthFeeGrantValidator {
	return EthFeeGrantValidator{
		feegrantKeeper: fk,
		evmKeeper:      evmKeeper,
	}
}

func (ev EthFeeGrantValidator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	params := ev.evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(ev.evmKeeper.ChainID())
	blockNum := big.NewInt(ctx.BlockHeight())
	signer := ethtypes.MakeSigner(ethCfg, blockNum)
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
		}
		sender, err := signer.Sender(msgEthTx.AsTransaction())
		if err != nil {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrorInvalidSigner,
				"couldn't retrieve sender address ('%s') from the ethereum transaction: %s",
				msgEthTx.From,
				err.Error(),
			)
		}
		fromAddr := sender.Bytes()

		feePayerAddr := msgEthTx.GetFeePayer()
		if feePayerAddr != nil {
			if _, err := ev.feegrantKeeper.GetAllowance(ctx, feePayerAddr, fromAddr); err != nil {
				return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type %T, expected %T", tx, (*evmtypes.MsgEthereumTx)(nil))
			}
		}

	}
	return next(ctx, tx, simulate)
}
