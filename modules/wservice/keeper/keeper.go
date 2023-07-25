package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cometbft/cometbft/libs/log"
	servicekeeper "github.com/irisnet/irismod/modules/service/keeper"

	"github.com/bianjieai/irita/modules/wservice/types"
)

type IKeeper interface {
	Logger(ctx sdk.Context) log.Logger
	SetReqSequence(ctx sdk.Context, reqSequence, value []byte)
	GetReqSequence(ctx sdk.Context, reqSequence []byte) []byte
	QueryReqSequencesFromPrefix(ctx sdk.Context, prefix []byte) sdk.Iterator
	ExistReqSequence(ctx sdk.Context, reqSequence []byte) bool
	DeleteReqSequence(ctx sdk.Context, reqSequence []byte)
	GetServiceKeeper() servicekeeper.Keeper
}

type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey

	serviceKeeper servicekeeper.Keeper
}

// NewKeeper returns a record keeper
func NewKeeper(
	cdc codec.Codec,
	key storetypes.StoreKey,
	serviceKeeper servicekeeper.Keeper,
) IKeeper {
	keeper := &Keeper{
		storeKey:      key,
		cdc:           cdc,
		serviceKeeper: serviceKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("irita/%s", types.ModuleName))
}

func (k *Keeper) SetReqSequence(ctx sdk.Context, reqSequence, value []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(reqSequence, value)
}

func (k *Keeper) GetReqSequence(ctx sdk.Context, reqSequence []byte) []byte {
	store := ctx.KVStore(k.storeKey)
	return store.Get(reqSequence)
}

func (k *Keeper) QueryReqSequencesFromPrefix(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k *Keeper) ExistReqSequence(ctx sdk.Context, reqSequence []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(reqSequence)
}

func (k *Keeper) DeleteReqSequence(ctx sdk.Context, reqSequence []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(reqSequence)
}

func (k *Keeper) GetServiceKeeper() servicekeeper.Keeper {
	return k.serviceKeeper
}
