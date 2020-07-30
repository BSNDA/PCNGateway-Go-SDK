package cert

import (
	"encoding/hex"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"testing"
)

func TestGetCSRPEM(t *testing.T) {

	name := "abc"
	fks, err := keystore.NewFileBasedKeyStore(nil, "./test/msp/keystore", false)

	csr, key, err := GetCSRPEM(name, enum.AppAlgorithmType_R1, fks)

	fmt.Println("key:", hex.EncodeToString(key.SKI()))

	fmt.Println("csrPEM：", csr)

}
