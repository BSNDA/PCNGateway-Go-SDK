package keystore

import (
	"crypto"
	"github.com/pkg/errors"

	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
)

func BCCSPKeyRequestGenerate(ks bccsp.KeyStore, keyOpts bccsp.KeyGenOpts) (bccsp.Key, crypto.Signer, error) {

	key, err := KeyGen(keyOpts)
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
