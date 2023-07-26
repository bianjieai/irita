package keeper

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/opb/types"
)

type EVMTransferCreator struct {
	opbKeeper   Keeper
	tokenKeeper types.TokenKeeper
	evmKeeper   types.EVMKeeper
	permKeeper  types.PermKeeper
}

func (k Keeper) NewEVMTransferCreator(
	tokenKeeper types.TokenKeeper,
	evmKeeper types.EVMKeeper,
	permKeeper types.PermKeeper,
) *EVMTransferCreator {
	return &EVMTransferCreator{
		opbKeeper:   k,
		tokenKeeper: tokenKeeper,
		evmKeeper:   evmKeeper,
		permKeeper:  permKeeper,
	}
}

func (ov *EVMTransferCreator) CanTransfer(ctx sdk.Context) vm.CanTransferFunc {
	return func(db vm.StateDB, userAddr common.Address, amount *big.Int) bool {
		userCosmosAddr := sdk.AccAddress(userAddr.Bytes())

		// get opb params
		params := ov.evmKeeper.GetParams(ctx)
		restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ctx)

		// check only if the transfer restriction is enabled
		if restrictionEnabled {
			owner, err := ov.getOwner(ctx, params.EvmDenom)
			if err != nil {
				return false
			}
			if ov.hasPlatformUserPerm(ctx, userCosmosAddr) {
				return false
			}
			if userCosmosAddr.String() != owner {
				return false
			}
		}

		return db.GetBalance(userAddr).Cmp(amount) >= 0
	}
}

func (ov *EVMTransferCreator) Transfer(ctx sdk.Context) vm.TransferFunc {
	return func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
		senderCosmosAddr := sdk.AccAddress(sender.Bytes())
		recipientCosmosAddr := sdk.AccAddress(recipient.Bytes())

		params := ov.evmKeeper.GetParams(ctx)
		restrictionEnabled := !ov.opbKeeper.UnrestrictedTokenTransfer(ctx)
		// check only if the transfer restriction is enabled
		if restrictionEnabled {
			owner, err := ov.getOwner(ctx, params.EvmDenom)
			if err != nil {
				//ov.logger.Error("unauthorized operation", "err_msg", err.Error())
				ov.opbKeeper.Logger(ctx).Error(
					"unauthorized operation",
					"err_msg", err.Error(),
					"amount", amount.Int64(),
				)
				return
			}
			if senderCosmosAddr.String() != owner && recipientCosmosAddr.String() != owner {
				errMsg := fmt.Sprintf(
					"either the sender or recipient must be the owner %s for token %s",
					owner,
					params.EvmDenom,
				)
				//ov.logger.Error("unauthorized operation", "err_msg", errMsg)
				ov.opbKeeper.Logger(ctx).Error(
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
}

func (ov *EVMTransferCreator) getOwner(ctx sdk.Context, denom string) (owner string, err error) {
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
func (ov *EVMTransferCreator) hasPlatformUserPermFromArr(ctx sdk.Context, addresses []string) bool {
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
func (ov *EVMTransferCreator) hasPlatformUserPerm(ctx sdk.Context, address sdk.AccAddress) bool {
	return ov.permKeeper.IsRootAdmin(ctx, address) || ov.permKeeper.IsBaseM1Admin(ctx, address) ||
		ov.permKeeper.IsPlatformUser(ctx, address)
}
