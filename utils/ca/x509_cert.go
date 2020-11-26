package ca

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// X509Cert defines custom X509 certificate
type X509Cert struct {
	*x509.Certificate
	PrivateKey interface{}
}

func ReadX509CertFromMem(data []byte) (Cert, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	return X509Cert{cert, nil}, err
}

func (xc X509Cert) WritePrivateKeytoMem() ([]byte, error) {
	der, err := x509.MarshalPKCS8PrivateKey(xc.PrivateKey)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	}

	return pem.EncodeToMemory(block), nil
}

func (xc X509Cert) VerifyCertFromRoot(rootCert Cert) error {
	if rc, ok := rootCert.(X509Cert); ok {
		return xc.Certificate.CheckSignatureFrom(rc.Certificate)
	}
	return errors.New("can not verify x509 certificate by other algorithm certificate")
}
