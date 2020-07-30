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

	config, err := config2.NewFabricSMTrustConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	name := "user05"

	var args []string
	args = append(args, "{\"baseKey\":\"test20200407\",\"baseValue\":\"this is string \"}")

	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_app0001202007271538152051987_01",
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
	fmt.Println(v)

}

func TestFabricClient_ReqChainCode(t *testing.T) {

	config, err := config2.NewFabricSMNoTrustConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println("cert", config.GetCert())

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	name := ""

	var args []string
	args = append(args, "{\"baseKey\":\"test2020069\",\"baseValue\":\"this is string \"}")

	nonce, _ := crypto.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_app0001202007291443281737652_01",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, _ := fabricClient.ReqChainCode(body)

	fmt.Println(res)

}

func TestFabricClient_GetTransInfo(t *testing.T) {
	config, err := config2.NewFabricSMNoTrustConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	tx := req.TxTransReqDataBody{
		TxId: "d74808fa0c355e85aeb6e6f5289dca4295f30343dd1bed35a67eef776624d7d4",
	}

	res, _ := fabricClient.GetTransInfo(tx)

	fmt.Println(res)
	if res.Header.Code == 0 {
		tm := time.Unix(res.Body.TimeSpanSec, res.Body.TimeSpanNsec)

		fmt.Println(tm.Format("2006-01-02 15:04:05.000 -0700 MST"))
	}

}

func TestFabricClient_GetBlockInfo(t *testing.T) {

	config, err := config2.NewFabricSMTrustConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	tx := req.BlockReqDataBody{

		BlockHash: "b340b4d2cd2f8e9e905f5ea14a6495aa0a1834fb68d08cd37283e9dcd9304379",
	}

	res, _ := fabricClient.GetBlockInfo(tx)

	v := fabricClient.Verify(res.Mac, res.GetEncryptionValue())
	fmt.Println(v)

	fmt.Println(res)

}

func TestFabricClient_GetLedgerInfo(t *testing.T) {

	config, err := config2.NewFabricSMTrustConfig()

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
