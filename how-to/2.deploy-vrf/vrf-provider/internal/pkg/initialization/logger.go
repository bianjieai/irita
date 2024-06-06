package initialization

import (
	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/configs"
)

func Logger(cfg *configs.Config) *log.Logger {
	logger := log.New()
	if cfg.App.Env == "prod" {
		logger.SetFormatter(&log.JSONFormatter{})
	} else {
		logger.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}
	switch cfg.App.LogLevel {
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
	return logger
}
