package app

import (
	tibctypes "github.com/bianjieai/irita/modules/tibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/bianjieai/iritamod/modules/perm"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"

)

type HandlerOptions struct {
	permKeeper      perm.Keeper
	accountKeeper   authkeeper.AccountKeeper
	bankKeeper      bankkeeper.Keeper
	feegrantKeeper  authante.FeegrantKeeper
	tokenKeeper     tokenkeeper.Keeper
	opbKeeper       opbkeeper.Keeper
	sigGasConsumer  ante.SignatureVerificationGasConsumer
	signModeHandler signing.SignModeHandler
}

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, deducts fees from the first
// signer, and performs other module-specific logic.
func NewAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
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
	)
}

func RegisterAccessControl(permKeeper perm.Keeper) perm.Keeper {
	// permission auth
	permKeeper.RegisterMsgAuth(&perm.MsgAssignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnassignRoles{}, perm.RoleRootAdmin, perm.RolePermAdmin)

	// blacklist auth
	permKeeper.RegisterMsgAuth(&perm.MsgBlockAccount{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)
	permKeeper.RegisterMsgAuth(&perm.MsgUnblockAccount{}, perm.RoleRootAdmin, perm.RoleBlacklistAdmin)

	// node auth
	permKeeper.RegisterModuleAuth(node.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)
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
	permKeeper.RegisterMsgAuth(&servicetypes.MsgDefineService{}, perm.RoleRootAdmin, perm.RolePowerUser)
	permKeeper.RegisterMsgAuth(&servicetypes.MsgBindService{}, perm.RoleRootAdmin, perm.RolePowerUser)

	// upgrade auth
	permKeeper.RegisterModuleAuth(upgradetypes.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)

	// tibc auth
	permKeeper.RegisterModuleAuth(tibctypes.ModuleName, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgCreateClient{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgRegisterRelayer{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)
	permKeeper.RegisterMsgAuth(&tibctypes.MsgUpgradeClient{}, perm.RoleRootAdmin, perm.RoleNodeAdmin)

	return permKeeper
}
