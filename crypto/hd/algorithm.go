package hd

import (
	cosmoshd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	ethd "github.com/evmos/ethermint/crypto/hd"
)

const (
	// Sm2Type defines the ECDSA sm2 used on Ethereum
	Sm2Type = cosmoshd.Sm2Type
)

var (
	// SupportedAlgorithms defines the list of signing algorithms used on Ethermint:
	//  - eth_secp256k1 (Ethereum)
	//  - secp256k1 (Tendermint)
	SupportedAlgorithms = keyring.SigningAlgoList{ethd.EthSecp256k1, cosmoshd.Secp256k1, cosmoshd.Sm2}
	// SupportedAlgorithmsLedger defines the list of signing algorithms used on Ethermint for the Ledger device:
	//  - eth_secp256k1 (Ethereum)
	//  - secp256k1 (Tendermint)
	SupportedAlgorithmsLedger = keyring.SigningAlgoList{ethd.EthSecp256k1, cosmoshd.Secp256k1}
)

// KeyringOption defines a function keys options for the ethereum Secp256k1 curve.
// It supports eth_secp256k1, secp256k1 and sm2 keys for accounts.
func KeyringOption() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = SupportedAlgorithms
		options.SupportedAlgosLedger = SupportedAlgorithmsLedger
	}
}

// SetSupportedAlgorithms sets the supported signing algorithms on Ethermint
func SetSupportedAlgorithms() {
	ethd.SupportedAlgorithms = SupportedAlgorithms
	ethd.SupportedAlgorithmsLedger = SupportedAlgorithmsLedger
}
