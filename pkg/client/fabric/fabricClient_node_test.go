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

	name := "user0731"

	var args []string
	args = append(args, "{\"baseKey\":\"test20200423\",\"baseValue\":\"this is string \"}")

	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_app0001202007302150221237627_01",
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

	config, err := config2.NewFabricSMTrustConfig()

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
	args = append(args, "{\"baseKey\":\"test2020100\",\"baseValue\":\"this is string \"}")

	nonce, _ := crypto.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_app0001202007310152164084640_01",
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
	config, err := config2.NewFabricSMTrustConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	tx := req.TxTransReqDataBody{
		TxId: "391b5fd1e83e272c4a5593616e3f9b191de806a47ddd6d6049fdc6d79e1e977c",
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

		BlockHash: "8b29e587b15a4d6edf1b66f8a26e9f773db64f3aa292897b2577e1808bffbeb1",
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
