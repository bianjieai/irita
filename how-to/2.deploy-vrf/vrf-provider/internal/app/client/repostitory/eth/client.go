package eth

import (
	"bytes"
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	gethcmn "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore/keys/vrfkey"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/signatures/secp256k1"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/vrf/proof"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory/eth/contracts"
)

const CtxTimeout = 10 * time.Second

const TryGetGasPriceTimeInterval = 10 * time.Second

type Eth struct {
	contractCfgGroup *ContractCfgGroup
	contracts        *contractGroup
	bindOpts         *bindOpts

	adminVrfKey vrfkey.KeyV2
	kMaster     keystore.Master

	ethClient  *gethethclient.Client
	gethCli    *gethclient.Client
	gethRpcCli *gethrpc.Client
}

func NewEth(config *ChainConfig) (*Eth, error) {

	//pk := `{"PublicKey":"0x68e8b817b819e35cbd871ce33d8754f72347a1e44225bd5427ddcb87a44bf0ea00","vrf_key":{"address":"411efb75b302bdc8fdea29bd106eb34124e6c738","crypto":{"cipher":"aes-128-ctr","ciphertext":"1b9d18bc3142ab3eae54dd2facb6620a235fe3a64c9eb227738774a02078e848","cipherparams":{"iv":"c11da367f8794c0d5e8113b8fd5ceb2d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"0dfb59f91b1161e4ab1e5fc80b0c21fe3df510e894263f1ad655369c763ed809"},"mac":"55ae66c3426cbf1cc2191c8481f8e49b97f7c3f794a813e06c94eefe31e4bbbf"},"version":3}}`

	kMaster := keystore.New()

	adminVrfKey, err := kMaster.VRF().Import([]byte(config.VrfAdminKey), "12345678")
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, config.ChainURI)
	if err != nil {
		return nil, err
	}

	ethClient := gethethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	contractGroup, err := newContractGroup(ethClient, config.ContractCfgGroup)
	if err != nil {
		return nil, err
	}

	tmpBindOpts, err := newBindOpts(config.ContractBindOptsCfg)

	if err != nil {
		return nil, err
	}

	return &Eth{
		adminVrfKey:      adminVrfKey,
		kMaster:          kMaster,
		contractCfgGroup: config.ContractCfgGroup,
		ethClient:        ethClient,
		gethCli:          gethCli,
		gethRpcCli:       rpcClient,
		contracts:        contractGroup,
		bindOpts:         tmpBindOpts,
	}, nil
}

func (eth Eth) GetLatestHeight() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	return eth.ethClient.BlockNumber(ctx)
}

func (eth Eth) GetRandomWordsRequestedEvent(height uint64) ([]*contracts.VrfRandomWordsRequested, error) {
	address := gethcmn.HexToAddress(eth.contractCfgGroup.VRF.Addr)
	topic := eth.contractCfgGroup.VRF.Topic
	logs, err := eth.getLogs(address, topic, height, height)
	if err != nil {
		return nil, err
	}
	var vrfRqs []*contracts.VrfRandomWordsRequested
	for _, log := range logs {
		vrfRq, err := eth.contracts.VRF.ParseRandomWordsRequested(log)
		if err != nil {
			return nil, err
		}
		vrfRqs = append(vrfRqs, vrfRq)
	}

	return vrfRqs, nil
}

func (eth Eth) FulfillRandomWords(requestLog *contracts.VrfRandomWordsRequested) (string, error) {
	preSeed, err := proof.BigToSeed(requestLog.PreSeed)
	if err != nil {
		return "", err
	}
	proofResponse, rc, err := proof.GenerateProofResponseV2(
		eth.kMaster.VRF(), eth.adminVrfKey.ID(),
		proof.PreSeedDataV2{
			PreSeed:          preSeed,
			BlockHash:        requestLog.Raw.BlockHash,
			BlockNum:         requestLog.Raw.BlockNumber,
			SubId:            requestLog.SubId,
			CallbackGasLimit: requestLog.CallbackGasLimit,
			NumWords:         requestLog.NumWords,
			Sender:           requestLog.Sender,
		})
	proofParam := contracts.VRFProof{
		Pk:            proofResponse.Pk,
		Gamma:         proofResponse.Gamma,
		C:             proofResponse.C,
		S:             proofResponse.S,
		Seed:          proofResponse.Seed,
		UWitness:      proofResponse.UWitness,
		CGammaWitness: proofResponse.CGammaWitness,
		SHashWitness:  proofResponse.SHashWitness,
		ZInv:          proofResponse.ZInv,
	}
	rcParam := contracts.VRFCoordinatorRequestCommitment{
		BlockNum:         rc.BlockNum,
		SubId:            rc.SubId,
		CallbackGasLimit: rc.CallbackGasLimit,
		NumWords:         rc.NumWords,
		Sender:           rc.Sender,
	}
	result, err := eth.contracts.VRF.FulfillRandomWords(
		eth.bindOpts.vrfOpts,
		proofParam,
		rcParam,
	)
	if err != nil {
		return "", err
	}
	return result.Hash().String(), nil
}

func (eth Eth) RegisterProvingKey(key vrfkey.KeyV2) (string, error) {
	//注册 provider
	keyPoint, err := key.PublicKey.Point()
	if err != nil {
		return "", err
	}
	resultHash, err := eth.contracts.VRF.RegisterProvingKey(
		eth.bindOpts.vrfOpts,
		eth.pair(secp256k1.Coordinates(keyPoint)),
	)
	if err != nil {
		return "", err
	}
	return resultHash.Hash().String(), nil
}

func (eth Eth) StoreBlockHash(height uint64, hash common.Hash) (string, error) {
	heightBigint := new(big.Int).SetUint64(height)

	result, err := eth.contracts.VRF.StoreBlockHash(eth.bindOpts.vrfOpts, heightBigint, hash)
	if err != nil {
		return "", err
	}
	return result.Hash().String(), nil
}

func (eth Eth) GetProvider(key vrfkey.KeyV2) (bool, error) {
	_, _, providers, err := eth.contracts.VRF.GetRequestConfig(nil)
	if err != nil {
		return false, err
	}

	keyPoint, err := key.PublicKey.Point()
	if err != nil {
		return false, err
	}

	hashKey, err := eth.contracts.VRF.HashOfKey(
		nil,
		eth.pair(secp256k1.Coordinates(keyPoint)))
	if err != nil {
		return false, err
	}

	for _, provider := range providers {
		if bytes.Equal(hashKey[:], provider[:]) {
			return true, nil
		}
	}
	return false, nil
}

func (eth Eth) GetBlockHashByHeight(height uint64) (common.Hash, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	block, err := eth.ethClient.BlockByNumber(ctx, new(big.Int).SetUint64(height))
	if err != nil {
		return common.Hash{}, err
	}
	return block.Hash(), nil
}

func (eth *Eth) GetResult(hash string) (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()

	cmnHash := gethcmn.HexToHash(hash)
	result, err := eth.ethClient.TransactionReceipt(ctx, cmnHash)
	if err != nil {
		return 0, err
	}
	return result.Status, nil
}

func (eth Eth) AdminKeyProvider() vrfkey.KeyV2 {
	return eth.adminVrfKey
}

func (eth Eth) ServiceName() string {
	return "vrf"
}

func (eth Eth) Contracts() *contractGroup {
	return eth.contracts
}

func (eth *Eth) getLogs(address gethcmn.Address, topic string, fromBlock, toBlock uint64) ([]gethtypes.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []gethcmn.Address{address},
		Topics:    [][]gethcmn.Hash{{gethcrypto.Keccak256Hash([]byte(topic))}},
	}
	return eth.ethClient.FilterLogs(context.Background(), query)
}

func (eth *Eth) pair(x, y *big.Int) [2]*big.Int { return [2]*big.Int{x, y} }
