package client

import (
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/services"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/configs"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/initialization"
)

func Serve(cfg *configs.Config) {

	logger := initialization.Logger(cfg)
	logger.Info("1. service init relayers ")
	channelMap := initialization.ChannelMap(cfg, logger)
	logger.Info("2. service init listener ")
	listener := services.NewListener(channelMap, logger)
	logger.Info("3. service start & crontab start")

	logger.Fatal(listener.Listen())
}
