package keyring

import (
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"

	tmbcrypt "github.com/tendermint/crypto/bcrypt"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/armor"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

const (
	blockTypePrivKey = "TENDERMINT PRIVATE KEY"

	headerType   = "type"
	addressIndex = "index"

	passwordLen = 10
)

var SigningAlgoList = keyring.SigningAlgoList{
	hd.Secp256k1,
	hd.Sm2,
}

// BcryptSecurityParameter is security parameter var, and it can be changed within the lcd test.
// Making the bcrypt security parameter a var shouldn't be a security issue:
// One can't verify an invalid key by maliciously changing the bcrypt
// parameter during a runtime vulnerability. The main security
// threat this then exposes would be something that changes this during
// runtime before the user creates their key. This vulnerability must
// succeed to update this to that same value before every subsequent call
// to the keys command in future startups / or the attacker must get access
// to the filesystem. However, with a similar threat model (changing
// variables in runtime), one can cause the user to sign a different tx
// than what they see, which is a significantly cheaper attack then breaking
// a bcrypt hash. (Recall that the nonce still exists to break rainbow tables)
// For further notes on security parameter choice, see README.md
var BcryptSecurityParameter = 12

func GenHash(text string) string {
	hash := tmhash.Sum([]byte(text))
	return hex.EncodeToString(hash)
}

func VerifyHash(srcHash, srcText string) bool {
	dstHash := hex.EncodeToString(tmhash.Sum([]byte(srcText)))
	return srcHash == dstHash
}

// Encrypt and armor the private key.
func EncryptArmorPrivKey(privKey cryptotypes.PrivKey, passphrase string, header map[string]string) string {
	saltBytes, encBytes := encryptPrivKey(privKey, passphrase)
	if header == nil {
		header = map[string]string{}
	}
	header["kdf"] = "bcrypt"
	header["salt"] = fmt.Sprintf("%X", saltBytes)
	return armor.EncodeArmor(blockTypePrivKey, header, encBytes)
}

// encrypt the given privKey with the passphrase using a randomly
// generated salt and the xsalsa20 cipher. returns the salt and the
// encrypted priv key.
func encryptPrivKey(privKey cryptotypes.PrivKey, passphrase string) (saltBytes []byte, encBytes []byte) {
	saltBytes = crypto.CRandBytes(16)
	key, err := tmbcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)

	if err != nil {
		panic(errors.Wrap(err, "generating bcrypt key from passphrase"))
	}

	key = crypto.Sha256(key) // get 32 bytes
	privKeyBytes := privKey.Bytes()

	return saltBytes, xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
}
