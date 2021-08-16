package keystore

import (
	"encoding/hex"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/userstore"
	kscert "github.com/BSNDA/bsn-sdk-crypto/keystore/cert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoadKey(t *testing.T) {

	ks, _ := NewFileBasedKeyStore(nil, "F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp/keystore", false)
	us := userstore.NewUserStore("F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp")

	user := &msp.UserData{
		UserName: "sdktest",
		AppCode:  "app0006202004071529586812466",
	}

	us.Load(user)

	LoadKey(user, ks, enum.AppAlgorithmType_R1)

	fmt.Println(string(user.EnrollmentCertificate))
	fmt.Println(user.PrivateKey)
}

func TestCertService_SaveUserCert(t *testing.T) {
	certFiles, err := GetAllFiles("D:/certs/pem", "pem")
	//ks, _ := NewFileBasedKeyStore(nil, "D:/certs/pem/msp/keystore", false)

	for _, certFile := range certFiles {
		bytes, err := ioutil.ReadFile(certFile)
		filename := strings.Split(certFile, ";")
		println(filename[1])
		bb := strings.Split(filename[1], "@")
		println(bb)
		if err != nil {

		}
		cert := string(bytes)
		println(cert)
		key, err := kscert.ImportCert(bytes)
		if err != nil {

		}
		prikeyname := hex.EncodeToString(key.SKI())
		aa, _ := ioutil.ReadFile("D:/certs/pem/msp/keystore/" + prikeyname + "_sk")
		println(aa)

	}
	println(certFiles)
	println(err)
}
