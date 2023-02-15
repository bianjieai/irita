package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory/eth/contracts"
)

// ==================================================================================================================
// contract client group
type contractGroup struct {
	VRF *contracts.Vrf
}

func newContractGroup(ethClient *gethethclient.Client, cfgGroup *ContractCfgGroup) (*contractGroup, error) {
	vrfAddr := gethcmn.HexToAddress(cfgGroup.VRF.Addr)
	packetFilter, err := contracts.NewVrf(vrfAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return &contractGroup{
		VRF: packetFilter,
	}, nil
}

// ==================================================================================================================
// contract bind opts
type bindOpts struct {
	vrfOpts *bind.TransactOpts
}

func newBindOpts(cfg *ContractBindOptsCfg) (*bindOpts, error) {

	vrfPriv, err := gethcrypto.HexToECDSA(cfg.VRFPrivKey)
	if err != nil {
		return nil, err
	}
	vrfOpts, err := bind.NewKeyedTransactorWithChainID(vrfPriv, new(big.Int).SetUint64(cfg.ChainID))
	if err != nil {
		return nil, err
	}
	vrfOpts.GasLimit = cfg.GasLimit
	vrfOpts.GasPrice = new(big.Int).SetUint64(cfg.MaxGasPrice)

	return &bindOpts{
		vrfOpts: vrfOpts,
	}, nil
}
