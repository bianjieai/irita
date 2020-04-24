package clitest

import (
	"fmt"
	"github.com/irismod/token/exported"
	"testing"

	"github.com/stretchr/testify/require"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irismod/token"

	"github.com/bianjieai/irita/app"
)

func TestIritaCLIIssueToken(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	f.ModifyTokenState()

	fooAddr := f.KeyAddress(keyFoo)

	// start irita server
	proc := f.GDStart()
	defer proc.Stop(false)

	symbol := "btc"
	name := "Bitcoin"
	minUnit := "satoshi"
	initialSupply := uint64(10000000000)
	maxSupply := uint64(20000000000)
	scale := uint8(18)
	mintable := true

	success, txInfo, stderr := f.TxIssueToken(symbol, minUnit, name, scale, initialSupply, maxSupply, mintable, fooAddr.String(), "-y")

	fmt.Println(txInfo)

	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	searchResult := f.QueryTxs(1, 50, "message.action=issue_token", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	name2 := "BITCOIN"

	success, _, stderr = f.TxEditToken(symbol, name2, maxSupply, mintable, fooAddr.String(), "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	searchResult = f.QueryTxs(1, 50, "message.action=edit_token", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	token := f.QueryToken(symbol)
	require.EqualValues(t, name2, token.Name)

	mintAmt := uint64(100)
	minter, err := sdk.AccAddressFromBech32("faa16kux6v9shcs7j8p42a2s2whdptjslkugaswx2y")
	require.NoError(t, err)

	success, _, stderr = f.TxMintToken(symbol, mintAmt, fooAddr.String(), minter.String(), "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	searchResult = f.QueryTxs(1, 50, "message.action=mint_token", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	account := f.QueryAccount(minter)
	amt := account.GetCoins().AmountOf(minUnit)
	coin, err := token.ToMainCoin(sdk.NewInt64Coin(minUnit, int64(mintAmt)))
	require.NoError(t, err)
	require.EqualValues(t, coin.Amount.TruncateInt(), amt)

	success, _, stderr = f.TxTransferTokenOwner(symbol, fooAddr.String(), minter.String(), "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	searchResult = f.QueryTxs(1, 50, "message.action=transfer_token_owner", fmt.Sprintf("message.sender=%s", fooAddr))
	require.Len(t, searchResult.Txs, 1)

	tokens := f.QueryTokenOfOwner(minter.String())
	require.NotEmpty(f.T, tokens)
	require.EqualValues(t, minter, tokens[0].GetOwner())

	// Cleanup testing directories
	f.Cleanup()
}

// TxCreateRecord is iritacli tx record create
func (f *Fixtures) TxIssueToken(symbol, minUnit, name string,
	scale uint8,
	initialSupply,
	maxSupply uint64,
	mintable bool,
	from string,
	flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx token issue --symbol=%s --name=%s --min-unit=%s --initial-supply=%d --max-supply=%d --scale=%d --mintable=%v %v --keyring-backend=test --from=%s", f.IritaCLIBinary, symbol, name, minUnit, initialSupply, maxSupply, scale, mintable, f.Flags(), from)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), "y", clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxEditToken(symbol, name string,
	maxSupply uint64,
	mintable bool,
	from string,
	flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx token edit %s --from=%s --name=%s --max-supply=%d --mintable=%v %v --keyring-backend=test", f.IritaCLIBinary, symbol, from, name, maxSupply, mintable, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxMintToken(symbol string,
	amount uint64,
	from, to string,
	flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx token mint %s --from=%s --amount=%d --to=%s %v --keyring-backend=test", f.IritaCLIBinary, symbol, from, amount, to, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), "y", clientkeys.DefaultKeyPass)
}

func (f *Fixtures) TxTransferTokenOwner(symbol string, from, to string,
	flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("%s tx token transfer %s --from=%s --to=%s %v --keyring-backend=test", f.IritaCLIBinary, symbol, from, to, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), "y", clientkeys.DefaultKeyPass)
}

func (f *Fixtures) QueryToken(symbol string) (t token.Token) {
	cmd := fmt.Sprintf("%s query token token %s --output=%s %v", f.IritaCLIBinary, symbol, "json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	cdc.MustUnmarshalJSON([]byte(out), &t)
	return
}

func (f *Fixtures) QueryTokenOfOwner(owner string) (ts []exported.TokenI) {
	cmd := fmt.Sprintf("%s query token tokens %s --output=%s %v", f.IritaCLIBinary, owner, "json", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	cdc := app.MakeCodec()
	cdc.MustUnmarshalJSON([]byte(out), &ts)
	return
}

func (f *Fixtures) ModifyTokenState() {
	genesis := f.GenesisState()[token.ModuleName]
	cdc := codec.New()

	var state token.GenesisState
	cdc.MustUnmarshalJSON(genesis, &state)

	state.Params.IssueTokenBaseFee.Amount = sdk.NewInt(30)

	bz := cdc.MustMarshalJSON(state)
	f.Override(token.ModuleName, bz)
}
