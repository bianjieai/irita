package wallet

import (
	"os"
)

const (
	FlagHome    = "home"
	FlagKeyAlgo = "algo"
	flagType    = "type"
	flagRecover = "recover"

	TypePassphrase = "passphrase"
	TypeMnemonic   = "mnemonic"
)

var (
	DefaultWalletHome = os.ExpandEnv("$HOME/.iritawallet")
)
