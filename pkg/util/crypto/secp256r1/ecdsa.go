package secp256r1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
)

const publicKeyLength = 64

type ECDSASignature struct {
	R, S *big.Int
}

var (
	// curveHalfOrders contains the precomputed curve group orders halved.
	// It is used to ensure that signature' S value is lower or equal to the
	// curve group order halved. We accept only low-S signatures.
	// They are precomputed for efficiency reasons.
	curveHalfOrders = map[elliptic.Curve]*big.Int{
		elliptic.P224(): new(big.Int).Rsh(elliptic.P224().Params().N, 1),
		elliptic.P256(): new(big.Int).Rsh(elliptic.P256().Params().N, 1),
		elliptic.P384(): new(big.Int).Rsh(elliptic.P384().Params().N, 1),
		elliptic.P521(): new(big.Int).Rsh(elliptic.P521().Params().N, 1),
	}
)

// IsLow checks that s is a low-S
func IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := curveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}

	return s.Cmp(halfOrder) != 1, nil

}

func ToLowS(k *ecdsa.PublicKey, s *big.Int) (*big.Int, bool, error) {
	lowS, err := IsLowS(k, s)
	if err != nil {
		return nil, false, err
	}

	if !lowS {
		// Set s to N - s that will be then in the lower part of signature space
		// less or equal to half order
		s.Sub(k.Params().N, s)

		return s, true, nil
	}

	return s, false, nil
}

func UnmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	// Unmarshal
	sig := new(ECDSASignature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, errors.New("invalid signature, R must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, errors.New("invalid signature, S must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, R must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("invalid signature, S must be larger than zero")
	}

	return sig.R, sig.S, nil
}

func GetCurveHalfOrdersAt(c elliptic.Curve) *big.Int {
	return big.NewInt(0).Set(curveHalfOrders[c])
}

func VerifyECDSA(k *ecdsa.PublicKey, signature, digest []byte) (bool, error) {
	r, s, err := UnmarshalECDSASignature(signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	lowS, err := IsLowS(k, s)
	if err != nil {
		return false, err
	}

	if !lowS {
		return false, fmt.Errorf("Invalid S. Must be smaller than half the order [%s][%s].", s, GetCurveHalfOrdersAt(k.Curve))
	}

	return ecdsa.Verify(k, digest, r, s), nil
}

func LoadPrivateKeyByFile(file string) (*ecdsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	bl, _ := pem.Decode(b)
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from " + file)
	}
	key, err := x509.ParsePKCS8PrivateKey(bl.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key from " + file)
	}
	return key.(*ecdsa.PrivateKey), nil
}

func LoadPrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {

	bl, _ := pem.Decode([]byte(privateKey))
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from PrivateKey")
	}
	key, err := x509.ParsePKCS8PrivateKey(bl.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}
	return key.(*ecdsa.PrivateKey), nil
}

func LoadPublicKeyByFile(file string) (*ecdsa.PublicKey, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	bl, _ := pem.Decode(b)
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from " + file)
	}
	key, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key from " + file)
	}
	return key.PublicKey.(*ecdsa.PublicKey), nil
}

func LoadPublicKeyNotCert(cert string) (*ecdsa.PublicKey, error) {
	bl, _ := pem.Decode([]byte(cert))
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from Certificate")
	}

	key, err := x509.ParsePKIXPublicKey(bl.Bytes)

	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}

	return key.(*ecdsa.PublicKey), nil
}

func LoadPublicKey(cert string) (*ecdsa.PublicKey, error) {

	bl, _ := pem.Decode([]byte(cert))
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from Certificate")
	}
	key, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key from Certificate")
	}
	return key.PublicKey.(*ecdsa.PublicKey), nil
}

func SignECDSA(k *ecdsa.PrivateKey, digest []byte) (signature []byte, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}

	s, _, err = ToLowS(&k.PublicKey, s)
	if err != nil {
		return nil, err
	}

	return marshalECDSASignature(r, s)
}

func marshalECDSASignature(r, s *big.Int) ([]byte, error) {
	return asn1.Marshal(ECDSASignature{r, s})
}

func GetSHA256HASH(data string) []byte {
	bmsg := []byte(data)

	h := sha256.New()
	h.Write([]byte(bmsg))
	hash := h.Sum(nil)
	return hash

}

// ECDSAPubBytes return esdsa public key as slice
func ECDSAPubBytes(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	pubBytes := make([]byte, publicKeyLength)
	copy(pubBytes[:], pub.X.Bytes())
	copy(pubBytes[publicKeyLength/2:], pub.Y.Bytes())
	return pubBytes
}
