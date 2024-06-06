package vrf

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/gethwrappers/generated/vrf_coordinator_v2"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/signatures/secp256k1"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/vrf/proof"
)

func TestKeyStoreMaster(t *testing.T) {
	keyStoreMaster := keystore.New()
	vrfkey, err := keyStoreMaster.VRF().Create()
	require.NoError(t, err)
	keyJSON, err := keyStoreMaster.VRF().Export(vrfkey.ID(), "12345678")
	require.NoError(t, err)
	vrfKey1, err := keyStoreMaster.VRF().Delete(vrfkey.ID())
	require.NoError(t, err)
	t.Log(vrfKey1.PublicKey.String())
	t.Log("=============")
	t.Log(string(keyJSON))
	t.Log("=============")
	vrfKey2, err := keyStoreMaster.VRF().Import(keyJSON, "12345678")
	require.NoError(t, err)
	t.Log(vrfKey2.PublicKey.String())
	p, err := vrfKey2.PublicKey.Point()
	require.NoError(t, err)
	publicProvingKeyX, publicProvingKeyY := secp256k1.Coordinates(p)
	t.Log(publicProvingKeyX, publicProvingKeyY)
}

func pair(x, y *big.Int) [2]*big.Int { return [2]*big.Int{x, y} }

func TestFulfillmentCost(t *testing.T) {

	keyStoreMaster := keystore.New()
	vrfkey, err := keyStoreMaster.VRF().Create()
	require.NoError(t, err)

	carolContractAddress := common.HexToAddress("0x0CEEBB3BCF6CCF2F3523DCFD7865B2B0D752DC32")

	nw := 1
	gasRequested := 50000

	requestLog := &vrf_coordinator_v2.VRFCoordinatorV2RandomWordsRequested{
		RequestId: big.NewInt(11),
		PreSeed:   big.NewInt(1),
		Raw: types.Log{
			BlockNumber: 100,
			TxHash:      common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
		},
	}

	s, err := proof.BigToSeed(requestLog.PreSeed)
	require.NoError(t, err)

	subId := uint64(1)

	proof, rc, err := proof.GenerateProofResponseV2(
		keyStoreMaster.VRF(), vrfkey.ID(),
		proof.PreSeedDataV2{
			PreSeed:          s,
			BlockHash:        requestLog.Raw.BlockHash,
			BlockNum:         requestLog.Raw.BlockNumber,
			SubId:            subId,
			CallbackGasLimit: uint32(gasRequested),
			NumWords:         uint32(nw),
			Sender:           carolContractAddress,
		})
	require.NoError(t, err)
	t.Log(proof)
	t.Log(rc)

}
