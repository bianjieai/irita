package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto/tmhash"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	testAddress = sdk.AccAddress(tmhash.SumTruncated([]byte("test-address")))
	testDenom   = sdk.DefaultBondDenom
	testAmount  = uint64(1000)

	emptyAddress = sdk.AccAddress{}
)

// TestMsgMintRoute tests Route for MsgMint
func TestMsgMintRoute(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress, testAddress)
	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgMintType tests Type for MsgMint
func TestMsgMintType(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress, testAddress)
	require.Equal(t, TypeMsgMint, msg.Type())
}

// TestMsgMintValidation tests ValidateBasic for MsgMint
func TestMsgMintValidation(t *testing.T) {
	testMsgs := []*MsgMint{
		NewMsgMint(testAmount, testAddress, testAddress),  // valid msg
		NewMsgMint(testAmount, testAddress, emptyAddress), // missing operator address
		NewMsgMint(0, testAddress, testAddress),           // amount must be greater than 0
		NewMsgMint(testAmount, emptyAddress, testAddress), // missing recipient address
	}

	testCases := []struct {
		msg     *MsgMint
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "amount must be greater than 0"},
		{testMsgs[3], false, "missing recipient address"},
	}

	for i, tc := range testCases {
		err := tc.msg.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "Msg %d failed: %v", i, err)
		} else {
			require.Error(t, err, "Invalid Msg %d passed: %s", i, tc.errMsg)
		}
	}
}

// TestMsgMintGetSignBytes tests GetSignBytes for MsgMint
func TestMsgMintGetSignBytes(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress, testAddress)
	res := msg.GetSignBytes()

	expected := `{"type":"irita/opb/MsgMint","value":{"amount":"1000","operator":"cosmos1hjppmlx4fgtnpsya0pzqyg7el9qrq5lw58dd9x","recipient":"cosmos1hjppmlx4fgtnpsya0pzqyg7el9qrq5lw58dd9x"}}`
	require.Equal(t, expected, string(res))
}

// TestMsgMintGetSigners tests GetSigners for MsgMint
func TestMsgMintGetSigners(t *testing.T) {
	msg := NewMsgMint(testAmount, testAddress, testAddress)
	res := msg.GetSigners()

	expected := "[BC821DFCD54A1730C09D78440223D9F9403053EE]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}

// TestMsgReclaimRoute tests Route for MsgReclaim
func TestMsgReclaimRoute(t *testing.T) {
	msg := NewMsgReclaim(testDenom, testAddress, testAddress)
	require.Equal(t, RouterKey, msg.Route())
}

// TestMsgReclaimType tests Type for MsgReclaim
func TestMsgReclaimType(t *testing.T) {
	msg := NewMsgReclaim(testDenom, testAddress, testAddress)
	require.Equal(t, TypeMsgReclaim, msg.Type())
}

// TestMsgReclaimValidation tests ValidateBasic for MsgReclaim
func TestMsgReclaimValidation(t *testing.T) {
	invalidDenom := "invalid+denom"

	testMsgs := []*MsgReclaim{
		NewMsgReclaim(testDenom, testAddress, testAddress),    // valid msg
		NewMsgReclaim(testDenom, testAddress, emptyAddress),   // missing operator address
		NewMsgReclaim(invalidDenom, testAddress, testAddress), // invalid denom
		NewMsgReclaim(testDenom, emptyAddress, testAddress),   // missing recipient address
	}

	testCases := []struct {
		msg     *MsgReclaim
		expPass bool
		errMsg  string
	}{
		{testMsgs[0], true, ""},
		{testMsgs[1], false, "missing operator address"},
		{testMsgs[2], false, "invalid denom"},
		{testMsgs[3], false, "missing recipient address"},
	}

	for i, tc := range testCases {
		err := tc.msg.ValidateBasic()
		if tc.expPass {
			require.NoError(t, err, "Msg %d failed: %v", i, err)
		} else {
			require.Error(t, err, "Invalid Msg %d passed: %s", i, tc.errMsg)
		}
	}
}

// TestMsgReclaimGetSignBytes tests GetSignBytes for MsgReclaim
func TestMsgReclaimGetSignBytes(t *testing.T) {
	msg := NewMsgReclaim(testDenom, testAddress, testAddress)
	res := msg.GetSignBytes()

	expected := `{"type":"irita/opb/MsgReclaim","value":{"denom":"stake","operator":"cosmos1hjppmlx4fgtnpsya0pzqyg7el9qrq5lw58dd9x","recipient":"cosmos1hjppmlx4fgtnpsya0pzqyg7el9qrq5lw58dd9x"}}`
	require.Equal(t, expected, string(res))
}

// TestMsgReclaimGetSigners tests GetSigners for MsgReclaim
func TestMsgReclaimGetSigners(t *testing.T) {
	msg := NewMsgReclaim(testDenom, testAddress, testAddress)
	res := msg.GetSigners()

	expected := "[BC821DFCD54A1730C09D78440223D9F9403053EE]"
	require.Equal(t, expected, fmt.Sprintf("%v", res))
}
