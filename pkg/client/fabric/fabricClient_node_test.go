package fabric

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	"github.com/BSNDA/bsn-sdk-crypto/common"
)

func getFabricClient(t *testing.T) *FabricClient {

	//if call test net
	//Uncomment code
	//config2.SetTest()

	config, err := config2.NewMockFabricConfig()

	if err != nil {
		t.Fatal(err)
	}
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
	args = append(args, "{\"baseKey\":\"test20201222\",\"baseValue\":\"this is string11 \"}")

	nonce, _ := common.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_app0001202012111600499234472_02",
		FuncName:     "update",
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
		TxId: "b448496c1f161bd62709915ed3cdb79d595fc585dbd1237e60fe2d3911410e06",
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

		BlockHash: "2a491d3bed3e97b304797390182000a6fa625e11d74567142e1a8fc670c1bbf9",
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
