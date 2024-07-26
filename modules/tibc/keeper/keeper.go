package keeper

import (
	tibckeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
	nftkeeper "mods.irisnet.org/modules/nft/keeper"
)

// Keeper defines each TICS keeper for TIBC
type (
	Keeper struct {
		*tibckeeper.Keeper
	}

	NFTKeeper struct {
		nk nftkeeper.Keeper
	}
)

func NewKeeper(k *tibckeeper.Keeper) *Keeper {
	return &Keeper{k}
}
