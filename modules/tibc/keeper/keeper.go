package keeper

import (
	tibckeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
)

// Keeper defines each TICS keeper for TIBC
type Keeper struct {
	*tibckeeper.Keeper
}

func NewKeeper(k *tibckeeper.Keeper) *Keeper {
	return &Keeper{k}
}
