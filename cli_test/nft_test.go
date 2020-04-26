package clitest

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irismod/nft"
	"testing"

	"github.com/bianjieai/irita/app"
	"github.com/cosmos/cosmos-sdk/tests"
	"github.com/irismod/nft/exported"
	"github.com/stretchr/testify/require"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
)

func TestIritaCLINFT(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	fooAddr := f.KeyAddress(keyFoo)

	// start irita server
	proc := f.GDStart()
	defer proc.Stop(false)

	nftDenom := "service"
	tokenID := "pricing"
	tokenURI := "https://github.com/bianjieai"

	success, _, stderr := f.TxMintNFT(nftDenom, tokenID, tokenURI, fooAddr.String(), "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	searchResult := f.QueryTxs(1, 50, "message.action=mint_nft", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	//for query
	nft,err := f.QueryNFT(nftDenom,tokenID)
	require.NoError(t,err)
	require.Equal(t,tokenID,nft.GetID())
	require.Equal(t,tokenURI,nft.GetTokenURI())
	require.Equal(t,fooAddr,nft.GetOwner())

	supply := f.QueryNFTSupply(fooAddr.String(), nftDenom)
	require.Equal(t,uint64(1),supply)

	collection,err := f.QueryCollection(nftDenom)
	require.NoError(t,err)
	require.Equal(t, nftDenom,collection.Denom)
	require.Len(t,collection.NFTs,1)
	require.Equal(t,tokenID,collection.NFTs[0].GetID())

	denoms,err := f.QueryDenoms()
	require.NoError(t,err)
	require.Len(t,collection.NFTs,1)
	require.Equal(t, nftDenom,denoms[0])

	//for edit nft
	tokenURI2 := "https://github.com/bianjieai/irita/blob/master/docs/features/service-pricing.json"
	success, _, stderr = f.TxEditNFT(nftDenom, tokenID, tokenURI2, fooAddr.String(), "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	searchResult = f.QueryTxs(1, 50, "message.action=edit_nft", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	nft,err = f.QueryNFT(nftDenom,tokenID)
	require.NoError(t,err)
	require.Equal(t,nft.GetTokenURI(),tokenURI2)

	//for Transfer NFT
	barAddr := f.KeyAddress(keyBar)
	success, _, stderr = f.TxTransferNFT(nftDenom,tokenID,tokenURI2,fooAddr.String(),barAddr.String(),"-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	searchResult = f.QueryTxs(1, 50, "message.action=transfer_nft", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	owner := f.QueryOwner(barAddr.String(), nftDenom)
	require.Len(t,owner.IDCollections,1)
	require.Len(t,owner.IDCollections[0].IDs,1)
	require.Equal(t,owner.IDCollections[0].IDs[0],tokenID)

	success, _, stderr = f.TxSend(keyFoo,barAddr,sdk.NewCoin(denom, sdk.TokensFromConsensusPower(5)),"-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	success, _, stderr = f.TxBurnNFT(nftDenom,tokenID,barAddr.String(),"-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	searchResult = f.QueryTxs(1, 50, "message.action=burn_nft", fmt.Sprintf("message.sender=%s", barAddr))
	require.Len(t, searchResult.Txs, 1)

	_,err = f.QueryNFT(nftDenom,tokenID)
	require.Error(t,err)

	supply = f.QueryNFTSupply(barAddr.String(), nftDenom)
	require.Equal(t,uint64(0),supply)

	_,err = f.QueryCollection(nftDenom)
	require.Error(t,err)

	denoms,err = f.QueryDenoms()
	require.NoError(t,err)
	require.Empty(t,denoms)
}

func (f *Fixtures) TxMintNFT(denom, tokenID, tokenURI string,
	from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx nft mint %s %s --from=%s --token-uri=%s %v --keyring-backend=test", f.IritaCLIBinary, denom, tokenID, from, tokenURI, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxEditNFT(denom, tokenID, tokenURI string,
	from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx nft edit %s %s --from=%s --token-uri=%s %v --keyring-backend=test", f.IritaCLIBinary, denom, tokenID, from, tokenURI, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxTransferNFT(denom, tokenID, tokenURI string,
	from ,recipient string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx nft transfer %s %s %s --from=%s --token-uri=%s %v --keyring-backend=test", f.IritaCLIBinary, recipient,denom, tokenID, from, tokenURI, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxBurnNFT(denom, tokenID ,from string,flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx nft burn %s %s --from=%s %v --keyring-backend=test", f.IritaCLIBinary, denom, tokenID, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

func (f *Fixtures) QueryNFTSupply(addr,denom string) (supply uint64) {
	cmd := fmt.Sprintf("%s q nft supply %s --owner=%s --output=%s %v", f.IritaCLIBinary, denom, addr,"json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	cdc.MustUnmarshalJSON([]byte(out), &supply)
	return
}

func (f *Fixtures) QueryOwner(addr,denom string) (owner nft.Owner){
	cmd := fmt.Sprintf("%s q nft owner %s --denom=%s --output=%s %v", f.IritaCLIBinary, addr, denom, "json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	cdc.MustUnmarshalJSON([]byte(out), &owner)
	return
}

func (f *Fixtures) QueryCollection(denom string) (collection nft.Collection,err error){
	cmd := fmt.Sprintf("%s q nft collection %s --output=%s %v", f.IritaCLIBinary, denom,"json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	err = cdc.UnmarshalJSON([]byte(out), &collection)
	return
}

func (f *Fixtures) QueryDenoms() (denoms []string ,err error){
	cmd := fmt.Sprintf("%s q nft denoms --output=%s %v", f.IritaCLIBinary,"json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	err = cdc.UnmarshalJSON([]byte(out), &denoms)
	return
}

func (f *Fixtures) QueryNFT(denom, tokenID string) (nft exported.NFT,err error) {
	cmd := fmt.Sprintf("%s q nft token %s %s --output=%s %v", f.IritaCLIBinary, denom, tokenID, "json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	err = cdc.UnmarshalJSON([]byte(out), &nft)
	return
}
