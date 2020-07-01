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

func TestFabricClient_SdkTran(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	name := ""

	var args []string
	args = append(args, "{\"baseKey\":\"test20200406\",\"baseValue\":\"this is string \"}")

	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_base",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := fabricClient.SdkTran(body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

}

func TestFabricClient_ReqChainCode(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("cert", config.GetCert())

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	name := "user01"

	var args []string
	args = append(args, "{\"baseKey\":\"test2020066\",\"baseValue\":\"this is string \"}")

	nonce, _ := crypto.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_base",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, _ := fabricClient.ReqChainCode(body)

	fmt.Println(res)

}

func TestFabricClient_GetTransInfo(t *testing.T) {
	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	tx := req.TxTransReqDataBody{
		TxId: "255cf25cc7ed1ef94a2c23e7d52747426f193470a2f716ba36c93d17323e6272",
	}

	res, _ := fabricClient.GetTransInfo(tx)

	fmt.Println(res)
	if res.Header.Code == 0 {
		tm := time.Unix(res.Body.TimeSpanSec, res.Body.TimeSpanNsec)

		fmt.Println(tm.Format("2006-01-02 15:04:05.000 -0700 MST"))
	}

}

func TestFabricClient_GetBlockInfo(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	tx := req.BlockReqDataBody{
		BlockHash: "d04fdb12073abfee4f3ee45472468a1cf0434e74d6342671a8c43b713f6a5e92",
	}

	res, _ := fabricClient.GetBlockInfo(tx)

	v := fabricClient.Verify(res.Mac, res.GetEncryptionValue())
	fmt.Println(v)

	fmt.Println(res)

}

func TestFabricClient_GetLedgerInfo(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	res, _ := fabricClient.GetLedgerInfo()

	fmt.Println(res)
}
