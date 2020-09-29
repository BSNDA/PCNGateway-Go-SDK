package fabric

import (
	"encoding/base64"
	"fmt"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto"
	"testing"
	"time"
)

func getFabricClient(t *testing.T) *FabricClient {

	//if call test net
	//Uncomment code
	//config2.SetTest()

	config, err := config2.NewMockFabricConfig()

	if err != nil {
		t.Fatal(err)
	}
	//fmt.Println("cert", config.GetCert())
	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err)
	}
	return fabricClient
}

func TestFabricClient_SdkTran(t *testing.T) {

	fabricClient := getFabricClient(t)

	name := "testuser3"

	var args []string
	args = append(args, "{\"baseKey\":\"test20200456\",\"baseValue\":\"this is string \"}")

	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_99613acedfe94e879252f25a50f5bb27",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := fabricClient.SdkTran(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	v := fabricClient.Verify(res.Mac, res.GetEncryptionValue())
	fmt.Println("网关签名验证：", v)

}

func TestFabricClient_ReqChainCode(t *testing.T) {

	fabricClient := getFabricClient(t)

	name := ""

	var args []string
	//args = append(args, "test2020104")
	args = append(args, "{\"baseKey\":\"test2020105\",\"baseValue\":\"this is string \"}")

	nonce, _ := crypto.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_99613acedfe94e879252f25a50f5bb27",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, _ := fabricClient.ReqChainCode(body)

	fmt.Println(res)

	v := fabricClient.Verify(res.Mac, res.GetEncryptionValue())
	fmt.Println(v)

}

func TestFabricClient_GetTransInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.TxTransReqDataBody{
		TxId: "9092ccfe3af254f5afc4b0aba770f521c1357ea8e47e97757a0eca7da047cb94",
	}

	res, _ := fabricClient.GetTransInfo(tx)

	fmt.Println(res)
	if res.Header.Code == 0 {
		tm := time.Unix(res.Body.TimeSpanSec, res.Body.TimeSpanNsec)

		fmt.Println(tm.Format("2006-01-02 15:04:05.000 -0700 MST"))
	}

}

func TestFabricClient_GetBlockInfo(t *testing.T) {

	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{

		BlockHash: "8b29e587b15a4d6edf1b66f8a26e9f773db64f3aa292897b2577e1808bffbeb1",
	}

	res, _ := fabricClient.GetBlockInfo(tx)

	v := fabricClient.Verify(res.Mac, res.GetEncryptionValue())
	fmt.Println(v)

	fmt.Println(res)

}

func TestFabricClient_GetLedgerInfo(t *testing.T) {

	fabricClient := getFabricClient(t)

	res, _ := fabricClient.GetLedgerInfo()

	fmt.Println(res)
	fmt.Println(fabricClient.Verify(res.Mac, res.GetEncryptionValue()))
}
