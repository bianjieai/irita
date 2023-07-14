package evm

import (
	"fmt"
	"math/big"

	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

type EthOpbValidator struct {
	opbKeeper   opbkeeper.Keeper
	tokenKeeper tokenkeeper.Keeper
	evmKeeper   EVMKeeper
	permKeeper  permkeeper.Keeper

	ctx sdk.Context
}

func NewEthOpbValidator(
	ctx sdk.Context,
	opbKeeper opbkeeper.Keeper,
	tokenKeeper tokenkeeper.Keeper,
	evmKeeper EVMKeeper,
	permKeeper permkeeper.Keeper,
) *EthOpbValidator {
	return &EthOpbValidator{
		opbKeeper:   opbKeeper,
		tokenKeeper: tokenKeeper,
		evmKeeper:   evmKeeper,
		permKeeper:  permKeeper,
		ctx:         ctx,
	}
}

func (ov EthOpbValidator) Transfer(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {

	senderCosmosAddr := sdk.AccAddress(sender.Bytes())
	recipientCosmosAddr := sdk.AccAddress(recipient.Bytes())

	params := ov.evmKeeper.GetParams(ov.ctx)
	restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ov.ctx)
	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		owner, err := ov.getOwner(ov.ctx, params.EvmDenom)
		if err != nil {
			//ov.logger.Error("unauthorized operation", "err_msg", err.Error())
			ov.opbKeeper.Logger(ov.ctx).Error(
				"unauthorized operation",
				"err_msg", err.Error(),
				"amount", amount.Int64(),
			)
			return
		}
		if senderCosmosAddr.String() != owner && recipientCosmosAddr.String() != owner {
			errMsg := fmt.Sprintf("either the sender or recipient must be the owner %s for token %s", owner, params.EvmDenom)
			//ov.logger.Error("unauthorized operation", "err_msg", errMsg)
			ov.opbKeeper.Logger(ov.ctx).Error(
				"unauthorized operation",
				"err_msg", errMsg,
				"amount", amount.Int64(),
			)
			return
		}
	}
	// go-ethereum
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

func (ov EthOpbValidator) CanTransfer(db vm.StateDB, userAddr common.Address, amount *big.Int) bool {

	userCosmosAddr := sdk.AccAddress(userAddr.Bytes())

	// get opb params
	params := ov.evmKeeper.GetParams(ov.ctx)
	restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ov.ctx)

	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		owner, err := ov.getOwner(ov.ctx, params.EvmDenom)
		if err != nil {
			return false
		}
		if ov.hasPlatformUserPerm(ov.ctx, userCosmosAddr) {
			return false
		}
		if userCosmosAddr.String() != owner {
			return false
		}
	}

	return db.GetBalance(userAddr).Cmp(amount) >= 0
}

func (ov EthOpbValidator) getOwner(ctx sdk.Context, denom string) (owner string, err error) {
	baseTokenDenom := ov.opbKeeper.BaseTokenDenom(ctx)

	if denom == baseTokenDenom {
		owner = ov.opbKeeper.BaseTokenManager(ctx)
	} else {
		ownerAddr, err := ov.tokenKeeper.GetOwner(ctx, denom)
		if err == nil {
			owner = ownerAddr.String()
		}
	}
	return
}

// hasPlatformUserPermFromArr determine whether the account is a platform user from addresses
func (ov EthOpbValidator) hasPlatformUserPermFromArr(ctx sdk.Context, addresses []string) bool {
	for _, addr := range addresses {
		fromAddress, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return false
		}
		if !ov.hasPlatformUserPerm(ctx, fromAddress) {
			return false
		}
	}

	return true
}

// hasPlatformUserPerm determine whether the account is a platform user
func (ov EthOpbValidator) hasPlatformUserPerm(ctx sdk.Context, address sdk.AccAddress) bool {
	return ov.permKeeper.IsRootAdmin(ctx, address) || ov.permKeeper.IsBaseM1Admin(ctx, address) || ov.permKeeper.IsPlatformUser(ctx, address)
}
