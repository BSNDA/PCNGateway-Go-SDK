package cert

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto/eth"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/helpers"
	"github.com/tjfoc/gmsm/sm2"
	"net"
	"net/mail"
	"net/url"
)

func GetCSRPEM(name string, algorithmType enum.App_AlgorithmType, ks bccsp.KeyStore) (string, bccsp.Key, error) {

	cr := &csr.CertificateRequest{}
	cr.CN = name
	cr.KeyRequest = newCfsslBasicKeyRequest()
	cr.Names = append(cr.Names, csr.Name{
		OU: "client",
		O:  "Bsn",
	})
	var keyOpts bccsp.KeyGenOpts
	keyOpts = &keystore.ECDSAP256KeyGenOpts{Temporary: true}

	if algorithmType == enum.AppAlgorithmType_SM2 {
		cr.KeyRequest = newSm2BasicKeyRequest()
		keyOpts = &keystore.SM2KeyGenOpts{Temporary: true}
	}

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(ks, keyOpts)

	if err != nil {
		return "", nil, err
	}

	var csrPEM []byte

	if algorithmType == enum.AppAlgorithmType_SM2 {
		csrPEM, err = GenerateSM2(cspSigner, cr)
	} else {
		csrPEM, err = csr.Generate(cspSigner, cr)
	}
	if err != nil {
		return "", nil, err
	}

	return string(csrPEM), key, nil
}

//Create Certificate Request
func GetCertificateRequest(name string) *csr.CertificateRequest {

	cr := &csr.CertificateRequest{}
	cr.CN = name
	cr.KeyRequest = newCfsslBasicKeyRequest()

	return cr

}

func newCfsslBasicKeyRequest() csr.KeyRequest {
	return &csr.BasicKeyRequest{A: "ecdsa", S: 256}
}

func newSm2BasicKeyRequest() csr.KeyRequest {
	return &csr.BasicKeyRequest{A: "sm2", S: 256}
}

func NewUser(userType enum.App_AlgorithmType) (interface{}, []byte, error) {

	switch userType {
	case enum.AppAlgorithmType_SM2:
		key, err := sm2.GenerateKey()
		if err != nil {
			return nil, nil, err
		}
		keyByte, err := sm2.WritePrivateKeytoMem(key, nil)
		if err != nil {
			return nil, nil, err
		}
		return key, keyByte, nil

	case enum.AppAlgorithmType_K1:
		key, err := eth.NewSecp256k1Key()
		if err != nil {
			return nil, nil, err
		}
		keyByte, err := eth.PrivateKeyToPEM(key)
		if err != nil {
			return nil, nil, err
		}
		return key, keyByte, nil
	}

	return nil, nil, errors.New("Not implemented")
}

func GetUserKey(keyBytes []byte, userType enum.App_AlgorithmType) (interface{}, error) {
	switch userType {
	case enum.AppAlgorithmType_SM2:
		return sm2.ReadPrivateKeyFromMem(keyBytes, nil)

	case enum.AppAlgorithmType_K1:
		return eth.LoadPrivateKey(keyBytes)
	}

	return nil, errors.New("Not implemented")
}

// Generate creates a new CSR from a CertificateRequest structure and
// an existing key. The KeyRequest field is ignored.
func GenerateSM2(priv crypto.Signer, req *csr.CertificateRequest) (csr []byte, err error) {
	sigAlgo := helpers.SignerAlgo(priv)
	var sm2SigAlgo sm2.SignatureAlgorithm
	if sigAlgo == x509.UnknownSignatureAlgorithm {
		sm2SigAlgo = sm2.SM2WithSM3
		//return nil, cferr.New(cferr.PrivateKeyError, cferr.Unavailable)
	}

	var tpl = sm2.CertificateRequest{
		Subject:            req.Name(),
		SignatureAlgorithm: sm2SigAlgo,
	}

	for i := range req.Hosts {
		if ip := net.ParseIP(req.Hosts[i]); ip != nil {
			tpl.IPAddresses = append(tpl.IPAddresses, ip)
		} else if email, err := mail.ParseAddress(req.Hosts[i]); err == nil && email != nil {
			tpl.EmailAddresses = append(tpl.EmailAddresses, email.Address)
		} else if uri, err := url.ParseRequestURI(req.Hosts[i]); err == nil && uri != nil {
			//tpl.URIs = append(tpl.URIs, uri)
		} else {
			tpl.DNSNames = append(tpl.DNSNames, req.Hosts[i])
		}
	}

	if req.CA != nil {
		err = appendCAInfoToCSRSm2(req.CA, &tpl)
		if err != nil {
			err = errors.New("")
			return
		}
	}

	csr, err = sm2.CreateCertificateRequest(rand.Reader, &tpl, priv)
	if err != nil {
		//log.Fatalf("failed to generate a CSR: %v", err)
		err = errors.New("")
		return
	}
	block := pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csr,
	}

	csr = pem.EncodeToMemory(&block)
	return
}

type BasicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

// appendCAInfoToCSR appends CAConfig BasicConstraint extension to a CSR
func appendCAInfoToCSRSm2(reqConf *csr.CAConfig, csr *sm2.CertificateRequest) error {
	pathlen := reqConf.PathLength
	if pathlen == 0 && !reqConf.PathLenZero {
		pathlen = -1
	}
	val, err := asn1.Marshal(BasicConstraints{true, pathlen})

	if err != nil {
		return err
	}

	csr.ExtraExtensions = []pkix.Extension{
		{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 19},
			Value:    val,
			Critical: true,
		},
	}

	return nil
}
