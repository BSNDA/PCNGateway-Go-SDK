package cita

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/node"
	"testing"
)

func getCitaClient(t *testing.T) *CitaClient {

	config, err := config2.NewMockCitaConfig()

	if err != nil {
		t.Fatal(err)
	}
	citaClient, err := NewCitaClient(config)

	if err != nil {
		t.Fatal(err)
	}
	return citaClient
}

func TestCitaClient_ReqChainCode_insert(t *testing.T) {

	citaClient := getCitaClient(t)

	name := "test852"
	var param = []string{}

	//s:= fmt.Sprintf("%064s",hex.EncodeToString([]byte("test10281536")))
	param = append(param, "0000000000000000000000000000000000000000000000000000000074657374")
	//param = append(param,hex.EncodeToString([]byte("abcdf")))
	param = append(param, "31313131")
	jb, _ := json.Marshal(param)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "CitaBsnBaseContract",
		FuncName:     "insert",
		FuncParam:    string(jb),
	}

	res, _ := citaClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))

}
func TestCitaClient_ReqChainCode_update(t *testing.T) {

	citaClient := getCitaClient(t)

	name := "zhmtest1"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("test10281456")))
	param = append(param, s)
	param = append(param, hex.EncodeToString([]byte("abcdwfm")))

	jb, _ := json.Marshal(param)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "CitaBsnBaseContract",
		FuncName:     "update",
		FuncParam:    string(jb),
	}

	res, _ := citaClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))

}
func TestCitaClient_ReqChainCode_get(t *testing.T) {

	citaClient := getCitaClient(t)

	name := "zhmtest1"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("test")))
	param = append(param, s)
	//param = append(param,hex.EncodeToString([]byte("abcdf")))

	jb, _ := json.Marshal(param)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "CitaBsnBaseContract",
		FuncName:     "retrieve",
		FuncParam:    string(jb),
	}

	res, _ := citaClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))

}
func TestCitaClient_ReqChainCode_remove(t *testing.T) {

	citaClient := getCitaClient(t)

	name := "zhmtest1"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("test")))
	param = append(param, s)
	//param = append(param,hex.EncodeToString([]byte("abcdf")))

	jb, _ := json.Marshal(param)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "CitaBsnBaseContract",
		FuncName:     "remove",
		FuncParam:    string(jb),
	}

	res, _ := citaClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))

}

func TestCitaClient_GetBlockInfo(t *testing.T) {

	citaClient := getCitaClient(t)

	data := nodereq.BlockReqDataBody{
		BlockNumber: "1",
		//BlockHash:"0x64dbb27dfc3af603bb24314d94c1d19c4ee2390686c26b87a49295c818742ea0",
	}

	res, err := citaClient.GetBlockInfo(data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
func TestCitaClient_GetBlockHeight(t *testing.T) {

	citaClient := getCitaClient(t)

	res, err := citaClient.GetBlockHeight()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}
func TestCitaClient_GetTxReceiptByTxHash(t *testing.T) {

	citaClient := getCitaClient(t)

	data := nodereq.TxTransReqDataBody{
		TxId: "0x2f26f1c5cccf0c4db5e06cc08626534711187b83ded089d58c6005c15d0e37c7",
		//TxId:"0xe558c63b889bb7b7cb1c3c0745e4278f5d6c47e97e7bfa035c5f9925514b3be0",
	}
	res, err := citaClient.GetTxReceiptByTxHash(data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}
func TestCitaClient_GetTxInfoByTxHash(t *testing.T) {

	citaClient := getCitaClient(t)

	data := nodereq.TxTransReqDataBody{

		TxId: "0x428ff18d3f753b03117820592929f23d7b425b977951ca64bb05488fe0a20080",
	}
	res, err := citaClient.GetTxInfoByTxHash(data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))
}
