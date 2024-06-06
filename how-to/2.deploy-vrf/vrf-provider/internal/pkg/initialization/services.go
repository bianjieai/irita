package initialization

import (
	log "github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory"
	repoeth "gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory/eth"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/pkg/configs"
)

func vrfService(cfg *configs.Config, logger *log.Logger) repostitory.IChain {
	loggerEntry := logger.WithFields(log.Fields{
		"chain_name": cfg.ContractServices.VRF.ServiceName,
	})

	loggerEntry.Info(" init eth chain start")

	contractCfgGroup := repoeth.NewContractCfgGroup()

	contractCfgGroup.VRF.Addr = cfg.ContractServices.VRF.Contract.Addr
	contractCfgGroup.VRF.Topic = cfg.ContractServices.VRF.Contract.Topic
	contractCfgGroup.VRF.OptPrivKey = cfg.ContractServices.VRF.Contract.OptPrivKey

	contractBindOptsCfg := repoeth.NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = cfg.Eth.ChainID
	contractBindOptsCfg.GasLimit = cfg.Eth.GasLimit
	contractBindOptsCfg.MaxGasPrice = cfg.Eth.MaxGasPrice

	contractBindOptsCfg.VRFPrivKey = cfg.ContractServices.VRF.Contract.OptPrivKey

	ethChainCfg := repoeth.NewChainConfig()
	ethChainCfg.ContractCfgGroup = contractCfgGroup
	ethChainCfg.ContractBindOptsCfg = contractBindOptsCfg

	ethChainCfg.ChainID = cfg.Eth.ChainID
	ethChainCfg.ChainURI = cfg.Eth.URI
	ethChainCfg.VrfAdminKey = cfg.Eth.VrfAdminKey

	ethRepo, err := repoeth.NewEth(ethChainCfg)
	if err != nil {
		loggerEntry.WithFields(log.Fields{
			"err_msg": err,
		}).Fatal("failed to init chain")
	}

	return ethRepo
}
