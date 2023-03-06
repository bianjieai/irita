package keystore

import (
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore/keys/vrfkey"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore/store"
	"gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/utils"
)

type Master interface {
	VRF() VRF
}

type master struct {
	vrf *vrf
}

func New() Master {
	return newMaster()
}

func newMaster() *master {
	km := &keyManager{
		orm:          store.NewMemory(),
		scryptParams: utils.DefaultScryptParams,
	}
	return &master{
		vrf: newVRFKeyStore(km),
	}
}

func (ks *master) VRF() VRF {
	return ks.vrf
}

type keyManager struct {
	orm store.MemoryDAO

	scryptParams utils.ScryptParams
}

func (km *keyManager) save(id string, key vrfkey.KeyV2) error {
	return km.orm.Write(id, "", key)
}

func (km *keyManager) safeAddKey(unknownKey vrfkey.KeyV2) error {

	// save keyring to DB
	err := km.save(unknownKey.ID(), unknownKey)
	if err != nil {
		return err
	}
	return nil
}

// caller must hold lock!
func (km *keyManager) safeRemoveKey(unknownKey vrfkey.KeyV2) (err error) {
	return km.orm.Delete(unknownKey.ID(), "")
}
