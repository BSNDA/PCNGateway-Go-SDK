package fabric

import (
	"encoding/base64"
	"fmt"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	"github.com/BSNDA/bsn-sdk-crypto/common"
	"testing"
	"time"
)

func getFabricClient(t *testing.T) *FabricClient {

	//if Call test net
	//Uncomment code
	//config2.SetTest()

	api := "http://192.168.1.61:17512"
	userCode := "USER0001202108171423128315082"
	appCode := "app0001202204251015286226029"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQglETKLei6cn/6KwEU
IPINz6tOMWE/eJGNZVtTu65xtjOgCgYIKoEcz1UBgi2hRANCAATnpj8cIMAWCVZM
QpCKmYv2FBF+6Q0mQvPGfySWQn95CSt9xSL/0FIF7H7qQtn5D5PWhatLPurAo1z1
ye9gfY6d
-----END PRIVATE KEY-----`
	config2.SetTest()

	config, err := config2.NewConfig(api, userCode, appCode, privK, "E:\\test")

	if err != nil {
		t.Fatal(err)
	}

	fabricClient, err := InitFabricClient(config, WithDefaultNodeName("a"))

	if err != nil {
		t.Fatal(err)
	}
	return fabricClient
}

func TestFabricClient_SdkTran(t *testing.T) {

	fabricClient := getFabricClient(t)

	name := "user20220421"

	user, _ := fabricClient.LoadUser(name)

	var args []string
	args = append(args, "{\"baseKey\":\"test202110333\",\"baseValue\":\"this is string \"}")
	//args = append(args,"key1","value")
	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_app0001202204201700217624980_01",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := fabricClient.SdkTran(body, user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

}

func TestFabricClient_SdkTranQ(t *testing.T) {

	fabricClient := getFabricClient(t)

	name := "user20220421"
	var args []string
	args = append(args, "test202110333")
	//args = append(args,"key1","value")
	body := req.TransReqDataBody{
		UserName:     name,
		ChainCode:    "cc_app0001202204201700217624980_01",
		FuncName:     "get",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, err := fabricClient.SdkTran(body, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

}

func TestFabricClient_ReqChainCode(t *testing.T) {

	fabricClient := getFabricClient(t)

	name := ""

	var args []string
	//args = append(args, "test2020104")
	args = append(args, "{\"baseKey\":\"test20210926\",\"baseValue\":\"this is string11 \"}")
	//args = append(args, "key1", "value")
	nonce, _ := common.GetRandomNonce()

	body := req.TransReqDataBody{
		UserName:     name,
		Nonce:        base64.StdEncoding.EncodeToString(nonce),
		ChainCode:    "cc_app0001202111121647396153631_01",
		FuncName:     "set",
		Args:         args,
		TransientMap: make(map[string]string),
	}

	res, _ := fabricClient.ReqChainCode(body)

	fmt.Println(res)

}

func TestFabricClient_GetTransInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.TxTransReqDataBody{
		TxId: "623edd1185c1cdc32b54904eba107fee2f10938b6c446ca829642a2e10101aa9",
	}

	res, _ := fabricClient.GetTransInfo(tx)

	fmt.Println(res)
	if res.Header.Code == 0 {
		tm := time.Unix(res.Body.TimeSpanSec, res.Body.TimeSpanNsec)

		fmt.Println(tm.Format("2006-01-02 15:04:05.000 -0700 MST"))
	}

}
func TestFabricClient_GetTransDataInfo(t *testing.T) {
	fabricClient := getFabricClient(t)

	tx := req.TxTransReqDataBody{
		TxId: "55fcfe1e48fb96ffd2cf63501e7fd3e5460a456239c8b3956a8ef71305886822",
	}

	res, trans, _ := fabricClient.GetTransData(tx)

	fmt.Println(trans)
	fmt.Println(res)

}

func TestFabricClient_GetBlockInfo(t *testing.T) {

	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{
		BlockNumber: 1,
		//BlockHash: "f66e68e0ca0e45f092ec55aeb1a503afda35ab234d6da527bfb7d8abbc004d2a",
		//TxId: "d3715eac4e4af04e2662da21461e04887ebbde1aafd70d283e1900ced3b1a0fd",
	}

	res, _ := fabricClient.GetBlockInfo(tx)

	fmt.Println(res)

}

func TestFabricClient_GetLedgerInfo(t *testing.T) {

	fabricClient := getFabricClient(t)

	res, _ := fabricClient.GetLedgerInfo()

	fmt.Println(res.Mac)
	fmt.Println(res.GetEncryptionValue())

}

func TestFabricClient_GetBlockDataInfo(t *testing.T) {

	fabricClient := getFabricClient(t)

	tx := req.BlockReqDataBody{
		//BlockHash: "f66e68e0ca0e45f092ec55aeb1a503afda35ab234d6da527bfb7d8abbc004d2a",
		BlockNumber: 1,
		//TxId: "d3715eac4e4af04e2662da21461e04887ebbde1aafd70d283e1900ced3b1a0fd",
	}

	_, block, _ := fabricClient.GetBlockData(tx)

	fmt.Println(block)

}
