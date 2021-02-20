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

	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/params"
	"github.com/bianjieai/iritamod/modules/perm"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, deducts fees from the first
// signer, and performs other module-specific logic.
func NewAnteHandler(
	permKeeper perm.Keeper,
	ak authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	tokenKeeper tokenkeeper.Keeper,
	opbKeeper opbkeeper.Keeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer,
	signModeHandler signing.SignModeHandler,
) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		perm.NewAuthDecorator(permKeeper),
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		ante.NewRejectFeeGranterDecorator(),
		ante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		ante.NewDeductFeeDecorator(ak, bankKeeper),
		ante.NewSigGasConsumeDecorator(ak, sigGasConsumer),
		ante.NewSigVerificationDecorator(ak, signModeHandler),
		ante.NewIncrementSequenceDecorator(ak),
		tokenkeeper.NewValidateTokenFeeDecorator(tokenKeeper, bankKeeper),
		opbkeeper.NewValidateTokenTransferDecorator(opbKeeper, tokenKeeper),
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

	return permKeeper
}
