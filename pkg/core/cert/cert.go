package cert

import (
	"encoding/hex"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"github.com/cloudflare/cfssl/csr"
)

func GetCSRPEM(name string, ks bccsp.KeyStore) (string, error) {

	cr := GetCertificateRequest(name)

	key, cspSigner, err := keystore.BCCSPKeyRequestGenerate(ks)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	csrPEM, err := csr.Generate(cspSigner, cr)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("key:", hex.EncodeToString(key.SKI()))

	fmt.Println("csrPEM：", string(csrPEM))

	return string(csrPEM), nil
}

//Create Certificate Request
func GetCertificateRequest(name string) *csr.CertificateRequest {

	cr := &csr.CertificateRequest{}
	cr.CN = name
	cr.KeyRequest = newCfsslBasicKeyRequest()

	return cr

}

func newCfsslBasicKeyRequest() *csr.BasicKeyRequest {
	return &csr.BasicKeyRequest{A: "ecdsa", S: 256}
}
