// @Title  xuperchain_node_test
// @Description
// @Author  zxl  2020/7/22 19:32
// @Version 1.0.0
// @Update  2020/7/22 19:32
package xuperchain

import (
	"encoding/base64"
	"fmt"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/node"
	trans "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain/pb"
	"github.com/golang/protobuf/proto"
	"testing"
)

func getXuperChainClient(t *testing.T) *XuperChainClient {
	config, err := config2.NewMockXuperchainConfig()
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

// 上传公钥模式invoke合约
// initiator:代表xupchain的账户地址；可调用方法GetAddressFromPublicKey生成address信息，文件地址（trans/xuperchain/account/account_ext.go）
func TestXuperChainClient_SdkTran_Increase(t *testing.T) {
	client := getXuperChainClient(t)
	// step 1 准备测试数据
	var arg = struct {
		contractName string
		methodName   string
		initiator    string
		args         map[string][]byte
	}{
		contractName: "cc_appxc_01",
		methodName:   "increase",
		initiator:    "rJHdYSyCWcnhZ9bjhfbhoN5mfhMVCWGhX",
		args: map[string][]byte{
			"key": []byte("dev_002"),
		},
	}
	// step 2 执行pre invoke
	invokeRequest := trans.GenerateInvokeIR(arg.contractName, arg.methodName, arg.args)
	bytes, _ := proto.Marshal(invokeRequest)
	body := req.UPukCallContractReqDataReqDataBody{
		Initiator: arg.initiator,
		Flag:      0,
		TransData: base64.StdEncoding.EncodeToString(bytes),
	}
	res, err := client.SdkTran(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
	t.Log(fmt.Sprintf("preExecRes:%s", res.Body.PreExecRes))

	// step 3 执行invoke
	privateKey := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgikcrsA9vRp14VAl0
lhNmOMc7pl0j4xlF9Eno+eJNgSegCgYIKoEcz1UBgi2hRANCAAQuR25rig9+Isir
q7eapxSsDCo8FCdt0qfCa7eGCC7BEdGerbX4qa5j3qljl1/d6I00c7veFAvgGPP8
DFT+md7r
-----END PRIVATE KEY-----`
	publicKey := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAELkdua4oPfiLIq6u3mqcUrAwqPBQn
bdKnwmu3hgguwRHRnq21+KmuY96pY5df3eiNNHO73hQL4Bjz/AxU/pne6w==
-----END PUBLIC KEY-----`

	var preExecResponse = new(pb.InvokeResponse)

	buf, err := base64.StdEncoding.DecodeString(res.Body.PreExecRes)
	if err != nil {
		t.Fatal(err)
	}
	err = proto.Unmarshal(buf, preExecResponse)
	if err != nil {
		t.Fatal(err)
	}
	transaction, err := trans.GenerateTransaction(arg.initiator, preExecResponse, privateKey, publicKey)
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ = proto.Marshal(transaction)
	body = req.UPukCallContractReqDataReqDataBody{
		Initiator: arg.initiator,
		Flag:      1,
		TransData: base64.StdEncoding.EncodeToString(bytes),
	}
	res, err = client.SdkTran(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
	t.Log(fmt.Sprintf("txid:%s", res.Body.TxId))
}

// 上传公钥模式query合约
// initiator:代表xupchain的账户地址；可调用方法GetAddressFromPublicKey生成address信息，文件地址（trans/xuperchain/account/account_ext.go）
func TestXuperChainClient_SdkTran_Get(t *testing.T) {
	client := getXuperChainClient(t)
	// step 1 准备测试数据
	var arg = struct {
		contractName string
		methodName   string
		initiator    string
		args         map[string][]byte
	}{
		contractName: "cc_appxc_01",
		methodName:   "get",
		initiator:    "rJHdYSyCWcnhZ9bjhfbhoN5mfhMVCWGhX",
		args: map[string][]byte{
			"key": []byte("dev_002"),
		},
	}
	// step 2 执行pre
	invokeRequest := trans.GenerateInvokeIR(arg.contractName, arg.methodName, arg.args)
	bytes, _ := proto.Marshal(invokeRequest)
	body := req.UPukCallContractReqDataReqDataBody{
		Initiator: arg.initiator,
		Flag:      0,
		TransData: base64.StdEncoding.EncodeToString(bytes),
	}
	res, err := client.SdkTran(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
	t.Log(fmt.Sprintf("QueryInfo:%s", res.Body.QueryInfo))
}

func TestXuperChainClient_ReqChainCode_Insert_Data(t *testing.T) {

	client := getXuperChainClient(t)

	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201708",
		UserAddr:     "2CzpFVY3KYQcZZXBfEa6hmrFn17o4FMBdf",
		ContractName: "cc_appxc_01",
		FuncName:     "insert_data",
		FuncParam:    "{\"base_key\":\"dev_0001\",\"base_value\":\"aaron1\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
func TestXuperChainClient_ReqChainCode_Update_Data(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201708",
		UserAddr:     "2CzpFVY3KYQcZZXBfEa6hmrFn17o4FMBdf",
		ContractName: "cc_appxc_01",
		FuncName:     "update_data",
		FuncParam:    "{\"base_key\":\"dev_0001\",\"base_value\":\"aaron.zhang\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
func TestXuperChainClient_ReqChainCode_Select_Data(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201708",
		UserAddr:     "2CzpFVY3KYQcZZXBfEa6hmrFn17o4FMBdf",
		ContractName: "cc_appxc_01",
		FuncName:     "select_data",
		FuncParam:    "{\"base_key\":\"dev_0001\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestXuperChainClient_ReqChainCode_Remove_Data(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201708",
		UserAddr:     "2CzpFVY3KYQcZZXBfEa6hmrFn17o4FMBdf",
		ContractName: "cc_appxc_01",
		FuncName:     "remove_data",
		FuncParam:    "{\"base_key\":\"dev_0001\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestXuperChainClient_GetTxInfo(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.GetTxInfoReqDataBody{
		TxHash: "8ac7856fb8521ed90a40c9667ff1b617ab2bbaa076e051f58067915f1d9c3e6e",
	}
	res, err := client.GetTxInfo(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestXuperChainClient_GetBlockInfo_ByBlockHeight(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.GetBlockInfoReqDataBody{
		BlockHeight: 276229,
		BlockHash:   "", //must be empty
	}
	res, err := client.GetBlockInfo(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
func TestXuperChainClient_GetBlockInfo_ByBlockHash(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.GetBlockInfoReqDataBody{
		BlockHeight: 0, //must be zero
		BlockHash:   "bae9328a98409aa5d0c9d0a061d0ef59dfc7f22add23aa50fb9cea7fcf6a5ea8",
	}
	res, err := client.GetBlockInfo(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestXuperChainClient_ReqChainCode_Increase(t *testing.T) {

	client := getXuperChainClient(t)

	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl20201020",
		UserAddr:     "222mRGrQ1sqEu8B8x57EfihTxirJi5a3v7",
		ContractName: "cc_appxc_01",
		FuncName:     "increase",
		FuncParam:    "{\"key\":\"zxlcounter\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
func TestXuperChainClient_ReqChainCode_get(t *testing.T) {

	client := getXuperChainClient(t)

	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl20201020",
		UserAddr:     "222mRGrQ1sqEu8B8x57EfihTxirJi5a3v7",
		ContractName: "cc_appxc_01",
		FuncName:     "get",
		FuncParam:    "{\"key\":\"zxlcounter\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestXuperChainClient_ReqChainCode_Get(t *testing.T) {

	client := getXuperChainClient(t)

	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl20201020",
		UserAddr:     "2AVKHf8EXRVTmoCAEdjXbQ153hjfmcW9cw",
		ContractName: "cc_appxc_01",
		FuncName:     "get",
		FuncParam:    "{\"key\":\"zxlcounter\"}",
	}
	res, err := client.ReqChainCode(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
