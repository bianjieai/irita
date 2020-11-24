package keyring

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/99designs/keyring"
	jose "github.com/dvsekhvalnov/jose2go"
	"github.com/mitchellh/go-homedir"
	"github.com/mtibben/percent"

	"github.com/cosmos/cosmos-sdk/client/input"
)

var (
	filenameEscape = func(s string) string {
		return percent.Encode(s, "/")
	}
	filenameUnescape = percent.Decode
)

type Keyring interface {
	keyring.Keyring
	Encrypt(bytes []byte, pass ...string) (string, error)
	Decrypt(bytes []byte) (string, error)
	Reset(password string) error
	RemoveAll() error
}

func NewFileKeyring(dir string, buf io.Reader) Keyring {
	fileDir := filepath.Join(dir, keyringFileDirName)
	return &fileKeyring{
		dir: fileDir,
		passwordFunc: func() (string, error) {
			return promptPassword(fileDir, buf)
		},
	}
}

type fileKeyring struct {
	password     string
	dir          string
	passwordFunc func() (string, error)
}

func (k fileKeyring) Get(key string) (keyring.Item, error) {
	filename, err := k.filename(key)
	if err != nil {
		return keyring.Item{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return keyring.Item{}, keyring.ErrKeyNotFound
	} else if err != nil {
		return keyring.Item{}, err
	}

	var decoded keyring.Item
	err = json.Unmarshal(bytes, &decoded)

	return decoded, err
}

func (k fileKeyring) GetMetadata(key string) (keyring.Metadata, error) {
	filename, err := k.filename(key)
	if err != nil {
		return keyring.Metadata{}, err
	}

	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return keyring.Metadata{}, errors.New("key not found")
	} else if err != nil {
		return keyring.Metadata{}, err
	}

	return keyring.Metadata{
		ModificationTime: stat.ModTime(),
	}, nil
}

func (k fileKeyring) Set(i keyring.Item) error {
	bytes, err := json.Marshal(i)
	if err != nil {
		return err
	}

	filename, err := k.filename(i.Key)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0600)
}

func (k *fileKeyring) Reset(password string) error {
	passwordHash, err := hex.DecodeString(GenHash(password))
	if err != nil {
		return err
	}

	filename := path.Join(k.dir, "/keyhash")
	if err := os.Remove(filename); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, passwordHash, 0555); err != nil {
		return err
	}
	k.password = password
	return nil
}

func (k fileKeyring) Remove(key string) error {
	filename, err := k.filename(key)
	if err != nil {
		return err
	}

	return os.Remove(filename)
}

func (k fileKeyring) RemoveAll() error {
	dir, err := k.resolveDir()
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

func (k fileKeyring) Keys() ([]string, error) {
	dir, err := k.resolveDir()
	if err != nil {
		return nil, err
	}

	var keys []string
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		keys = append(keys, filenameUnescape(f.Name()))
	}

	return keys, nil
}

func (k *fileKeyring) Encrypt(bytes []byte, pass ...string) (string, error) {
	if len(pass) == 0 {
		if err := k.unlock(); err != nil {
			return "", err
		}
	} else {
		k.password = pass[0]
	}

	return jose.Encrypt(
		string(bytes),
		jose.PBES2_HS256_A128KW,
		jose.A256GCM,
		k.password,
		jose.Headers(
			map[string]interface{}{"created": time.Now().String()},
		),
	)
}

func (k *fileKeyring) Decrypt(bytes []byte) (string, error) {
	if err := k.unlock(); err != nil {
		return "", err
	}

	payload, _, err := jose.Decode(string(bytes), k.password)
	if err != nil {
		return "", err
	}
	return payload, nil
}

func (k *fileKeyring) filename(key string) (string, error) {
	dir, err := k.resolveDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, filenameEscape(key)), nil
}

func (k *fileKeyring) resolveDir() (string, error) {
	if k.dir == "" {
		return "", fmt.Errorf("no directory provided for file keyring")
	}

	dir := k.dir

	// expand tilde for home directory
	if strings.HasPrefix(dir, "~") {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		dir = strings.Replace(dir, "~", home, 1)
	}

	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
	} else if err != nil && !stat.IsDir() {
		err = fmt.Errorf("%s is a file, not a directory", dir)
	}

	return dir, err
}

func (k *fileKeyring) unlock() error {
	if k.password == "" {
		pwd, err := k.passwordFunc()
		if err != nil {
			return err
		}
		k.password = pwd
	}
	return nil
}

func promptPassword(dir string, buf io.Reader) (string, error) {
	keyhashStored := false
	keyhashFilePath := filepath.Join(dir, "keyhash")

	var keyhash []byte

	_, err := os.Stat(keyhashFilePath)

	switch {
	case err == nil:
		keyhash, err = ioutil.ReadFile(keyhashFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to read %s: %v", keyhashFilePath, err)
		}

		keyhashStored = true

	case os.IsNotExist(err):
		keyhashStored = false

	default:
		return "", fmt.Errorf("failed to open %s: %v", keyhashFilePath, err)
	}

	failureCounter := 0

	for {
		failureCounter++
		if failureCounter > maxPassphraseEntryAttempts {
			return "", fmt.Errorf("too many failed passphrase attempts")
		}

		buf := bufio.NewReader(buf)
		pass, err := input.GetPassword("Enter root passphrase:", buf)
		if err != nil { //nolint:unparam
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}

		if keyhashStored {
			srcHash := hex.EncodeToString(keyhash)
			if !VerifyHash(srcHash, pass) {
				_, _ = fmt.Fprintln(os.Stderr, "incorrect passphrase")
				continue
			}
			return pass, nil
		}

		reEnteredPass, err := input.GetPassword("Re-enter root passphrase:", buf)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}

		if pass != reEnteredPass {
			_, _ = fmt.Fprintln(os.Stderr, "passphrases do not match")
			continue
		}

		passwordHash, err := hex.DecodeString(GenHash(pass))
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			continue
		}

		filename := path.Join(dir, "/keyhash")
		if err := ioutil.WriteFile(filename, passwordHash, 0555); err != nil {
			return "", err
		}
		return pass, nil
	}
}
