package initialization

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/services/channels"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/configs"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/types/cache"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/tools"
)

const TypSource = "source"

func ChannelMap(cfg *configs.Config, logger *log.Logger) map[string]channels.IChannel {
	sourceChain := vrfService(cfg, logger)
	// 注册 admin Provider
	// todo
	registerAdminProvider(sourceChain, logger)
	return channelMap(cfg, sourceChain, logger)
}

func channelMap(cfg *configs.Config, sourceChain repostitory.IChain, logger *log.Logger) map[string]channels.IChannel {

	// init source chain channel
	sourceChannel := channel(cfg, sourceChain, TypSource, logger)
	sourceChannel = channels.NewWriterMW(
		sourceChannel, sourceChain.ServiceName(), logger,
		tools.DefaultHomePath, tools.DefaultCacheDirName,
		cfg.ContractServices.VRF.Cache.Filename,
	)
	channelMap := map[string]channels.IChannel{}
	if !cfg.ContractServices.VRF.Enabled {
		logger.Fatal("channel_types should be equal 1")
	}

	channelMap[sourceChain.ServiceName()] = sourceChannel

	return channelMap
}

func channel(cfg *configs.Config, sourceChain repostitory.IChain, typ string, logger *log.Logger) channels.IChannel {

	var channel channels.IChannel
	var channelErr error
	var filename string
	switch typ {
	case TypSource:
		filename = path.Join(tools.DefaultHomePath, tools.DefaultCacheDirName, cfg.ContractServices.VRF.Cache.Filename)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file does not exist, the initial height is the startHeight in the configuration
		switch typ {
		case TypSource:
			channel, channelErr = channels.NewChannel(
				sourceChain, cfg.ContractServices.VRF.Cache.StartHeight, logger)
		}

	} else {
		// If the file exists, the initial height is the latest_height in the file
		file, err := os.Open(filename)
		if err != nil {
			logger.Fatal("read cache file err: ", err)
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			logger.Fatal("read cache file err: ", err)
		}

		cacheData := &cache.Data{}
		err = json.Unmarshal(content, cacheData)
		if err != nil {
			logger.Fatal("read cache file unmarshal err: ", err)
		}
		channel, channelErr = channels.NewChannel(sourceChain, cacheData.LatestHeight, logger)
	}
	if channelErr != nil {
		logger.Fatal("failed to init channel err: ", channelErr)
	}

	return channel
}

func registerAdminProvider(chain repostitory.IChain, logger *log.Logger) {
	logger.Info("regsiter admin provider")
	found, err := chain.GetProvider(chain.AdminKeyProvider())
	if err != nil {
		logger.Fatal(err)
	}
	if found {
		return
	}
	resultHash, err := chain.RegisterProvingKey(chain.AdminKeyProvider())
	if err != nil {
		logger.Fatal(err)
	}
	for {
		txStatus, err := chain.GetResult(resultHash)
		if err == nil {
			if txStatus == 1 {
				logger.Info(" admin provider register success ")
				break
			}
			logger.Fatal(" status == 0, register error")
		}
	}
}
