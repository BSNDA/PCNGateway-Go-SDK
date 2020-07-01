package keystore

import (
	"crypto"
	"crypto/elliptic"
	"github.com/pkg/errors"

	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
)

func BCCSPKeyRequestGenerate(ks bccsp.KeyStore) (bccsp.Key, crypto.Signer, error) {

	keyOpts := &ECDSAP256KeyGenOpts{Temporary: true}

	myCSP := &ecdsaKeyGenerator{curve: elliptic.P256()}

	key, err := myCSP.KeyGen(keyOpts)
	if err != nil {
		return nil, nil, err
	}

	ks.StoreKey(key)

	cspSigner, err := New(key)
	if err != nil {
		return nil, nil, errors.WithMessage(err, "Failed initializing CryptoSigner")
	}
	return key, cspSigner, nil
}
