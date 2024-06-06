package keyring

import (
	"fmt"
	"io"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ghodss/yaml"
)

type rootInfo struct {
	Mnemonic     string        `json:"mnemonic"`
	MnemonicHash string        `json:"mnemonic_hash"`
	Index        uint32        `json:"index"`
	Algo         hd.PubKeyType `json:"algo"`
}

type offlineInfo struct {
	Name   string             `json:"name"`
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Path   string             `json:"path"`
	Algo   hd.PubKeyType      `json:"algo"`
}

func (l offlineInfo) GetType() cosmoskeyring.KeyType {
	return cosmoskeyring.TypeOffline
}

func (l offlineInfo) GetName() string {
	return l.Name
}

func (l offlineInfo) GetPubKey() cryptotypes.PubKey {
	return l.PubKey
}

func (l offlineInfo) GetAddress() sdk.AccAddress {
	return l.PubKey.Address().Bytes()
}

func (l offlineInfo) GetPath() (*hd.BIP44Params, error) {
	return hd.NewParamsFromPath(l.Path)
}

func (l offlineInfo) GetAlgo() hd.PubKeyType {
	return l.Algo
}

func newOfflineInfo(name, hdPath string, pub cryptotypes.PubKey, algo hd.PubKeyType) cosmoskeyring.Info {
	return &offlineInfo{
		Name:   name,
		Path:   hdPath,
		PubKey: pub,
		Algo:   algo,
	}
}

func PrintInfo(w io.Writer, keyInfo ...cosmoskeyring.Info) {
	var output []KeyOutput
	for _, info := range keyInfo {
		ko, err := bech32KeyOutput(info)
		if err != nil {
			panic(err)
		}
		output = append(output, ko)
	}
	printTextInfos(w, output)
}

func printTextInfos(w io.Writer, kos []KeyOutput) {
	out, err := yaml.Marshal(&kos)
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintln(w, string(out))
}

// Bech32KeyOutput create a KeyOutput in with "acc" Bech32 prefixes. If the
// public key is a multisig public key, then the threshold and constituent
// public keys will be added.
func bech32KeyOutput(keyInfo cosmoskeyring.Info) (KeyOutput, error) {

	registry := codectypes.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz, err := cdc.MarshalInterfaceJSON(keyInfo.GetPubKey())
	if err != nil {
		return KeyOutput{}, err
	}
	pubKey := string(bz)

	hdPath, err := keyInfo.GetPath()
	if err != nil {
		return KeyOutput{}, err
	}

	return KeyOutput{
		Name:       keyInfo.GetName(),
		Type:       keyInfo.GetType().String(),
		Address:    sdk.AccAddress(keyInfo.GetPubKey().Address().Bytes()).String(),
		PubKey:     pubKey,
		AddressIdx: hdPath.AddressIndex,
	}, nil
}

// KeyOutput defines a structure wrapping around an Info object used for output
// functionality.
type KeyOutput struct {
	Name       string `json:"name" yaml:"name"`
	Type       string `json:"type" yaml:"type"`
	Address    string `json:"address" yaml:"address"`
	PubKey     string `json:"pub_key" yaml:"pub_key"`
	AddressIdx uint32 `json:"address_idx" yaml:"address_idx"`
}
