package secp256k1

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/pem"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto"
)

func getPuk(pub string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pub))

	if block == nil {
		return nil, errors.New("load public key failed")
	}

	if block.Type == crypto.PublicKeyType {

		return LoadPublicKey([]byte(pub))
	}

	if block.Type == crypto.CertType {
		return LoadPublicKeyByCert([]byte(pub))

	}

	return nil, errors.New("cert loading failed")
}

func NewEcdsaK1Handle(pub, pri string) (*ecdsaK1Handle, error) {
	priKey, err := LoadPrivateKey([]byte(pri))

	if err != nil {
		return nil, errors.New("cert loading failed")
	}

	var pubKey *ecdsa.PublicKey
	if pub == "" {
		pubKey = &priKey.PublicKey
	} else {
		pubKey, err = getPuk(pub)
		if err != nil {
			return nil, errors.New("cert loading failed")
		}
	}

	ecdsa := &ecdsaK1Handle{
		pubKey: pubKey,
		priKey: priKey,
	}

	return ecdsa, nil
}

type ecdsaK1Handle struct {
	pubKey *ecdsa.PublicKey
	priKey *ecdsa.PrivateKey
}

func (e *ecdsaK1Handle) Hash(msg []byte) ([]byte, error) {

	h := sha256.New()

	h.Write([]byte(msg))
	hash := h.Sum(nil)

	return hash, nil
}

func (e *ecdsaK1Handle) Sign(digest []byte) ([]byte, error) {
	return SignECDSA(e.priKey, digest)

}

func (e *ecdsaK1Handle) Verify(sign, digest []byte) (bool, error) {
	return VerifyECDSA(e.pubKey, sign, digest)

}
