package app

import (
	"fmt"
	"runtime/debug"

	appante "github.com/bianjieai/irita/modules/evm"
	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	tibctypes "github.com/bianjieai/irita/modules/tibc/types"
	wservicekeeper "github.com/bianjieai/irita/modules/wservice/keeper"
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/bianjieai/iritamod/modules/perm"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	tokentypes "github.com/irisnet/irismod/modules/token/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
)

type HandlerOptions struct {
	permKeeper      perm.Keeper
	accountKeeper   authkeeper.AccountKeeper
	bankKeeper      bankkeeper.Keeper
	feegrantKeeper  authante.FeegrantKeeper
	tokenKeeper     tokenkeeper.Keeper
	opbKeeper       opbkeeper.Keeper
	wserviceKeeper  wservicekeeper.IKeeper
	sigGasConsumer  ante.SignatureVerificationGasConsumer
	signModeHandler signing.SignModeHandler

	// evm config
	evmKeeper          appante.EVMKeeper
	evmFeeMarketKeeper evmtypes.FeeMarketKeeper
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, deducts fees from the first
// signer, and performs other module-specific logic.
func NewAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, sim bool,
	) (newCtx sdk.Context, err error) {
		var anteHandler sdk.AnteHandler

		//defer Recover(ctx.Logger(), &err)
		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				switch typeURL := opts[0].GetTypeUrl(); typeURL {
				case "/ethermint.evm.v1.ExtensionOptionsEthereumTx":
					// handle as *evmtypes.MsgEthereumTx
					anteHandler = sdk.ChainAnteDecorators(
						ante.NewMempoolFeeDecorator(),
						ante.NewTxTimeoutHeightDecorator(),
						ante.NewValidateMemoDecorator(options.accountKeeper),
						appante.NewEthValidateBasicDecorator(options.evmKeeper),
						appante.NewEthContractCallableDecorator(options.permKeeper),
						appante.NewEthSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
						appante.NewEthSigVerificationDecorator(options.evmKeeper, options.accountKeeper, options.signModeHandler),
						appante.NewEthAccountVerificationDecorator(options.accountKeeper, options.bankKeeper, options.evmKeeper),
						appante.NewEthNonceVerificationDecorator(options.accountKeeper),
						appante.NewEthGasConsumeDecorator(options.evmKeeper),
						appante.NewCanTransferDecorator(options.evmKeeper, options.evmFeeMarketKeeper),
						appante.NewEthIncrementSenderSequenceDecorator(options.accountKeeper), // innermost AnteDecorator.
						perm.NewAuthDecorator(options.permKeeper),
					)

				default:
					return ctx, sdkerrors.Wrapf(
						sdkerrors.ErrUnknownExtensionOptions,
						"rejecting tx with unsupported extension option: %s",
						typeURL,
					)
				}

				return anteHandler(ctx, tx, sim)
			}
		}
		switch tx.(type) {
		case sdk.Tx:
			anteHandler = sdk.ChainAnteDecorators(
				perm.NewAuthDecorator(options.permKeeper),
				ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
				ante.NewMempoolFeeDecorator(),
				ante.NewValidateBasicDecorator(),
				ante.NewValidateMemoDecorator(options.accountKeeper),
				ante.NewConsumeGasForTxSizeDecorator(options.accountKeeper),
				ante.NewSetPubKeyDecorator(options.accountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
				ante.NewValidateSigCountDecorator(options.accountKeeper),
				ante.NewDeductFeeDecorator(options.accountKeeper, options.bankKeeper, options.feegrantKeeper),
				ante.NewSigGasConsumeDecorator(options.accountKeeper, options.sigGasConsumer),
				ante.NewSigVerificationDecorator(options.accountKeeper, options.signModeHandler),
				ante.NewIncrementSequenceDecorator(options.accountKeeper),
				ante.NewRejectExtensionOptionsDecorator(),
				ante.NewTxTimeoutHeightDecorator(),
				tokenkeeper.NewValidateTokenFeeDecorator(options.tokenKeeper, options.bankKeeper),
				opbkeeper.NewValidateTokenTransferDecorator(options.opbKeeper, options.tokenKeeper),
				wservicekeeper.NewDeduplicationTxDecorator(options.wserviceKeeper),
			)
		default:
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "invalid transaction type: %T", tx)
		}

		return anteHandler(ctx, tx, sim)

	}
}

func Recover(logger tmlog.Logger, err *error) {
	if r := recover(); r != nil {
		*err = sdkerrors.Wrapf(sdkerrors.ErrPanic, "%v", r)

		if e, ok := r.(error); ok {
			logger.Error(
				"ante handler panicked",
				"error", e,
				"stack trace", string(debug.Stack()),
			)
		} else {
			logger.Error(
				"ante handler panicked",
				"recover", fmt.Sprintf("%v", r),
			)
		}
	}
}

func RegisterAccessControl(permKeeper perm.Keeper) perm.Keeper {
	// permission auth
	permKeeper.RegisterMsgAuth(&perm.MsgAssignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnassignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin)

	// blacklist auth
	permKeeper.RegisterMsgAuth(&perm.MsgBlockAccount{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnblockAccount{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgBlockContract{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnblockContract{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)

	// node auth
	permKeeper.RegisterModuleAuth(node.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&node.MsgRemoveValidator{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&node.MsgCreateValidator{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&node.MsgUpdateValidator{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterModuleAuth(slashingtypes.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)

	// param auth
	permKeeper.RegisterModuleAuth(params.ModuleName, perm.RoleRootAdmin, perm.RoleParamAdmin)

	// identity auth
	permKeeper.RegisterMsgAuth(&identity.MsgCreateIdentity{}, perm.RoleRootAdmin, perm.RoleIDAdmin)

	// oracle auth
	permKeeper.RegisterModuleAuth(oracletypes.ModuleName, perm.RoleRootAdmin, perm.RolePowerUser)

	// power user auth
	permKeeper.RegisterMsgAuth(&tokentypes.MsgIssueToken{}, perm.RoleRootAdmin, perm.RolePowerUser)
	permKeeper.RegisterMsgAuth(&nfttypes.MsgIssueDenom{}, perm.RoleRootAdmin, perm.RolePowerUser)
	permKeeper.RegisterMsgAuth(&mttypes.MsgIssueDenom{}, perm.RoleRootAdmin, perm.RolePowerUser)
	permKeeper.RegisterMsgAuth(&servicetypes.MsgDefineService{}, perm.RoleRootAdmin, perm.RolePowerUser)
	permKeeper.RegisterMsgAuth(&servicetypes.MsgBindService{}, perm.RoleRootAdmin, perm.RolePowerUser)

	// upgrade auth
	permKeeper.RegisterModuleAuth(upgradetypes.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)

	// tibc auth
	permKeeper.RegisterModuleAuth(tibctypes.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgCreateClient{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgRegisterRelayer{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgUpgradeClient{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgSetRoutingRules{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)

	return permKeeper
}
