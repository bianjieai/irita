package eth

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/signatures/secp256k1"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/vrf/proof"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/internal/app/client/repostitory/eth/contracts"
)

const (
	chainID = 12231

	host       = "http://testnet.bianjie.ai:8545"
	optPrivKey = "3f2ca07c1f351caed872317dba6693ef917393121331fefdfa56012e1cbb1e5c"

	VRFCoordinatorAddr = "0xc632Fcba486e341Db8599bb7C07741f8a4f245Ae"
	VRFTopic           = "RandomWordsRequested(bytes32,uint256,uint256,uint64,uint16,uint32,uint32,address)"

	subId = 1
)

// getRequestConfig
func TestEthClient(t *testing.T) {

	contractCfgGroup := NewContractCfgGroup()
	contractCfgGroup.VRF.Addr = VRFCoordinatorAddr
	contractCfgGroup.VRF.Topic = VRFTopic
	contractCfgGroup.VRF.OptPrivKey = optPrivKey

	contractBindOptsCfg := NewContractBindOptsCfg()
	contractBindOptsCfg.ChainID = chainID
	contractBindOptsCfg.VRFPrivKey = optPrivKey
	contractBindOptsCfg.GasLimit = 2000000
	contractBindOptsCfg.MaxGasPrice = 1

	pk := `{"PublicKey":"0x68e8b817b819e35cbd871ce33d8754f72347a1e44225bd5427ddcb87a44bf0ea00","vrf_key":{"address":"411efb75b302bdc8fdea29bd106eb34124e6c738","crypto":{"cipher":"aes-128-ctr","ciphertext":"1b9d18bc3142ab3eae54dd2facb6620a235fe3a64c9eb227738774a02078e848","cipherparams":{"iv":"c11da367f8794c0d5e8113b8fd5ceb2d"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"0dfb59f91b1161e4ab1e5fc80b0c21fe3df510e894263f1ad655369c763ed809"},"mac":"55ae66c3426cbf1cc2191c8481f8e49b97f7c3f794a813e06c94eefe31e4bbbf"},"version":3}}`
	chainCfg := NewChainConfig()
	chainCfg.ContractCfgGroup = contractCfgGroup
	chainCfg.ContractBindOptsCfg = contractBindOptsCfg
	chainCfg.ChainURI = host
	chainCfg.ChainID = chainID
	chainCfg.VrfAdminKey = pk

	ethClient, err := NewEth(chainCfg)
	if err != nil {
		t.Fatal(err)
	}
	latestHeight, err := ethClient.GetLatestHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(latestHeight)

	vrfCfg, err := ethClient.contracts.VRF.GetConfig(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(vrfCfg)

	kMaster := keystore.New()
	myVrfkey, err := kMaster.VRF().Import([]byte(pk), "12345678")
	if err != nil {
		t.Fatal(err)
	}

	p, err := myVrfkey.PublicKey.Point()

	require.NoError(t, err)
	t.Log(p)

	//注册 privider
	result1, err := ethClient.contracts.VRF.RegisterProvingKey(
		ethClient.bindOpts.vrfOpts,
		ethClient.pair(secp256k1.Coordinates(p)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result1.Hash().String())

	_, _, providers, err := ethClient.contracts.VRF.GetRequestConfig(nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, provider := range providers {
		t.Log(hexutil.Encode(provider[:]))
	}

	//result2, err := ethClient.contracts.VRF.HashOfKey(
	//	nil,
	//	ethClient.pair(secp256k1.Coordinates(p)))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(hexutil.Encode(result2[:]))
	//cosumer := common.HexToAddress("0x4579DB44FD3A6F645194058914E0A8D5E8F20DB8")
	//resultG1, err := ethClient.contracts.VRF.AddConsumer(
	//	ethClient.bindOpts.vrfOpts,
	//	subId,
	//	cosumer,
	//)
	//require.NoError(t, err)
	//t.Log(resultG1.Hash().String())

	resultG1, err := ethClient.contracts.VRF.GetSubscription(
		nil,
		subId,
	)
	require.NoError(t, err)
	t.Log(resultG1)

	//resultG2, err := ethClient.SubscriptionConsumerAddedEvent(10602865)
	//require.NoError(t, err)
	//t.Log(resultG2)

	// 返回结果

	events, err := ethClient.GetRandomWordsRequestedEvent(10629868)
	require.NoError(t, err)

	requestLog := events[0]

	preSeed, err := proof.BigToSeed(requestLog.PreSeed)
	require.NoError(t, err)
	//requestLog.Raw.BlockHash

	proofResponse, rc, err := proof.GenerateProofResponseV2(
		kMaster.VRF(), myVrfkey.ID(),
		proof.PreSeedDataV2{
			PreSeed:          preSeed,
			BlockHash:        common.HexToHash("0x"),
			BlockNum:         requestLog.Raw.BlockNumber,
			SubId:            requestLog.SubId,
			CallbackGasLimit: requestLog.CallbackGasLimit,
			NumWords:         requestLog.NumWords,
			Sender:           requestLog.Sender,
		})
	require.NoError(t, err)

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
	t.Log("requestId: ", requestLog.RequestId.String())
	t.Log("blockNum: ", rcParam.BlockNum)
	t.Log("SubId: ", rcParam.SubId)
	t.Log("CallbackGasLimit: ", rcParam.CallbackGasLimit)
	t.Log("NumWords: ", rcParam.NumWords)
	t.Log("Sender: ", rcParam.Sender)

	result3, err := ethClient.contracts.VRF.FulfillRandomWords(
		ethClient.bindOpts.vrfOpts,
		proofParam,
		rcParam,
	)
	require.NoError(t, err)
	t.Log(result3.Hash().String())

}
