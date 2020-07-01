package keystore

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
)

func KeyImport(raw interface{}) (bccsp.Key, error) {
	lowLevelKey, ok := raw.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Invalid raw material. Expected *ecdsa.PublicKey.")
	}

	return &ecdsaPublicKey{lowLevelKey}, nil
}

func ImportCert(cert []byte) (bccsp.Key, error) {
	dcert, _ := pem.Decode(cert)
	if dcert == nil {
		return nil, errors.New("Unable to decode cert bytes [%v]")
	}

	x509Cert, err := x509.ParseCertificate(dcert.Bytes)
	if err != nil {
		return nil, errors.New("Unable to parse cert from decoded bytes: %s")
	}

	pk := x509Cert.PublicKey

	return KeyImport(pk)
}
