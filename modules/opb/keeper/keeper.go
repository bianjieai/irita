package keeper

import (
	"bytes"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/bianjieai/irita/modules/opb/types"
)

// Keeper defines the OPB keeper
type Keeper struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	tokenKeeper   types.TokenKeeper
	permKeeper    types.PermKeeper

	paramSpace paramstypes.Subspace
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.Codec,
	storeKey sdk.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	tokenKeeper types.TokenKeeper,
	permKeeper types.PermKeeper,
	paramSpace paramstypes.Subspace,
) Keeper {
	// ensure the OPB module account is set
	if addr := accountKeeper.GetModuleAddress(types.PointTokenFeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.PointTokenFeeCollectorName))
	}

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
		tokenKeeper:   tokenKeeper,
		permKeeper:    permKeeper,
		paramSpace:    paramSpace,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irita/%s", types.ModuleName))
}

// Mint mints the base native token by the specified amount
// NOTE: the operator must possess the BaseM1Admin or RootAdmin permission
func (k Keeper) Mint(ctx sdk.Context, amount uint64, recipient, operator sdk.AccAddress) error {
	// get the base token denom
	baseTokenDenom := k.BaseTokenDenom(ctx)

	if !k.hasBaseM1Perm(ctx, operator) {
		return sdkerrors.Wrapf(types.ErrUnauthorized, "address %s has no permission to mint %s", operator, baseTokenDenom)
	}

	// get the base token
	baseToken, err := k.tokenKeeper.GetToken(ctx, baseTokenDenom)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "token for %s does not exist", baseTokenDenom)
	}

	// NOTE: empty owner
	owner := sdk.AccAddress{}

	return k.tokenKeeper.MintToken(ctx, baseToken.GetSymbol(), amount, recipient, owner)
}

// Reclaim reclaims the native token of the specified denom from the corresponding escrow account
// NOTE: the operator must possess the certain permission
func (k Keeper) Reclaim(ctx sdk.Context, denom string, recipient, operator sdk.AccAddress) error {
	baseTokenDenom := k.BaseTokenDenom(ctx)
	pointTokenDenom := k.PointTokenDenom(ctx)

	var moduleAccName string

	switch denom {
	case baseTokenDenom, "ugas":
		if !k.hasBaseM1Perm(ctx, operator) {
			return sdkerrors.Wrapf(types.ErrUnauthorized, "address %s has no permission to reclaim %s", operator, denom)
		}

		moduleAccName = authtypes.FeeCollectorName

	case pointTokenDenom:
		owner, err := k.tokenKeeper.GetOwner(ctx, denom)
		if err != nil {
			return sdkerrors.Wrapf(types.ErrInvalidDenom, "token for %s does not exist", denom)
		}

		if !bytes.Equal(operator, owner) {
			return sdkerrors.Wrapf(types.ErrUnauthorized, "only %s is allowed to reclaim %s", owner, denom)
		}

		moduleAccName = types.PointTokenFeeCollectorName

	default:
		return sdkerrors.Wrapf(types.ErrInvalidDenom, "denom must be either %s or %s", baseTokenDenom, pointTokenDenom)
	}

	moduleAccAddr := k.accountKeeper.GetModuleAddress(moduleAccName)

	balance := k.bankKeeper.GetBalance(ctx, moduleAccAddr, denom)
	if balance.IsZero() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "no balance for %s in the module account", denom)
	}

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleAccName, recipient, sdk.NewCoins(balance))
}

// HasToken checks if the given token exists
func (k Keeper) HasToken(ctx sdk.Context, denom string) bool {
	if _, err := k.tokenKeeper.GetToken(ctx, denom); err != nil {
		return false
	}

	return true
}

// hasBaseM1Perm returns true if the given address is BaseM1Admin or RootAdmin
// False otherwise
func (k Keeper) hasBaseM1Perm(ctx sdk.Context, address sdk.AccAddress) bool {
	return k.permKeeper.IsRootAdmin(ctx, address) || k.permKeeper.IsBaseM1Admin(ctx, address)
}
