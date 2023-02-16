package repostitory

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore/keys/vrfkey"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory/eth/contracts"
)

type IChain interface {
	GetLatestHeight() (uint64, error)
	GetRandomWordsRequestedEvent(height uint64) ([]*contracts.VrfRandomWordsRequested, error)
	GetProvider(key vrfkey.KeyV2) (bool, error)

	FulfillRandomWords(requestLog *contracts.VrfRandomWordsRequested) (string, error)
	StoreBlockHash(height uint64, hash common.Hash) (string, error)

	RegisterProvingKey(key vrfkey.KeyV2) (string, error)

	GetResult(hash string) (uint64, error)

	ServiceName() string
	AdminKeyProvider() vrfkey.KeyV2
}
