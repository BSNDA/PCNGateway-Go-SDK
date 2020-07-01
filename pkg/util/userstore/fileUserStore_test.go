package userstore

import (
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp/utils"
	"testing"
)

func TestFileUserStore_Load(t *testing.T) {

	userData := &msp.UserData{
		UserName: "abcde",
		AppCode:  "cl0006202003181926573677572",
	}

	userStore := FileUserStore{
		FilePath: "F:\\Work\\RedBaaS\\04SourceCode\\Gateway_sdk\\src\\github.com/BSNDA/PCNGateway-Go-SDK\\test\\msp",
	}

	userStore.Load(userData)

	key, err := keystore.ImportCert(userData.EnrollmentCertificate)
	if err != nil {
		t.Error(err)
	}

	fks, err := keystore.NewFileBasedKeyStore(nil, "F:\\Work\\RedBaaS\\04SourceCode\\Gateway_sdk\\src\\github.com/BSNDA/PCNGateway-Go-SDK\\test\\msp\\keystore", false)
	if err != nil {
		t.Error(err)
	}

	prikey, err := fks.GetKey(key.SKI())

	if err != nil {
		t.Error(err)
	}

	pk := keystore.GetECDSAPrivateKey(prikey)
	userData.PrivateKey = pk
	rawKey, err := utils.PrivateKeyToPEM(pk, nil)

	fmt.Println(string(rawKey))

}

func TestGetPeerName(t *testing.T) {

	name := getPemName("abcde@abc-cert.pem", "abc")
	fmt.Println(name)
}
