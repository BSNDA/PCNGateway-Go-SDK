package secp256k1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/ethereum/go-ethereum/common/math"
	"io"
	"math/big"
	"time"
)

const ecPrivKeyVersion = 1

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
		SECP256K1():     new(big.Int).Rsh(SECP256K1().Params().N, 1),
	}
)

type ECDSASignature struct {
	R, S *big.Int
}

type pkcs8Info struct {
	Version             int
	PrivateKeyAlgorithm []asn1.ObjectIdentifier
	PrivateKey          []byte
}

type pkcs8 struct {
	Version    int
	Algo       pkix.AlgorithmIdentifier
	PrivateKey []byte
	// optional attributes omitted.
}

type ecPrivateKey struct {
	Version       int
	PrivateKey    []byte
	NamedCurveOID asn1.ObjectIdentifier `asn1:"optional,explicit,tag:0"`
	PublicKey     asn1.BitString        `asn1:"optional,explicit,tag:1"`
}

type pkixPublicKey struct {
	Algo      pkix.AlgorithmIdentifier
	BitString asn1.BitString
}

type certificate struct {
	Raw                asn1.RawContent
	TBSCertificate     tbsCertificate
	SignatureAlgorithm pkix.AlgorithmIdentifier
	SignatureValue     asn1.BitString
}

type tbsCertificate struct {
	Raw                asn1.RawContent
	Version            int `asn1:"optional,explicit,default:0,tag:0"`
	SerialNumber       *big.Int
	SignatureAlgorithm pkix.AlgorithmIdentifier
	Issuer             asn1.RawValue
	Validity           validity
	Subject            asn1.RawValue
	PublicKey          publicKeyInfo
	UniqueId           asn1.BitString   `asn1:"optional,tag:1"`
	SubjectUniqueId    asn1.BitString   `asn1:"optional,tag:2"`
	Extensions         []pkix.Extension `asn1:"optional,explicit,tag:3"`
}

type validity struct {
	NotBefore, NotAfter time.Time
}

var (
	oidNamedCurveP224 = asn1.ObjectIdentifier{1, 3, 132, 0, 33}
	oidNamedCurveP256 = asn1.ObjectIdentifier{1, 2, 840, 10045, 3, 1, 7}
	oidNamedCurveP384 = asn1.ObjectIdentifier{1, 3, 132, 0, 34}
	oidNamedCurveP521 = asn1.ObjectIdentifier{1, 3, 132, 0, 35}

	oidNamedCurveS256 = asn1.ObjectIdentifier{1, 3, 132, 0, 10}
)

var oidPublicKeyECDSA = asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}

func NewSecp256k1Key() (*ecdsa.PrivateKey, error) {

	secp256k1 := SECP256K1()

	privKey, err := ecdsa.GenerateKey(secp256k1, rand.Reader)

	return privKey, err

}

func PrivateKeyToPEM(k *ecdsa.PrivateKey) ([]byte, error) {

	// get the oid for the curve
	oidNamedCurve, ok := oidFromNamedCurve(k.Curve)
	if !ok {
		return nil, errors.New("unknown elliptic curve")
	}

	// based on https://golang.org/src/crypto/x509/sec1.go
	privateKeyBytes := k.D.Bytes()
	paddedPrivateKey := make([]byte, (k.Curve.Params().N.BitLen()+7)/8)
	copy(paddedPrivateKey[len(paddedPrivateKey)-len(privateKeyBytes):], privateKeyBytes)
	// omit NamedCurveOID for compatibility as it's optional
	asn1Bytes, err := asn1.Marshal(ecPrivateKey{
		Version:    1,
		PrivateKey: paddedPrivateKey,
		PublicKey:  asn1.BitString{Bytes: elliptic.Marshal(k.Curve, k.X, k.Y)},
	})

	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}

	var pkcs8Key pkcs8Info
	pkcs8Key.Version = 0
	pkcs8Key.PrivateKeyAlgorithm = make([]asn1.ObjectIdentifier, 2)
	pkcs8Key.PrivateKeyAlgorithm[0] = oidPublicKeyECDSA
	pkcs8Key.PrivateKeyAlgorithm[1] = oidNamedCurve
	pkcs8Key.PrivateKey = asn1Bytes

	pkcs8Bytes, err := asn1.Marshal(pkcs8Key)
	if err != nil {
		return nil, fmt.Errorf("error marshaling EC key to asn1 [%s]", err)
	}
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: pkcs8Bytes,
		},
	), nil

}

// PublicKeyToPEM marshals a public key to the pem format
func PublicKeyToPEM(k *ecdsa.PublicKey) ([]byte, error) {

	PubASN1, err := MarshalPKIXPublicKey(k)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: PubASN1,
		},
	), nil
}

func MarshalPKIXPublicKey(pub *ecdsa.PublicKey) ([]byte, error) {
	var publicKeyBytes []byte
	var publicKeyAlgorithm pkix.AlgorithmIdentifier
	var err error

	if publicKeyBytes, publicKeyAlgorithm, err = marshalPublicKey(pub); err != nil {
		return nil, err
	}

	pkix := pkixPublicKey{
		Algo: publicKeyAlgorithm,
		BitString: asn1.BitString{
			Bytes:     publicKeyBytes,
			BitLength: 8 * len(publicKeyBytes),
		},
	}

	ret, _ := asn1.Marshal(pkix)
	return ret, nil
}

func LoadPublicKey(publicKey []byte) (*ecdsa.PublicKey, error) {

	bl, _ := pem.Decode(publicKey)
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from Certificate")
	}

	key, err := ParsePKIXPublicKey(bl.Bytes)

	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}

	return key.(*ecdsa.PublicKey), nil
}

func LoadPublicKeyByCert(cert []byte) (*ecdsa.PublicKey, error) {

	bl, _ := pem.Decode([]byte(cert))
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from Certificate")
	}

	key, err := ParsePublicKeyByCert(bl.Bytes)

	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}

	return key.(*ecdsa.PublicKey), nil
}

func LoadPrivateKey(privateKey []byte) (*ecdsa.PrivateKey, error) {

	bl, _ := pem.Decode(privateKey)
	if bl == nil {
		return nil, errors.New("failed to decode PEM block from PrivateKey")
	}
	key, err := ParsePKCS8PrivateKey(bl.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse private key from PrivateKey")
	}
	return key.(*ecdsa.PrivateKey), nil
}

func ParsePKCS8PrivateKey(der []byte) (key interface{}, err error) {
	var privKey pkcs8
	if _, err := asn1.Unmarshal(der, &privKey); err != nil {
		return nil, err
	}
	bytes := privKey.Algo.Parameters.FullBytes
	namedCurveOID := new(asn1.ObjectIdentifier)
	if _, err := asn1.Unmarshal(bytes, namedCurveOID); err != nil {
		namedCurveOID = nil
	}
	key, err = parseECPrivateKey(namedCurveOID, privKey.PrivateKey)
	if err != nil {
		return nil, errors.New("x509: failed to parse EC private key embedded in PKCS#8: " + err.Error())
	}
	return key, nil
}

func parseECPrivateKey(namedCurveOID *asn1.ObjectIdentifier, der []byte) (key *ecdsa.PrivateKey, err error) {
	var privKey ecPrivateKey
	if _, err := asn1.Unmarshal(der, &privKey); err != nil {
		return nil, errors.New("x509: failed to parse EC private key: " + err.Error())
	}
	if privKey.Version != ecPrivKeyVersion {
		return nil, fmt.Errorf("x509: unknown EC private key version %d", privKey.Version)
	}

	var curve elliptic.Curve
	if namedCurveOID != nil {
		curve = namedCurveFromOID(*namedCurveOID)
	} else {
		curve = namedCurveFromOID(privKey.NamedCurveOID)
	}
	if curve == nil {
		return nil, errors.New("x509: unknown elliptic curve")
	}

	k := new(big.Int).SetBytes(privKey.PrivateKey)
	curveOrder := curve.Params().N
	if k.Cmp(curveOrder) >= 0 {
		return nil, errors.New("x509: invalid elliptic curve private key value")
	}
	priv := new(ecdsa.PrivateKey)
	priv.Curve = curve
	priv.D = k

	privateKey := make([]byte, (curveOrder.BitLen()+7)/8)

	// Some private keys have leading zero padding. This is invalid
	// according to [SEC1], but this code will ignore it.
	for len(privKey.PrivateKey) > len(privateKey) {
		if privKey.PrivateKey[0] != 0 {
			return nil, errors.New("x509: invalid private key length")
		}
		privKey.PrivateKey = privKey.PrivateKey[1:]
	}

	// Some private keys remove all leading zeros, this is also invalid
	// according to [SEC1] but since OpenSSL used to do this, we ignore
	// this too.
	copy(privateKey[len(privateKey)-len(privKey.PrivateKey):], privKey.PrivateKey)
	priv.X, priv.Y = curve.ScalarBaseMult(privateKey)

	return priv, nil
}

type publicKeyInfo struct {
	Raw       asn1.RawContent
	Algorithm pkix.AlgorithmIdentifier
	PublicKey asn1.BitString
}

func ParsePKIXPublicKey(derBytes []byte) (pub interface{}, err error) {
	var pki publicKeyInfo
	if rest, err := asn1.Unmarshal(derBytes, &pki); err != nil {
		return nil, err
	} else if len(rest) != 0 {
		return nil, errors.New("x509: trailing data after ASN.1 of public-key")
	}
	return parsePublicKey(&pki)
}

func parsePublicKey(keyData *publicKeyInfo) (interface{}, error) {
	asn1Data := keyData.PublicKey.RightAlign()
	paramsData := keyData.Algorithm.Parameters.FullBytes
	namedCurveOID := new(asn1.ObjectIdentifier)
	rest, err := asn1.Unmarshal(paramsData, namedCurveOID)
	if err != nil {
		return nil, err
	}
	if len(rest) != 0 {
		return nil, errors.New("x509: trailing data after ECDSA parameters")
	}
	namedCurve := namedCurveFromOID(*namedCurveOID)
	if namedCurve == nil {
		return nil, errors.New("x509: unsupported elliptic curve")
	}
	x, y := elliptic.Unmarshal(namedCurve, asn1Data)
	if x == nil {
		return nil, errors.New("x509: failed to unmarshal elliptic curve point")
	}
	pub := &ecdsa.PublicKey{
		Curve: namedCurve,
		X:     x,
		Y:     y,
	}
	return pub, nil
}

func ParsePublicKeyByCert(certBytes []byte) (pub interface{}, err error) {
	var cert certificate
	rest, err := asn1.Unmarshal(certBytes, &cert)
	if err != nil {
		return nil, err
	}
	if len(rest) > 0 {
		return nil, asn1.SyntaxError{Msg: "trailing data"}
	}

	publicKey, err := parsePublicKey(&cert.TBSCertificate.PublicKey)

	return publicKey, err
}

func namedCurveFromOID(oid asn1.ObjectIdentifier) elliptic.Curve {
	switch {
	case oid.Equal(oidNamedCurveP224):
		return elliptic.P224()
	case oid.Equal(oidNamedCurveP256):
		return elliptic.P256()
	case oid.Equal(oidNamedCurveP384):
		return elliptic.P384()
	case oid.Equal(oidNamedCurveP521):
		return elliptic.P521()
	case oid.Equal(oidNamedCurveS256):
		return SECP256K1()
	}
	return nil
}

func marshalPublicKey(pub *ecdsa.PublicKey) (publicKeyBytes []byte, publicKeyAlgorithm pkix.AlgorithmIdentifier, err error) {
	publicKeyBytes = elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	oid, ok := oidFromNamedCurve(pub.Curve)
	if !ok {
		return nil, pkix.AlgorithmIdentifier{}, errors.New("x509: unsupported elliptic curve")
	}
	publicKeyAlgorithm.Algorithm = oidPublicKeyECDSA
	var paramBytes []byte
	paramBytes, err = asn1.Marshal(oid)
	if err != nil {
		return
	}
	publicKeyAlgorithm.Parameters.FullBytes = paramBytes

	return publicKeyBytes, publicKeyAlgorithm, nil
}

func oidFromNamedCurve(curve elliptic.Curve) (asn1.ObjectIdentifier, bool) {
	switch curve {
	case elliptic.P224():
		return oidNamedCurveP224, true
	case elliptic.P256():
		return oidNamedCurveP256, true
	case elliptic.P384():
		return oidNamedCurveP384, true
	case elliptic.P521():
		return oidNamedCurveP521, true
	case SECP256K1():
		return oidNamedCurveS256, true
	}
	return nil, false
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

func GetCurveHalfOrdersAt(c elliptic.Curve) *big.Int {
	return big.NewInt(0).Set(curveHalfOrders[c])
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

func IsLowS(k *ecdsa.PublicKey, s *big.Int) (bool, error) {
	halfOrder, ok := curveHalfOrders[k.Curve]
	if !ok {
		return false, fmt.Errorf("curve not recognized [%s]", k.Curve)
	}

	return s.Cmp(halfOrder) != 1, nil

}

func CreateCert(key *ecdsa.PrivateKey, Subject pkix.Name) ([]byte, error) {

	//todo 未实现
	serialNumber := make([]byte, 20)
	_, err := io.ReadFull(rand.Reader, serialNumber)
	if err != nil {
		return nil, errors.New("")
	}
	serialNumber[0] &= 0x7F
	template := &x509.Certificate{
		SerialNumber: new(big.Int).SetBytes(serialNumber), //new(big.Int).SetBytes([]byte(csr.SerialNumber)),//
		Subject:      Subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * 365 * 24 * time.Hour),
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)

	if err != nil {
		return nil, err
	}

	certpem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICAET", Bytes: cert})

	return certpem, nil

}

func Hash(msg []byte) ([]byte, error) {

	h := sha256.New()

	h.Write([]byte(msg))
	hash := h.Sum(nil)

	return hash, nil
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}

func FromECDSA(priv *ecdsa.PrivateKey) []byte {
	if priv == nil {
		return nil
	}
	return math.PaddedBigBytes(priv.D, priv.Params().BitSize/8)
}
