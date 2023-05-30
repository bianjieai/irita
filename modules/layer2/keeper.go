package layer2

import (
	"errors"

	"github.com/bianjieai/iritamod/modules/perm"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	nftkeeper "github.com/irisnet/irismod/modules/nft/keeper"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"

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

	denom, _ := l2.nk.GetDenom(ctx, classID)
	if denom.UpdateRestricted {
		return l2.nk.TransferOwner(ctx, classID, tokenID, nfttypes.DoNotModify, nfttypes.DoNotModify, nfttypes.DoNotModify, nfttypes.DoNotModify, srcOwner, dstOwner)
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

	if denom.Creator != owner.String() {
		return errors.New("sender not the class owner")
	}

	denom.MintRestricted = mintRestricted

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

func NewPermKeeper(cdc codec.Codec, pk permkeeper.Keeper) PermKeeper {
	return PermKeeper{
		cdc:  cdc,
		perm: pk,
	}
}

func (k PermKeeper) HasL2UserRole(ctx sdk.Context, addr sdk.AccAddress) bool {
	if k.perm.IsRootAdmin(ctx, addr) {
		return true
	}

	if err := k.perm.Access(ctx, addr, perm.RoleLayer2User.Auth()); err != nil {
		return false
	}
	return true
}
