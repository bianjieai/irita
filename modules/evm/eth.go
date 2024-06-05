package evm

import (
	"math/big"

	"github.com/tharsis/ethermint/crypto/ethsecp256k1"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"

	evmcrypto "github.com/bianjieai/irita/modules/evm/crypto"

	ethermint "github.com/tharsis/ethermint/types"
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

type ContractCallable interface {
	GetBlockContract(sdk.Context, []byte) bool
}
