package store

import "gitlab.bianjie.ai/avata/contracts/vrf-provider/chainlink/core/services/keystore/keys/vrfkey"

// Use memory as storage, use with caution in build environment
type MemoryDAO struct {
	store map[string]vrfkey.KeyV2
}

func NewMemory() MemoryDAO {
	return MemoryDAO{
		store: make(map[string]vrfkey.KeyV2),
	}
}

func (m MemoryDAO) Write(name, password string, store vrfkey.KeyV2) error {
	m.store[name] = store
	return nil
}

func (m MemoryDAO) Read(name, password string) (vrfkey.KeyV2, error) {
	return m.store[name], nil
}

// ReadMetadata read a key information from the local store
func (m MemoryDAO) ReadMetadata(name string) (store vrfkey.KeyV2, err error) {
	return m.store[name], nil
}

func (m MemoryDAO) Delete(name, password string) error {
	delete(m.store, name)
	return nil
}

func (m MemoryDAO) Has(name string) bool {
	_, ok := m.store[name]
	return ok
}
