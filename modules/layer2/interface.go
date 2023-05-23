package layer2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"

	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
)

type (
	PermKeeper struct {
		cdc  codec.Codec
		perm permkeeper.Keeper
	}

	NftKeeper struct {
		cdc codec.Codec
		nk  nftkeeper.Keeper
	}

	NftToken struct {
		Owner sdk.AccAddress
	}

	NftClass struct {
		ClassId        string
		Owner          string
		MintRestricted bool
	}
)

func (c NftClass) GetID() string            { return c.ClassId }
func (c NftClass) GetOwner() string         { return c.Owner }
func (c NftClass) GetMintRestricted() bool  { return c.MintRestricted }
func (t NftToken) GetOwner() sdk.AccAddress { return t.Owner }
