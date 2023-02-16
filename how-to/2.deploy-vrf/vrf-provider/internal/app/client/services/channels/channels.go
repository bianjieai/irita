package channels

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/domain"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory"
	typeserr "gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/types/errors"
)

const RetryTimeout = 10 * time.Second
const RetryTimes = 20

type IChannel interface {
	Relay() error
	IsNotRelay() bool
	Context() *domain.Context
}

type Channel struct {
	source repostitory.IChain

	context *domain.Context

	logger *log.Logger
}

func NewChannel(source repostitory.IChain, startHeight uint64, logger *log.Logger) (IChannel, error) {

	return &Channel{
		logger:  logger,
		source:  source,
		context: domain.NewContext(startHeight, source.ServiceName()),
	}, nil
}

func (channel *Channel) Relay() error {
	return channel.relay()
}

func (channel *Channel) IsNotRelay() bool {
	curHeight := channel.Context().Height()
	latestHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		return false
	}

	if curHeight < latestHeight {
		return true
	}

	return false
}

func (channel *Channel) Context() *domain.Context {
	return channel.context
}

func (channel *Channel) relay() error {
	logger := channel.logger.WithFields(log.Fields{
		"start_height": channel.Context().Height(),
		"option":       "relay",
	})
	logger.Info()
	latestHeight, err := channel.source.GetLatestHeight()
	if err != nil {
		logger.Error("failed to get latest height")
		return typeserr.ErrGetLatestHeight
	}
	if latestHeight <= channel.Context().Height() {
		logger.Info("the current height cannot be relayed yet")
		return typeserr.ErrNotProduced
	}
	//1. 获取 Event
	events, err := channel.source.GetRandomWordsRequestedEvent(channel.Context().Height())
	if err != nil {
		logger.Info("failed to get random words request event")
		return typeserr.ErrGetRandomWordsRequestedEvent
	}

	for _, event := range events {
		// 2. 先查下合约中是否有当前区块高度的 hash
		// 如果没有写入到合约
		storeBlockHashResultHash, err := channel.source.StoreBlockHash(event.Raw.BlockNumber, event.Raw.BlockHash)
		if err != nil {
			logger.Info("failed to store block hash")
			return typeserr.ErrStoreBlockHash
		}
		logger.WithField("storeBlockHashResultHash", storeBlockHashResultHash).Info()
		if err := channel.reTryEthResult(storeBlockHashResultHash, 0); err != nil {
			logger.WithField("err_msg", err).Error("failed to StoreBlockHash: retry: ", err)
			return err
		}
		// 3. 向链上发交易
		fulfillRandomWordsResultHash, err := channel.source.FulfillRandomWords(event)
		if err != nil {
			logger.Info("failed to get random words requested event")
			return typeserr.ErrFulfillRandomWords
		}
		logger.WithField("fulfillRandomWordsResultHash", fulfillRandomWordsResultHash).Info()
		if err := channel.reTryEthResult(fulfillRandomWordsResultHash, 0); err != nil {
			logger.WithField("err_msg", err).Error("failed to FulfillRandomWords : retry: ", err)
			return err
		}
	}

	channel.Context().IncrHeight()
	return nil
}

func (channel *Channel) reTryEthResult(hash string, n uint64) error {
	channel.logger.Infof("retry %d time", n)
	if n == RetryTimes {
		return fmt.Errorf("retry %d times and return error", RetryTimes)
	}
	txStatus, err := channel.source.GetResult(hash)
	if err != nil {
		channel.logger.Info("re-request result ")
		time.Sleep(RetryTimeout)
		return channel.reTryEthResult(hash, n+1)
	}
	if txStatus == 0 {
		channel.logger.WithFields(log.Fields{
			"hash": hash,
			"flag": "result_error",
		}).Warning("re-request result is false")
		return nil
	}
	return nil
}
