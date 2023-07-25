package evm

import (
	"errors"
	"math/big"

	permtypes "github.com/bianjieai/iritamod/modules/perm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	ethermint "github.com/evmos/ethermint/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

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
func (vbd EthValidateBasicDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (sdk.Context, error) {
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
		if body.Memo != "" || body.TimeoutHeight != uint64(0) ||
			len(body.NonCriticalExtensionOptions) > 0 {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"for eth tx body Memo TimeoutHeight NonCriticalExtensionOptions should be empty")
		}

		if len(body.ExtensionOptions) != 1 {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"for eth tx length of ExtensionOptions should be 1",
			)
		}

		if len(protoTx.GetMsgs()) != 1 {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"only 1 ethereum msg supported per tx",
			)
		}
		msg := protoTx.GetMsgs()[0]
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"invalid transaction type %T, expected %T",
				tx,
				(*evmtypes.MsgEthereumTx)(nil),
			)
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
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"for eth tx AuthInfo SignerInfos should be empty",
			)
		}

		if authInfo.Fee.Payer != "" || authInfo.Fee.Granter != "" {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"for eth tx AuthInfo Fee payer and granter should be empty",
			)
		}

		if !authInfo.Fee.Amount.IsEqual(ethFeeAmount) {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"invalid eth tx AuthInfo Fee Amount",
			)
		}

		if authInfo.Fee.GasLimit != ethGasLimit {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"invalid eth tx AuthInfo Fee GasLimit",
			)
		}

		sigs := protoTx.Signatures
		if len(sigs) > 0 {
			return ctx, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"for eth tx Signatures should be empty",
			)
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

func NewEthContractCallableDecorator(
	contractCallable ContractCallable,
) EthContractCallableDecorator {
	return EthContractCallableDecorator{contractCallable: contractCallable}
}

func (e EthContractCallableDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"invalid transaction type %T, expected %T",
				tx,
				(*evmtypes.MsgEthereumTx)(nil),
			)
		}
		ethTx := msgEthTx.AsTransaction()
		if ethTx.To() != nil {
			state := e.contractCallable.GetBlockContract(ctx, ethTx.To().Bytes())
			if state {
				return ctx, sdkerrors.Wrapf(
					permtypes.ErrContractDisable,
					"the contract %s is in contract deny list ! ",
					ethTx.To(),
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

func (ev EthFeeGrantValidator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	params := ev.evmKeeper.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(ev.evmKeeper.ChainID())
	blockNum := big.NewInt(ctx.BlockHeight())
	signer := ethtypes.MakeSigner(ethCfg, blockNum)
	for _, msg := range tx.GetMsgs() {
		msgEthTx, ok := msg.(*evmtypes.MsgEthereumTx)
		if !ok {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrUnknownRequest,
				"invalid transaction type %T, expected %T",
				tx,
				(*evmtypes.MsgEthereumTx)(nil),
			)
		}
		ethTx := msgEthTx.AsTransaction()
		sender, err := signer.Sender(ethTx)
		if err != nil {
			return ctx, sdkerrors.Wrapf(
				sdkerrors.ErrorInvalidSigner,
				"couldn't retrieve sender address ('%s') from the ethereum transaction: %s",
				msgEthTx.From,
				err.Error(),
			)
		}
		txData, err := evmtypes.UnpackTxData(msgEthTx.Data)
		if err != nil {
			return ctx, sdkerrors.Wrap(err, "failed to unpack tx data")
		}
		feeGrantee := sender.Bytes()
		feeGranteeCosmosAddr := sdk.AccAddress(feeGrantee)
		feePayer := msgEthTx.GetFeePayer()
		feeAmt := txData.Fee()
		if feeAmt.Sign() == 0 {
			return ctx, sdkerrors.Wrap(err, "failed to fee amount")
		}

		fees := sdk.Coins{sdk.NewCoin(params.EvmDenom, sdk.NewIntFromBigInt(feeAmt))}

		msgs := []sdk.Msg{msg}

		if feePayer != nil {
			err := ev.feegrantKeeper.UseGrantedFees(ctx, feePayer, feeGrantee, fees, msgs)
			if err != nil {
				return ctx, sdkerrors.Wrapf(
					err,
					"%s(%s) not allowed to pay fees from %s",
					sender.Hex(),
					feeGranteeCosmosAddr,
					feePayer,
				)
			}

		}

	}
	return next(ctx, tx, simulate)
}
