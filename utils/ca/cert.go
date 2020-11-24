package ca

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/tjfoc/gmsm/sm2"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/algo"
	ed25519util "github.com/tendermint/tendermint/crypto/ed25519"
	tmsm2 "github.com/tendermint/tendermint/crypto/sm2"
)

type Cert interface {
	WritePrivateKeytoMem() ([]byte, error)
	VerifyCertFromRoot(rootCert Cert) error
}

func ReadCertificateFromMem(data []byte) (Cert, error) {
	switch algo.Algo {
	case algo.SM2:
		return ReadSM2CertFromMem(data)
	default:
		return ReadX509CertFromMem(data)
	}
}

func VerifyCertFromRoot(cert, rootCert Cert) error {
	return cert.VerifyCertFromRoot(rootCert)
}

// GetPubkeyFromCert gets the pubkey from certificate
func GetPubkeyFromCert(cert Cert) (crypto.PubKey, error) {
	switch c := cert.(type) {
	case Sm2Cert:
		expectedPubKeyAlgo := c.Certificate.PublicKeyAlgorithm
		pub, ok := c.Certificate.PublicKey.(*ecdsa.PublicKey)
		if !ok || expectedPubKeyAlgo != sm2.ECDSA {
			return nil, UnexpectedPubKeyAlgo("ECDSA", c.Certificate.PublicKey)
		}
		switch pub.Curve {
		case sm2.P256Sm2():
			sm2Pub := sm2.PublicKey{
				Curve: pub.Curve,
				X:     pub.X,
				Y:     pub.Y,
			}

			compPubkey := sm2.Compress(&sm2Pub)
			var pubKey tmsm2.PubKeySm2
			copy(pubKey[:], compPubkey)
			return pubKey, nil
		default:
			return nil, UnexpectedPubKeyAlgo("SM2", c.Certificate.PublicKey)
		}
	case X509Cert:
		expectedPubKeyAlgo := c.Certificate.PublicKeyAlgorithm
		pub, ok := c.Certificate.PublicKey.(ed25519.PublicKey)
		if !ok || expectedPubKeyAlgo != x509.Ed25519 {
			return nil, UnexpectedPubKeyAlgo(expectedPubKeyAlgo.String(), c.Certificate.PublicKey)
		}
		pubKey := make(ed25519util.PubKey, ed25519util.PubKeySize)
		copy(pubKey[:], pub)
		return pubKey, nil
	default:
		return nil, errors.New("unsupported algorithm type")
	}
}

func UnexpectedPubKeyAlgo(expected string, pubkey interface{}) error {
	return fmt.Errorf("x509: signature algorithm specifies an %s public key, but have public key of type %T", expected, pubkey)
}
