package evm

import (
	"fmt"
	"math/big"

	etherminttypes "github.com/evmos/ethermint/types"
)

var (
	EIP155ChainID = "1223"
)

func init() {
	etherminttypes.InjectChainIDParser(parseChainID)
}

func parseChainID(_ string) (*big.Int, error) {
	eip155ChainID, ok := new(big.Int).SetString(EIP155ChainID, 10)
	if !ok {
		return nil, fmt.Errorf("invalid chain-id: %s" + EIP155ChainID)
	}
	return eip155ChainID, nil
}
