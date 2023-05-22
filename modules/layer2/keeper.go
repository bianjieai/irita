package layer2

import (
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"

	layer2nft "github.com/bianjieai/iritamod/modules/layer2/types"
)

func NewNftKeeper(cdc codec.Codec, nk nftkeeper.Keeper) NftKeeper {
	return NftKeeper{
		cdc: cdc,
		nk:  nk,
	}
}

func (l2 NftKeeper) SaveNFT(ctx sdk.Context,
	classID,
	tokenID,
	tokenNm,
	tokenURI,
	tokenUriHash,
	tokenData string,
	receiver sdk.AccAddress) error {
	return l2.nk.MintNFT(ctx, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData, receiver)
}

func (l2 NftKeeper) UpdateNFT(ctx sdk.Context,
	classID,
	tokenID,
	tokenNm,
	tokenURI,
	tokenUriHash,
	tokenData string,
	owner sdk.AccAddress) error {
	return l2.nk.EditNFT(ctx, classID, tokenID, tokenNm, tokenURI, tokenUriHash, tokenData, owner)
}

func (l2 NftKeeper) RemoveNFT(ctx sdk.Context,
	classID,
	tokenID string,
	owner sdk.AccAddress) error {
	return l2.nk.BurnNFT(ctx, classID, tokenID, owner)
}

func (l2 NftKeeper) TransferNFT(ctx sdk.Context,
	classID,
	tokenID string,
	srcOwner,
	dstOwner sdk.AccAddress) error {
	nft, err := l2.nk.GetNFT(ctx, classID, tokenID)
	if err != nil {
		return err
	}

	return l2.nk.TransferOwner(ctx, classID, tokenID, nft.GetName(), nft.GetURI(), nft.GetURIHash(), nft.GetData(), srcOwner, dstOwner)
}

func (l2 NftKeeper) TransferClass(ctx sdk.Context,
	classID string,
	srcOwner,
	dstOwner sdk.AccAddress) error {
	return l2.nk.TransferDenomOwner(ctx, classID, srcOwner, dstOwner)
}

func (l2 NftKeeper) UpdateClassMintRestricted(ctx sdk.Context,
	classID string,
	mintRestricted bool,
	owner sdk.AccAddress) error {
	denom, exist := l2.nk.GetDenom(ctx, classID)
	if !exist {
		return errors.New("class not found")
	}

	denom.MintRestricted = mintRestricted
	denom.Creator = owner.String()

	return l2.nk.UpdateDenom(ctx, denom)
}

func (l2 NftKeeper) GetClass(ctx sdk.Context,
	classID string) (layer2nft.Class, error) {
	denom, exist := l2.nk.GetDenom(ctx, classID)
	if !exist {
		return nil, errors.New("class not found")
	}

	return NftClass{
		ClassId:        denom.Id,
		Owner:          denom.Creator,
		MintRestricted: denom.MintRestricted,
	}, nil
}

func (l2 NftKeeper) GetNFT(ctx sdk.Context,
	classID,
	tokenID string) (layer2nft.NFT, error) {
	token, err := l2.nk.GetNFT(ctx, classID, tokenID)
	if err != nil {
		return nil, err
	}

	return NftToken{
		Owner: token.GetOwner(),
	}, nil
}
