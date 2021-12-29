package utils

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	etherminthd "github.com/tharsis/ethermint/crypto/hd"
)

func SetEthermintSupportedAlgorithms() {
	etherminthd.SupportedAlgorithms = keyring.SigningAlgoList{etherminthd.EthSecp256k1, hd.Secp256k1, hd.Sm2}
	etherminthd.SupportedAlgorithmsLedger = keyring.SigningAlgoList{etherminthd.EthSecp256k1, hd.Secp256k1, hd.Sm2}

}
