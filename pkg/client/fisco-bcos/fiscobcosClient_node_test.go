package fisco_bcos

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/node"
	"math/big"
	"testing"
)

func getClient(t *testing.T) *FiscoBcosClient {
	config, err := config.NewMockFiscoConfig()

	if err != nil {
		t.Fatal(err)
	}

	fabricClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err)
	}
	return fabricClient
}

func getK1Client(t *testing.T) *FiscoBcosClient {
	config, err := config.NewMockFiscoConfig()

	if err != nil {
		t.Fatal(err)
	}

	fabricClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err)
	}
	return fabricClient
}

func getSMClient(t *testing.T) *FiscoBcosClient {
	config, err := config.NewMockFiscoConfig()

	if err != nil {
		t.Fatal(err)
	}

	fabricClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err)
	}
	return fabricClient
}

func TestFiscoBcosClient_ReqChainCode_insert(t *testing.T) {

	fiscoClient := getSMClient(t)

	name := "test0927"

	var args []interface{}
	args = append(args, "s0604")
	args = append(args, 12)
	args = append(args, "aa")

	jb, _ := json.Marshal(args)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "BsnBaseContract",
		FuncName:     "insert",
		FuncParam:    string(jb),
	}

	res, _ := fiscoClient.ReqChainCode(body)

	fmt.Println(res)

	fmt.Println(fiscoClient.Verify(res.Mac, res.GetEncryptionValue()))

}

func TestFiscoBcosClient_ReqChainCode_insertk1(t *testing.T) {

	fiscoClient := getK1Client(t)

	name := "test0623"

	var args []interface{}
	args = append(args, "s0605")
	//args = append(args, 12)
	//args = append(args, "aa")

	jb, _ := json.Marshal(args)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "BsnBaseGlobalContract",
		FuncName:     "select",
		FuncParam:    string(jb),
	}

	res, _ := fiscoClient.ReqChainCode(body)

	fmt.Println(res)

}

func TestFiscoBcosClient_ReqChainCode_query(t *testing.T) {

	fiscoClient := getClient(t)

	name := "test0623"

	var args []interface{}
	args = append(args, "s0604")
	//args = append(args, 1)
	//args = append(args, "b")

	jb, _ := json.Marshal(args)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "BsnBaseContract",
		FuncName:     "select",
		FuncParam:    string(jb),
	}

	res, _ := fiscoClient.ReqChainCode(body)

	fmt.Println(res)

}

func TestFiscoBcosClient_ReqChainCode_queryk1(t *testing.T) {

	fiscoClient := getK1Client(t)

	name := "test0611"

	var args []interface{}
	args = append(args, "s0604")
	//args = append(args, 1)
	//args = append(args, "b")

	jb, _ := json.Marshal(args)

	body := nodereq.TransReqDataBody{
		UserId:       name,
		ContractName: "BsnBaseContractk1",
		FuncName:     "select",
		FuncParam:    string(jb),
	}

	res, _ := fiscoClient.ReqChainCode(body)

	v := fiscoClient.Verify(res.Mac, res.GetEncryptionValue())

	fmt.Println(v)
	fmt.Println(res)

}

func TestFiscoBcosClient_GetBlockInfo(t *testing.T) {
	fiscoClient := getK1Client(t)
	data := nodereq.BlockReqDataBody{
		BlockNumber: "5",
	}

	res, err := fiscoClient.GetBlockInfo(data)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}

func TestFiscoBcosClient_GetBlockHeight(t *testing.T) {
	fiscoClient := getSMClient(t)

	res, err := fiscoClient.GetBlockHeight()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res.Body.Data)
	fmt.Println(fiscoClient.Verify(res.Mac, res.GetEncryptionValue()))
}

func TestFiscoBcosClient_GetTxCount(t *testing.T) {
	fiscoClient := getClient(t)

	res, err := fiscoClient.GetTxCount()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res.Body.Data)
}

func TestFiscoBcosClient_GetTxCountByBlockNumber(t *testing.T) {

	fiscoClient := getClient(t)

	res, err := fiscoClient.GetTxCountByBlockNumber(644)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res.Body.Data)
}

func TestFiscoBcosClient_GetTxReceiptByTxHash(t *testing.T) {
	fiscoClient := getClient(t)
	tx := nodereq.TxReqDataBody{
		TxHash: "0x76e8eda4229f1c8982089677f17a0d4c9959c8cdbb9478fa1fa064f65f2a493f",
	}

	res, _ := fiscoClient.GetTxReceiptByTxHash(tx)

	fmt.Println(res)
}

func TestFiscoBcosClient_GetTxInfoByTxHash(t *testing.T) {

	fiscoClient := getSMClient(t)

	tx := nodereq.TxReqDataBody{
		TxHash: "0x0ccf17d7110753e7cef3fd3fcea317332c7d873f0689c5bbb679a3078fb16881",
	}

	res, _ := fiscoClient.GetTxInfoByTxHash(tx)

	fmt.Println(res)
	fmt.Println(fiscoClient.Verify(res.Mac, res.GetEncryptionValue()))
}

func TestFiscoBcosClient_Trans(t *testing.T) {
	fiscoClient := getK1Client(t)

	name := "test0611"

	var args []interface{}
	args = append(args, "s0604")
	pr := new(big.Int).SetInt64(5)
	args = append(args, pr)
	args = append(args, "[0xf,0x0]")

	body := nodereq.TransData{
		UserName: name,
		Contract: nodereq.ContractData{
			ContractName:    "BsnBaseGlobalContract",
			ContractAddress: "0x5cc4d9d4945d8dd2eb371fb39b9d09b17eea70ff",
			ContractAbi:     `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`,
		},
		FuncName: "insert",
		Args:     args,
	}

	res, err := fiscoClient.Trans(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(fiscoClient.Verify(res.Mac, res.GetEncryptionValue()))
}

func TestFiscoBcosClient_Trans_SM(t *testing.T) {
	fiscoClient := getSMClient(t)

	name := "test0611"

	var args []interface{}
	args = append(args, "a1")
	pr := new(big.Int).SetInt64(5)
	args = append(args, pr)
	args = append(args, "aa")

	body := nodereq.TransData{
		UserName: name,
		Contract: nodereq.ContractData{
			ContractName:    "BsnBaseContract",
			ContractAddress: "0x9176d1e89033dcb6f1ed426935d267a04b0f6ccb", //0x9176d1e89033dcb6f1ed426935d267a04b0f6ccb
			ContractAbi:     `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`,
		},
		FuncName: "insert",
		Args:     args,
	}

	res, err := fiscoClient.Trans(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClient_Trans_Query(t *testing.T) {
	fiscoClient := getK1Client(t)

	name := "test0611"

	var args []interface{}
	args = append(args, "s0604")

	body := nodereq.TransData{
		UserName: name,
		Contract: nodereq.ContractData{
			ContractName:    "BsnBaseGlobalContract",
			ContractAddress: "0x5cc4d9d4945d8dd2eb371fb39b9d09b17eea70ff",
			ContractAbi:     `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`,
		},
		FuncName: "select",
		Args:     args,
	}

	res, err := fiscoClient.Trans(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res.Body.QueryInfo)

	var data [][]interface{}
	err = json.Unmarshal([]byte(res.Body.QueryInfo), &data)
	if err != nil {
		t.Fatal(err)
	}
	value := data[2]
	fmt.Println(value[len(value)-1].(string))
}

func TestFiscoBcosClient_Trans_Query_SM(t *testing.T) {
	fiscoClient := getSMClient(t)

	name := "test0611"

	var args []interface{}
	args = append(args, "s0604")

	body := nodereq.TransData{
		UserName: name,
		Contract: nodereq.ContractData{
			ContractName:    "BsnBaseContract",
			ContractAddress: "0xc206db9e77e547b015e2cb39d23ff8b0314746a4",
			ContractAbi:     `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`,
		},
		FuncName: "select",
		Args:     args,
	}

	res, err := fiscoClient.Trans(body)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res.Body.QueryInfo)

	var data [][]interface{}
	err = json.Unmarshal([]byte(res.Body.QueryInfo), &data)
	if err != nil {
		t.Fatal(err)
	}
	value := data[2]
	fmt.Println(value[len(value)-1].(string))
}
