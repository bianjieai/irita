package keyring

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
)

// CryptoCdc defines the codec required for keys and info
var CryptoCdc *codec.LegacyAmino

func init() {
	CryptoCdc = codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(CryptoCdc)
	RegisterLegacyAminoCodec(CryptoCdc)
	CryptoCdc.Seal()
}

// RegisterLegacyAminoCodec registers concrete types and interfaces on the given codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cosmoskeyring.RegisterLegacyAminoCodec(cdc)
}
