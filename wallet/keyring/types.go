package keyring

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
)

const (
	// bits of entropy to draw when creating a mnemonic
	addressSuffix = "address"
	infoSuffix    = "info"
	rootSuffix    = "ini"

	charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	clearCmd map[string]func() //create a map for storing clear funcs
)

func init() {
	clearCmd = make(map[string]func()) //Initialize it
	clearCmd["linux"] = func() {
		cmd := exec.Command("sh", "-c", `clear && printf '\033c'`) //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
	clearCmd["darwin"] = func() {
		cmd := exec.Command("sh", "-c", `clear && printf '\e[3J'`) //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
	clearCmd["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}

// encoding info
func marshalInfo(i cosmoskeyring.Info) []byte {
	return CryptoCdc.MustMarshalLengthPrefixed(i)
}

func marshalRoot(root rootInfo) []byte {
	return CryptoCdc.MustMarshalLengthPrefixed(root)
}

// decoding info
func unmarshalInfo(bz []byte) (info cosmoskeyring.Info, err error) {
	err = CryptoCdc.UnmarshalLengthPrefixed(bz, &info)
	return
}

// decoding info
func unmarshalRoot(bz []byte) (root rootInfo, err error) {
	err = CryptoCdc.UnmarshalLengthPrefixed(bz, &root)
	return
}

func randomString(l int) string {
	bytes := []byte(charset)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func clear() {
	cmd, ok := clearCmd[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                           //if we defined a clear func for that platform:
		cmd() //we execute it
	}
}
