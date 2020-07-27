// @Title  xuperchain_node_test
// @Description
// @Author  zxl  2020/7/22 19:32
// @Version 1.0.0
// @Update  2020/7/22 19:32
package xuperchain

import (
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/node"
	"testing"
)

func TestXuperChainClient_ReqChainCode_Insert_Data(t *testing.T) {
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201707",
		UserAddr:     "2BMwkb9R3khePcCPj1QzVRQa8pfNurAnAQ",
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201707",
		UserAddr:     "2BMwkb9R3khePcCPj1QzVRQa8pfNurAnAQ",
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201707",
		UserAddr:     "2BMwkb9R3khePcCPj1QzVRQa8pfNurAnAQ",
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.CallContractReqDataReqDataBody{
		UserId:       "zxl072201707",
		UserAddr:     "2BMwkb9R3khePcCPj1QzVRQa8pfNurAnAQ",
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.GetTxInfoReqDataBody{
		TxHash: "1a869cfe2dc09086fb2389ed6934f66baf70467cf4634cbfe756e4b2e478c384",
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.GetBlockInfoReqDataBody{
		BlockHeight: 275006,
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
	config, err := config2.NewMockXuperchainSMConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	client, err := NewXuperChainClient(config)
	if err != nil {
		t.Fatal(err.Error())
	}
	body := req.GetBlockInfoReqDataBody{
		BlockHeight: 0, //must be zero
		BlockHash:   "5b8237023cfe025c6fec9ad5a05f295f6097ceba9980199a5989361013a557ed",
	}
	res, err := client.GetBlockInfo(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
