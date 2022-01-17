package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibckeeper "github.com/bianjieai/tibc-go/modules/tibc/core/keeper"
	"github.com/irisnet/irismod/modules/nft/exported"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
)

var _ tibcnfttypes.NftKeeper = NFTKeeper{}

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

func WrapNftKeeper(nk nftkeeper.Keeper) NFTKeeper {
	return NFTKeeper{nk}
}

func (n NFTKeeper) MintNFT(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenData string, owner sdk.AccAddress) error {
	return n.nk.MintNFT(ctx, denomID, tokenID, tokenNm, tokenURI, "", tokenData, owner)
}

func (n NFTKeeper) BurnNFT(ctx sdk.Context, denomID, tokenID string, owner sdk.AccAddress) error {
	return n.nk.BurnNFT(ctx, denomID, tokenID, owner)
}

func (n NFTKeeper) GetNFT(ctx sdk.Context, denomID, tokenID string) (nft exported.NFT, err error) {
	return n.nk.GetNFT(ctx, denomID, tokenID)
}

func (n NFTKeeper) TransferOwner(ctx sdk.Context, denomID, tokenID, tokenNm, tokenURI, tokenData string, srcOwner, dstOwner sdk.AccAddress) error {
	return n.nk.TransferOwner(ctx, denomID, tokenID, tokenNm, tokenURI, nfttypes.DoNotModify, tokenData, srcOwner, dstOwner)
}

func (n NFTKeeper) GetDenom(ctx sdk.Context, id string) (denom nfttypes.Denom, found bool) {
	return n.nk.GetDenom(ctx, id)
}

func (n NFTKeeper) IssueDenom(ctx sdk.Context, id, name, schema, symbol string, creator sdk.AccAddress, mintRestricted, updateRestricted bool) error {
	return n.nk.IssueDenom(ctx, id, name, schema, symbol, creator, mintRestricted, updateRestricted, nfttypes.DoNotModify, nfttypes.DoNotModify, nfttypes.DoNotModify, nfttypes.DoNotModify)
}
