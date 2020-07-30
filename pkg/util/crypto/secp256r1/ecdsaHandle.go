/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  ecdsaHandle
 * @Version: 1.0.0
 * @Date: 2020/7/24 10:43
 */

package secp256r1

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
		return LoadPublicKeyNotCert(pub)
	}
	if block.Type == crypto.CertType {
		return LoadPublicKey(pub)
	}
	return nil, errors.New("cert loading failed")
}

func NewEcdsaR1Handle(pub, pri string) (*ecdsaHandle, error) {
	priKey, err := LoadPrivateKey(pri)
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

	ecdsa := &ecdsaHandle{
		pubKey: pubKey,
		priKey: priKey,
	}

	return ecdsa, nil
}

func NewTransUserR1Handle(priKey *ecdsa.PrivateKey) *ecdsaHandle {

	ecdsa := &ecdsaHandle{
		pubKey: &priKey.PublicKey,
		priKey: priKey,
	}

	return ecdsa
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
