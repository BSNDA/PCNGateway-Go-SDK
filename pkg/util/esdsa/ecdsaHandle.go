package ecdsa

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/pem"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
)

func getPuk(pub string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pub))

	fmt.Println(block.Type)

	if block.Type == "PUBLIC KEY" {

		return LoadPublicKeyNotCert(pub)
	}

	if block.Type == "CERTIFICATE" {
		return LoadPublicKey(pub)

	}

	return nil, errors.New("cert loading failed")
}

func NewEcdsaR1Handle(pub, pri string) (*ecdsaHandle, error) {

	pubKey, err := getPuk(pub)

	if err != nil {
		return nil, errors.New("cert loading failed")
	}

	priKey, err := LoadPrivateKey(pri)

	if err != nil {
		return nil, errors.New("cert loading failed")
	}

	ecdsa := &ecdsaHandle{
		pubKey: pubKey,
		priKey: priKey,
	}

	return ecdsa, nil
}

type ecdsaHandle struct {
	pubKey *ecdsa.PublicKey
	priKey *ecdsa.PrivateKey
}

func (e *ecdsaHandle) Hash(msg []byte) ([]byte, error) {

	h := sha256.New()

	h.Write([]byte(msg))
	hash := h.Sum(nil)

	return hash, nil
}

func (e *ecdsaHandle) Sign(digest []byte) ([]byte, error) {
	return SignECDSA(e.priKey, digest)

}

func (e *ecdsaHandle) Verify(sign, digest []byte) (bool, error) {
	return VerifyECDSA(e.pubKey, sign, digest)

}
