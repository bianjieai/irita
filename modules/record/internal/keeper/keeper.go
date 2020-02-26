package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/bianjieai/irita/modules/record/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the record store
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper returns a record keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey: key,
		cdc:      cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("%s", types.ModuleName))
}

// AddRecord add a record
func (k Keeper) AddRecord(ctx sdk.Context, record types.Record) []byte {
	store := ctx.KVStore(k.storeKey)
	recordBz := k.cdc.MustMarshalBinaryLengthPrefixed(record)
	intraTxCounter := k.GetIntraTxCounter(ctx)

	bz := make([]byte, 2+len(recordBz))
	copy(bz[:len(recordBz)], recordBz[:])
	binary.BigEndian.PutUint16(bz[len(recordBz):], intraTxCounter)

	recordID := getRecordId(bz)
	store.Set(types.GetRecordKey(recordID), recordBz)

	// update intraTxCounter + 1
	k.SetIntraTxCounter(ctx, intraTxCounter+1)
	return recordID
}

// GetRecord retrieves the record by specified recordID
func (k Keeper) GetRecord(ctx sdk.Context, recordID []byte) (record types.Record, found bool) {
	store := ctx.KVStore(k.storeKey)
	if bz := store.Get(types.GetRecordKey(recordID)); bz != nil {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &record)
		return record, true
	}
	return record, false
}

// RecordsIterator gets all records
func (k Keeper) RecordsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.RecordKey)
}

// GetIntraTxCounter gets the current in-block request operation counter
func (k Keeper) GetIntraTxCounter(ctx sdk.Context) uint16 {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.IntraTxCounterKey)
	if b == nil {
		return 0
	}

	var counter uint16
	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &counter)

	return counter
}

// SetIntraTxCounter sets the current in-block request counter
func (k Keeper) SetIntraTxCounter(ctx sdk.Context, counter uint16) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(counter)
	store.Set(types.IntraTxCounterKey, bz)
}

func getRecordId(bz []byte) []byte {
	return tmhash.Sum(bz)
}
