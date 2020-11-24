package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"

	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"
	tokentypes "github.com/irisnet/irismod/modules/token/types"

	"github.com/bianjieai/iritamod/modules/admin"
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/params"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"
	"github.com/bianjieai/iritamod/modules/validator"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(
	adminKeeper admin.Keeper,
	ak authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	tk tokenkeeper.Keeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		admin.NewAuthDecorator(adminKeeper),
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		ante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		ante.NewDeductFeeDecorator(ak, bankKeeper),
		ante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		ante.NewSigVerificationDecorator(ak, signModeHandler),
		ante.NewIncrementSequenceDecorator(ak),
		tokenkeeper.NewValidateTokenFeeDecorator(tk, bankKeeper),
	)
}

func RegisterAccessControl(adminKeeper admin.Keeper) admin.Keeper {
	// permission auth
	adminKeeper.RegisterMsgAuth(&admin.MsgAddRoles{}, admin.RoleRootAdmin, admin.RolePermAdmin)
	adminKeeper.RegisterMsgAuth(&admin.MsgRemoveRoles{}, admin.RoleRootAdmin, admin.RolePermAdmin)

	// blacklist auth
	adminKeeper.RegisterMsgAuth(&admin.MsgBlockAccount{}, admin.RoleRootAdmin, admin.RoleBlacklistAdmin)
	adminKeeper.RegisterMsgAuth(&admin.MsgUnblockAccount{}, admin.RoleRootAdmin, admin.RoleBlacklistAdmin)

	// node auth
	adminKeeper.RegisterModuleAuth(validator.ModuleName, admin.RoleRootAdmin, admin.RoleNodeAdmin)
	adminKeeper.RegisterModuleAuth(slashingtypes.ModuleName, admin.RoleRootAdmin, admin.RoleNodeAdmin)

	// param auth
	adminKeeper.RegisterModuleAuth(params.ModuleName, admin.RoleRootAdmin, admin.RoleParamAdmin)

	// identity auth
	adminKeeper.RegisterMsgAuth(&identity.MsgCreateIdentity{}, admin.RoleRootAdmin, admin.RoleIDAdmin)

	// oracle auth
	adminKeeper.RegisterModuleAuth(oracletypes.ModuleName, admin.RoleRootAdmin, admin.RolePowerUser)

	// power user auth
	adminKeeper.RegisterMsgAuth(&tokentypes.MsgIssueToken{}, admin.RoleRootAdmin, admin.RolePowerUser)
	adminKeeper.RegisterMsgAuth(&nfttypes.MsgIssueDenom{}, admin.RoleRootAdmin, admin.RolePowerUser)
	adminKeeper.RegisterMsgAuth(&servicetypes.MsgDefineService{}, admin.RoleRootAdmin, admin.RolePowerUser)
	adminKeeper.RegisterMsgAuth(&servicetypes.MsgBindService{}, admin.RoleRootAdmin, admin.RolePowerUser)

	// upgrade auth
	adminKeeper.RegisterModuleAuth(upgradetypes.ModuleName, admin.RoleRootAdmin, admin.RoleNodeAdmin)

	return adminKeeper
}
