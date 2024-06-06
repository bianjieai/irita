package keyring

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
)

const (
	mnemonicEntropySize        = 256
	RootName                   = "root"
	keyringFileDirName         = "keyring-file"
	maxPassphraseEntryAttempts = 3
)

// New creates a new instance of a keyring.
// Keyring ptions can be applied when generating the new instance.
// Available backends are "file".
func New(rootDir string, userInput io.Reader, opts ...cosmoskeyring.Option) (Keystore, error) {
	db := NewFileKeyring(rootDir, userInput)

	options := cosmoskeyring.Options{
		SupportedAlgos:       cosmoskeyring.SigningAlgoList{hd.Sm2, hd.Secp256k1},
		SupportedAlgosLedger: cosmoskeyring.SigningAlgoList{hd.Sm2, hd.Secp256k1},
	}

	for _, optionFn := range opts {
		optionFn(&options)
	}

	return Keystore{db, options}, nil
}

type Keystore struct {
	db      Keyring
	options cosmoskeyring.Options
}

// Init create a wallet instance and  produce a mnemonic
func (ks Keystore) Init(algo cosmoskeyring.SignatureAlgo, r *bufio.Reader, w *bufio.Writer) (cosmoskeyring.Info, error) {
	if !ks.isSupportedSigningAlgo(algo) {
		return nil, cosmoskeyring.ErrUnsupportedSigningAlgo
	}

	// read entropy seed straight from crypto.Rand and convert to mnemonic
	entropySeed, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		return nil, err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return nil, err
	}

	addressIdx := uint32(0)

	if err = ks.writeRoot(mnemonic, addressIdx, algo.Name()); err != nil {
		return nil, err
	}

	info, err := ks.addKey(RootName, mnemonic, ks.createHDPath(addressIdx), algo)
	if err != nil {
		return nil, err

	}

	if err = ks.confirmMnemonic(mnemonic, r, w); err != nil {
		_ = ks.db.RemoveAll()
	}

	return info, err
}

// HasInit determine if the wallet has been initialized
func (ks Keystore) HasInit() bool {
	key := rootKey()
	bs, err := ks.db.Get(string(key))
	if err == nil && len(bs.Data) > 0 {
		return true
	}
	return false
}

// Recover recover a wallet instance by a mnemonic
func (ks Keystore) Recover(mnemonic string, algo cosmoskeyring.SignatureAlgo) (cosmoskeyring.Info, error) {
	if !ks.isSupportedSigningAlgo(algo) {
		return nil, cosmoskeyring.ErrUnsupportedSigningAlgo
	}

	addressIdx := uint32(0)
	if err := ks.writeRoot(mnemonic, addressIdx, algo.Name()); err != nil {
		return nil, err
	}

	info, err := ks.addKey(RootName, mnemonic, ks.createHDPath(addressIdx), algo)
	if err != nil {
		return nil, err
	}

	return info, err
}

// UpdateRoot update a wallet root password by old password or mnemonic
func (ks Keystore) UpdateRoot(mnemonic, newPwd string) error {
	root, err := ks.getRoot()
	if err != nil {
		return err
	}

	if len(mnemonic) == 0 {
		mnemonic, err = ks.db.Decrypt([]byte(root.Mnemonic))
		if err != nil {
			return err
		}
	}

	if !VerifyHash(root.MnemonicHash, mnemonic) {
		return errors.New("invalid mnemonic or old encryptRoot passphrase")
	}

	mnemonicEncrypted, err := ks.db.Encrypt([]byte(mnemonic), newPwd)
	if err != nil {
		return err
	}

	newRoot := rootInfo{
		Mnemonic:     mnemonicEncrypted,
		MnemonicHash: root.MnemonicHash,
		Index:        root.Index,
		Algo:         root.Algo,
	}

	key := rootKey()
	value := marshalRoot(newRoot)

	if err := ks.db.Set(keyring.Item{
		Key:  string(key),
		Data: value,
	}); err != nil {
		return err
	}

	return ks.db.Reset(newPwd)
}

// NewKey create a new key
func (ks Keystore) NewKey(name string) (cosmoskeyring.Info, error) {
	if ks.Has(name) {
		return nil, fmt.Errorf("%s has exist", name)
	}

	root, err := ks.getVerifiedRoot()
	if err != nil {
		return nil, err
	}

	algo, err := cosmoskeyring.NewSigningAlgoFromString(string(root.Algo), SigningAlgoList)
	if err != nil {
		return nil, err
	}

	addressIdx := root.Index + 1
	if err := ks.writeRoot(root.Mnemonic, addressIdx, root.Algo); err != nil {
		return nil, err
	}

	hdPath := ks.createHDPath(addressIdx)
	return ks.addKey(name, root.Mnemonic, hdPath, algo)
}

// Has determine if a key exists
func (ks Keystore) Has(uid string) bool {
	key := infoKey(uid)

	bs, err := ks.db.Get(string(key))
	if err == nil && len(bs.Data) > 0 {
		return true
	}
	return false
}

// Key return a key information by key name
func (ks Keystore) Key(uid string) (cosmoskeyring.Info, error) {
	key := infoKey(uid)

	bs, err := ks.db.Get(string(key))
	if err != nil {
		return nil, err
	}

	if len(bs.Data) == 0 {
		return nil, fmt.Errorf("key %s not found in keybase", uid)
	}
	return unmarshalInfo(bs.Data)
}

// Key return a key information by the address
func (ks Keystore) KeyByAddress(address sdk.Address) (cosmoskeyring.Info, error) {
	ik, err := ks.db.Get(addrHexKeyAsString(address))
	if err != nil {
		return nil, err
	}

	if len(ik.Data) == 0 {
		return nil, fmt.Errorf("key with address %s not found", address)
	}

	bs, err := ks.db.Get(string(ik.Data))
	if err != nil {
		return nil, err
	}

	return unmarshalInfo(bs.Data)
}

// Key return all the key information
func (ks Keystore) List() ([]cosmoskeyring.Info, error) {
	var res []cosmoskeyring.Info

	keys, err := ks.db.Keys()
	if err != nil { //nolint:unparam
		return nil, err
	}

	sort.Strings(keys)

	for _, key := range keys {
		if strings.HasSuffix(key, infoSuffix) {
			rawInfo, err := ks.db.Get(key)
			if err != nil {
				return nil, err
			}

			if len(rawInfo.Data) == 0 {
				return nil, fmt.Errorf("key: %s not found", key)
			}

			info, err := unmarshalInfo(rawInfo.Data)
			if err != nil {
				return nil, err
			}

			res = append(res, info)
		}
	}
	return res, nil
}

// Export export a key private by a random password
func (ks Keystore) Export(key string) (armor, pwd string, err error) {
	var info cosmoskeyring.Info

	addr, err := sdk.AccAddressFromBech32(key)
	if err == nil {
		info, err = ks.KeyByAddress(addr)
		if err != nil {
			return "", "", err
		}
	} else {
		info, err = ks.Key(key)
		if err != nil {
			return "", "", err
		}
	}

	algo, err := cosmoskeyring.NewSigningAlgoFromString(string(info.GetAlgo()), SigningAlgoList)
	if err != nil {
		return "", "", err
	}

	root, err := ks.getVerifiedRoot()
	if err != nil {
		return "", "", err
	}

	bip44Params, err := info.GetPath()
	if err != nil {
		return "", "", err
	}

	hdPath := ks.createHDPath(bip44Params.AddressIndex)
	// create master key
	derivedPriv, err := algo.Derive()(root.Mnemonic, "", hdPath)
	if err != nil {
		return "", "", err
	}

	priv := algo.Generate()(derivedPriv)
	header := map[string]string{
		headerType:   string(info.GetAlgo()),
		addressIndex: fmt.Sprintf("%d", bip44Params.AddressIndex),
	}
	passphrase := randomString(passwordLen)
	return EncryptArmorPrivKey(priv, passphrase, header), passphrase, nil
}

func (ks Keystore) addKey(uid, mnemonic, hdPath string, algo cosmoskeyring.SignatureAlgo) (cosmoskeyring.Info, error) {
	// create master key and derive first key for keyring
	derivedPriv, err := algo.Derive()(mnemonic, "", hdPath)
	if err != nil {
		return nil, err
	}

	privKey := algo.Generate()(derivedPriv)

	return ks.writeLocalKey(uid, hdPath, privKey, algo.Name())
}

func (ks Keystore) writeLocalKey(name, hdPath string, priv cryptotypes.PrivKey, algo hd.PubKeyType) (cosmoskeyring.Info, error) {
	// encrypt private key using keyring
	pub := priv.PubKey()

	info := newOfflineInfo(name, hdPath, pub, algo)
	if name == RootName {
		return info, nil
	}

	if err := ks.writeInfo(info); err != nil {
		return nil, err
	}

	return info, nil
}

func (ks Keystore) writeRoot(mnemonic string, addressIdx uint32, algo hd.PubKeyType) error {
	mnemonicEncrypted, err := ks.db.Encrypt([]byte(mnemonic))
	if err != nil {
		return err
	}

	root := rootInfo{
		Mnemonic:     mnemonicEncrypted,
		MnemonicHash: GenHash(mnemonic),
		Index:        addressIdx,
		Algo:         algo,
	}

	key := rootKey()
	value := marshalRoot(root)
	return ks.db.Set(keyring.Item{
		Key:  string(key),
		Data: value,
	})
}

func (ks Keystore) writeInfo(info cosmoskeyring.Info) error {
	// write the info by key
	key := infoKey(info.GetName())
	serializedInfo := marshalInfo(info)

	exists, err := ks.existsInDb(info)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("public key already exist in keybase")
	}

	if err = ks.db.Set(keyring.Item{
		Key:  string(key),
		Data: serializedInfo,
	}); err != nil {
		return err
	}

	return ks.db.Set(keyring.Item{
		Key:  addrHexKeyAsString(info.GetAddress()),
		Data: key,
	})

}

func (ks Keystore) getVerifiedRoot() (rootInfo, error) {
	root, err := ks.getRoot()
	if err != nil {
		return rootInfo{}, err
	}

	mnemonic, err := ks.db.Decrypt([]byte(root.Mnemonic))
	if err != nil {
		return rootInfo{}, err
	}

	if !VerifyHash(root.MnemonicHash, mnemonic) {
		return rootInfo{}, errors.New("invalid mnemonic or old encryptRoot passphrase")
	}

	return rootInfo{
		Mnemonic: mnemonic,
		Index:    root.Index,
		Algo:     root.Algo,
	}, nil
}

func (ks Keystore) getRoot() (rootInfo, error) {
	ik, err := ks.db.Get(string(rootKey()))
	if err != nil {
		return rootInfo{}, fmt.Errorf("key with address %s not found", RootName)
	}
	if len(ik.Data) == 0 {
		return rootInfo{}, fmt.Errorf("key with address %s not found", RootName)
	}
	return unmarshalRoot(ik.Data)
}

func (ks Keystore) existsInDb(info cosmoskeyring.Info) (bool, error) {
	if _, err := ks.db.Get(addrHexKeyAsString(info.GetAddress())); err == nil {
		return true, nil // address lookup succeeds - info exists
	} else if err != keyring.ErrKeyNotFound {
		return false, err // received unexpected error - returns error
	}

	if _, err := ks.db.Get(string(infoKey(info.GetName()))); err == nil {
		return true, nil // uid lookup succeeds - info exists
	} else if err != keyring.ErrKeyNotFound {
		return false, err // received unexpected error - returns
	}

	// both lookups failed, info does not exist
	return false, nil
}

func (ks Keystore) isSupportedSigningAlgo(algo cosmoskeyring.SignatureAlgo) bool {
	return ks.options.SupportedAlgos.Contains(algo)
}

func (ks Keystore) createHDPath(idx uint32) string {
	return hd.CreateHDPath(sdk.CoinType, 0, idx).String()
}

func (ks Keystore) confirmMnemonic(mnemonic string, r *bufio.Reader, w *bufio.Writer) error {
	prompt := "\n**Important** Write down this mnemonic phrase in a safe place. It is the only way to recover your account if you ever forget your password\n"
	prompt += "*******************************************************************************************************************************************************"
	prompt += "\n\n" + mnemonic + "\n\n"
	prompt += "*******************************************************************************************************************************************************"
	prompt += "\nHave you written it down? The mnemonic will be verified in the next step. "

	ok, err := input.GetConfirmation(prompt, r, w)
	if err != nil || !ok {
		return errors.New("use cancel the operator")
	}

	//clear terminal
	clear()

	reEnterMnemonic, err := input.GetString("Enter your mnemonic:", r)
	if err != nil || len(reEnterMnemonic) == 0 {
		return errors.New("mnemonic is invalid")
	}

	root, err := ks.getRoot()
	if err != nil {
		return err
	}

	if !VerifyHash(root.MnemonicHash, reEnterMnemonic) {
		return errors.New("incorrect mnemonic")
	}

	//clear terminal
	clear()
	return nil
}

func rootKey() []byte {
	return []byte(fmt.Sprintf("%s.%s", RootName, rootSuffix))
}

func infoKey(name string) []byte {
	return []byte(fmt.Sprintf("%s.%s", name, infoSuffix))
}

func addrHexKeyAsString(address sdk.Address) string {
	return fmt.Sprintf("%s.%s", hex.EncodeToString(address.Bytes()), addressSuffix)
}
