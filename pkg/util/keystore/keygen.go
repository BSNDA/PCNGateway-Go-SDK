package keystore

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"

	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
)

type ecdsaKeyGenerator struct {
	curve elliptic.Curve
}

func (kg *ecdsaKeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {

	privKey, err := ecdsa.GenerateKey(kg.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", kg.curve, err)
	}

	return &ecdsaPrivateKey{privKey}, nil
}

func KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {

	algor := opts.Algorithm()

	switch algor {
	case SM2:
		privKey, err := sm2.GenerateKey()
		if err != nil {
			return nil, fmt.Errorf("Failed generating SM2 key for %s", err)
		}
		return &smPrivateKey{privKey: privKey}, nil
	case ECDSAP256:
		curve := elliptic.P256()
		privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			return nil, fmt.Errorf("Failed generating ECDSA key for [%v]: [%s]", curve, err)
		}

		return &ecdsaPrivateKey{privKey}, nil
	}

	return nil, fmt.Errorf("Failed generating ")

}
