package keeper

import (
	"strings"

	"github.com/bianjieai/irita/modules/opb/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (k Keeper) GetContractDenyList(ctx sdk.Context) []string {
	list, err := k.IteratorContractDanyList(ctx)
	if err != nil {
		return nil
	}
	return list
}

func (k Keeper) GetContractState(ctx sdk.Context, contractAddress string) (bool, error) {
	store := k.GetStore(ctx)
	contractAddr := common.HexToAddress(contractAddress)
	get := store.Get(types.ContractDenyListKey(contractAddr))
	if get != nil {
		return true, nil
	} else {
		return false, nil
	}
}

// GetStore get the store by keeper.storeKey
func (k Keeper) GetStore(ctx sdk.Context) sdk.KVStore {
	store := ctx.KVStore(k.storeKey)
	return store
}

// IteratorContractDanyList iterator the contract addresses in ContractDanyList
func (k Keeper) IteratorContractDanyList(goCtx sdk.Context) ([]string, error) {
	gm := make([]string, 0)
	store := k.GetStore(goCtx)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.KeyPrefixContractDenyList))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		split := strings.Split(string(key), "/")
		if len(split) < 2 {
			return nil, errors.Wrapf(types.ErrNotFound, "not found any contract address from ContractDanyList")
		}
		gm = append(gm, split[1])
	}
	return gm, nil
}
