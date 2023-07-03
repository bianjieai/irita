package sidechain

import (
	"github.com/cosmos/cosmos-sdk/codec"

	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type (
	PermKeeper struct {
		cdc  codec.Codec
		perm permkeeper.Keeper
	}
)
