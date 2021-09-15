package keys

import (
	"io"

	"github.com/bianjieai/irita/wallet/keyring"
)

type KeybaseGenerator func(buf io.Reader) (keyring.Keystore, error)
