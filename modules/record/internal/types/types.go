package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/bytes"
)

// Record represents a Record
type Record struct {
	TxHash   bytes.HexBytes `json:"tx_hash" yaml:"tx_hash"`
	Contents []Content      `json:"contents" yaml:"contents"`
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
}

// Content represents a sub-record
type Content struct {
	Digest     string `json:"digest" yaml:"digest"`
	DigestAlgo string `json:"digest_algo" yaml:"digest_algo"`
	URI        string `json:"uri" yaml:"uri"`
	Meta       string `json:"meta" yaml:"meta"`
}

// NewRecord constructs a record
func NewRecord(txHash bytes.HexBytes, contents []Content, creator sdk.AccAddress) Record {
	return Record{
		TxHash:   txHash,
		Contents: contents,
		Creator:  creator,
	}
}
