package ante

import (
	"fmt"
	"runtime/debug"

	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"

	"github.com/bianjieai/irita/modules/gas"

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

	evmmoduleante "github.com/bianjieai/irita/modules/evm"
	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	tibctypes "github.com/bianjieai/irita/modules/tibc/types"
	"github.com/bianjieai/iritamod/modules/identity"
	layer2types "github.com/bianjieai/iritamod/modules/layer2/types"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/bianjieai/iritamod/modules/perm"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"

	ethermintante "github.com/tharsis/ethermint/app/ante"
)

type HandlerOptions struct {
	PermKeeper      perm.Keeper
	AccountKeeper   authkeeper.AccountKeeper
	BankKeeper      bankkeeper.Keeper
	FeegrantKeeper  feegrantkeeper.Keeper
	TokenKeeper     tokenkeeper.Keeper
	OpbKeeper       opbkeeper.Keeper
	SigGasConsumer  ante.SignatureVerificationGasConsumer
	SignModeHandler signing.SignModeHandler

	// evm config
	EvmKeeper          evmmoduleante.EVMKeeper
	EvmFeeMarketKeeper evmtypes.FeeMarketKeeper
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
						ethermintante.NewEthSetUpContextDecorator(options.EvmKeeper), // outermost AnteDecorator. SetUpContext must be called first
						ante.NewMempoolFeeDecorator(),
						ante.NewTxTimeoutHeightDecorator(),
						ante.NewValidateMemoDecorator(options.AccountKeeper),
						evmmoduleante.NewEthValidateBasicDecorator(options.EvmKeeper),
						evmmoduleante.NewEthFeeGrantValidator(options.EvmKeeper, options.FeegrantKeeper),
						evmmoduleante.NewEthContractCallableDecorator(options.PermKeeper),
						evmmoduleante.NewEthSigVerificationDecorator(options.EvmKeeper, options.AccountKeeper, options.SignModeHandler),
						evmmoduleante.NewCanTransferDecorator(options.EvmKeeper, options.OpbKeeper, options.TokenKeeper, options.PermKeeper),

						ethermintante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.BankKeeper, options.EvmKeeper),
						ethermintante.NewEthGasConsumeDecorator(options.EvmKeeper),
						ethermintante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
						ethermintante.NewEthMempoolFeeDecorator(options.EvmKeeper),                  // Check eth effective gas price against minimal-gas-prices
						ethermintante.NewEthValidateBasicDecorator(options.EvmKeeper),

						perm.NewAuthDecorator(options.PermKeeper),
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
				gas.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
				perm.NewAuthDecorator(options.PermKeeper),
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
				opbkeeper.NewValidateTokenTransferDecorator(options.OpbKeeper, options.TokenKeeper, options.PermKeeper),
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
	permKeeper.RegisterMsgAuth(&perm.MsgAssignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin, perm.RolePowerUserAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnassignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin, perm.RolePowerUserAdmin)

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

	// layer2 auth
	permKeeper.RegisterMsgAuth(&layer2types.MsgCreateL2Space{}, perm.RoleRootAdmin, perm.RoleLayer2User)
	permKeeper.RegisterMsgAuth(&layer2types.MsgTransferL2Space{}, perm.RoleRootAdmin, perm.RoleLayer2User)
	permKeeper.RegisterMsgAuth(&layer2types.MsgCreateL2BlockHeader{}, perm.RoleRootAdmin, perm.RoleLayer2User)
	// NOTE： The following messages are currently closed to users.
	permKeeper.RegisterMsgAuth(&layer2types.MsgCreateNFTs{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgUpdateNFTs{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgDeleteNFTs{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgUpdateClassesForNFT{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgDepositClassForNFT{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgWithdrawClassForNFT{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgDepositTokenForNFT{}, perm.RoleRootAdmin)
	permKeeper.RegisterMsgAuth(&layer2types.MsgWithdrawTokenForNFT{}, perm.RoleRootAdmin)

	return permKeeper
}
