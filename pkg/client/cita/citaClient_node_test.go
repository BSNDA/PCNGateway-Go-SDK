package cita

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
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

	name := "testcurel"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("安慕希苹果味")))
	param = append(param, s)
	param = append(param, hex.EncodeToString([]byte("金典牛奶")))
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

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("金典牛奶")))
	param = append(param, s)
	param = append(param, hex.EncodeToString([]byte("安慕希")))

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

	name := "test0611"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("金典有机奶")))
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

	name := "test0611"
	var param = []string{}

	s := fmt.Sprintf("%064s", hex.EncodeToString([]byte("金典有机奶")))
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
func TestCitaClient_ReqChainCode_keyAtIndex(t *testing.T) {

	citaClient := getCitaClient(t)

	name := "test1123"
	var param []interface{}

	param = append(param, 1)

	jb, _ := json.Marshal(param)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "CitaBsnBaseContract",
		FuncName:     "keyAtIndex",
		FuncParam:    string(jb),
	}

	res, _ := citaClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))

}

func TestCitaClient_GetBlockInfo(t *testing.T) {

	citaClient := getCitaClient(t)

	data := nodereq.BlockReqDataBody{
		//BlockNumber: "1",
		BlockHash: "0xd9d76c64d6c2cb73adb431c23b1d0f2835ec0a9f39d9347f966f1237fdfb56f0",
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
		TxId: "0xf7376a31106a60778614b354e0628757e6ac588105195470cc4c3cb6e74cb3ab",
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

		TxId: "0x873a3ed633476808208299f3ebea985bf78fd44d36ab4a4ef42336f5451c8d32",
	}
	res, err := citaClient.GetTxInfoByTxHash(data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))
}

func TestCitaClient_Trans(t *testing.T) {
	citaClient := getCitaClient(t)
	if citaClient.Config.GetAppInfo().CAType == enum.AppCaType_Trust {
		fmt.Println("the trusteeship application cannot call the api")
		return
	}
	name := "test0611"
	var param []interface{}
	s, err := getbyte32(hex.EncodeToString([]byte("金典有机奶")))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	param = append(param, s)
	param = append(param, []byte("我不知道写啥好呢"))

	body := nodereq.TransData{
		UserName: name,
		Contract: nodereq.ContractData{
			ContractName:    "CitaBsnBaseContract",
			ContractAddress: "0x678711905fe98218406967812afa027ef0c5829e",
			ContractAbi:     `[{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"},{"name":"baseValue","type":"bytes"}],"name":"update","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"index","type":"uint256"}],"name":"keyAtIndex","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"}],"name":"remove","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"},{"name":"baseValue","type":"bytes"}],"name":"insert","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"baseKey","type":"bytes32"}],"name":"retrieve","outputs":[{"name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`,
		},
		//FuncName: "update",
		//FuncName: "retrieve",
		FuncName: "insert",
		//FuncName: "remove",
		Args: param,
	}

	res, err := citaClient.Trans(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))
}

func getbyte32(str string) ([32]byte, error) {
	var a [32]byte
	a1, _ := hex.DecodeString(str)
	//a1 := []byte(str)
	l := len(a1)
	if l > 32 {
		return a, errors.New("characters are  too long")
	}

	for i := 0; i < l; i++ {
		a[32-(l-i)] = a1[i]
	}
	return a, nil
}
