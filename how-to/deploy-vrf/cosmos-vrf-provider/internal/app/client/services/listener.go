package services

import (
	"context"
	"sync"
	"time"

	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/services/channels"

	log "github.com/sirupsen/logrus"
)

const DefaultTimeout = 10

type IListener interface {
	Listen() error
}

type Listener struct {
	channelMap map[string]channels.IChannel

	ctxMap sync.Map
	logger *log.Logger
}

func NewListener(
	channelMap map[string]channels.IChannel,
	logger *log.Logger) IListener {
	listener := &Listener{
		channelMap: channelMap,
		ctxMap:     sync.Map{},
		logger:     logger,
	}
	return listener
}

func (listener *Listener) Listen() error {

	// 启动N个goroutine去处理
	for chainName := range listener.channelMap {
		ctx, cancel := context.WithCancel(context.Background())
		listener.ctxMap.Store(chainName, cancel)
		go listener.start(ctx, chainName)
	}
	listener.ctxMap.Range(listener.walk)
	select {}
}

func (listener *Listener) start(ctx context.Context, chainName string) {

	for {
		select {
		case <-ctx.Done():
			listener.logger.WithFields(log.Fields{
				"chain_name": chainName,
			}).Info("canceled")
			return
		default:
			if !listener.channelMap[chainName].IsNotRelay() {
				time.Sleep(DefaultTimeout * time.Second)
			} else {
				err := listener.channelMap[chainName].Relay()
				if err != nil {
					time.Sleep(DefaultTimeout * time.Second)
				}
			}

		}
	}
}

func (listener *Listener) cancelCtx(locality string) {
	if value, ok := listener.ctxMap.Load(locality); ok {
		cancel := value.(context.CancelFunc)
		cancel()
		listener.ctxMap.Delete(locality)
	}
}

func (listener *Listener) walk(key, value interface{}) bool {
	listener.logger.WithFields(log.Fields{"chain": key}).Info("start")
	return true
}
