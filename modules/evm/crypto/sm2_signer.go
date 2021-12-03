package crypto

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Sm2Signer struct {
	types.EIP155Signer
}

func NewSm2Signer(chainId *big.Int) types.Signer {
	return Sm2Signer{types.NewEIP155Signer(chainId)}
}

func (s Sm2Signer) Sender(tx *types.Transaction) (common.Address, error) {
	return s.EIP155Signer.Sender(tx)
}

func (s Sm2Signer) ChainID() *big.Int {
	return s.EIP155Signer.ChainID()
}

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func (s Sm2Signer) Hash(tx *types.Transaction) common.Hash {
	return rlpHash([]interface{}{
		tx.Nonce(),
		tx.GasPrice(),
		tx.Gas(),
		tx.To(),
		tx.Value(),
		tx.Data(),
		s.ChainID(), uint(0), uint(0),
	})
}

func (s Sm2Signer) Equal(signer types.Signer) bool {
	return s.EIP155Signer.Equal(signer)
}

var _ types.Signer = &Sm2Signer{}

func decodeSignature(sig []byte) (r, s, v *big.Int) {
	//if len(sig) != crypto.SignatureLength {
	//	panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	//}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	//v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v
}

func (s Sm2Signer) SignatureValues(tx *types.Transaction, sig []byte) (R, S, V *big.Int, err error) {
	// because it indicates that the chain ID was not specified in the tx.
	if tx.ChainId().Sign() != 0 && tx.ChainId().Cmp(s.ChainID()) != 0 {
		return nil, nil, nil, types.ErrInvalidChainId
	}
	R, S, _ = decodeSignature(sig)
	//V = big.NewInt(int64(sig[64]))
	return R, S, V, nil
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])
	return h
}
