/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  trans_test
 * @Version: 1.0.0
 * @Date: 2020/6/2 10:53
 */

package fiscobcos

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/BSNDA/bsn-sdk-crypto/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/rlp"
	"math/big"

	"strings"
	"testing"
)

func Test1(t *testing.T) {

	abiString := "[{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"aaaa\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"a\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"b\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"FourTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Stored\",\"type\":\"event\"}]"

	funcName := "set"
	//groupId :=130

	contractAddress := "0xa977d5f4b19b966d489d641c280a16462750b4f7"

	var list []interface{}

	pr := new(big.Int).SetInt64(101)
	//list = append(list, "101")
	list = append(list, pr)
	//list = append(list, "abc")

	abi, err := abi.JSON(strings.NewReader(abiString))

	if err != nil {
		t.Fatal(err)
	}

	p, err := abi.Pack(funcName, list...)

	if err != nil {
		t.Fatal(err)
	}

	hex1 := hexutil.Encode(p)

	fmt.Println(hex1)

	address := common.HexToAddress(contractAddress)

	nonce, _ := new(big.Int).SetString("1682842859963368860006956442320140456128348364022072681970215451535956266928", 10)
	GAS_PRICE := new(big.Int).SetInt64(22000000000)
	GAS_LIMIT := new(big.Int).SetInt64(4300000)
	blockLimit := new(big.Int).SetInt64(4)

	chainId := new(big.Int).SetInt64(1)
	groupId := new(big.Int).SetInt64(118)
	fmt.Println("funcData", hexutil.Encode(p))
	tx := NewTransaction(nonce, address, nil, GAS_LIMIT, GAS_PRICE, blockLimit, p, chainId, groupId, nil, false)
	txb, err := rlp.EncodeToBytes(tx.data)

	s1 := hexutil.Encode(txb)
	ss := "0xf9012ba003b8748ede4016399647b2a6f8e619556cc4d3f560e59be3b3fbf6d97b844fb085051f4d5c0083419ce00494a977d5f4b19b966d489d641c280a16462750b4f780b8e4ebf3b24f0000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000001610000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000036162630000000000000000000000000000000000000000000000000000000000017680"
	//tx2 :=&Transaction{}

	//bb2 ,_:=hexutil.Decode(ss)
	//rlp.DecodeBytes(bb2,tx2)

	//fmt.Println(strings.ToLower(tx2.data.Recipient.Hex()))

	//mm,err :=abi.MethodById( tx2.data.Payload)

	//fmt.Println(mm.Name)
	fmt.Println(s1)
	fmt.Println(ss)
	fmt.Println(s1 == ss)

	prik := `-----BEGIN PRIVATE KEY-----
MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgbsKArzr1MVA77N3PNzR9
ss9CLHlmZ/9NhcjVhTE2aMShRANCAASaYIp2L1Rl/+GSsFv3tdPVAunBYHjSfoTV
bQVByl2ZoXPc2dNsowENzM9d8+aBoXfUVVx8Si3bm81F9b3aBdBa
-----END PRIVATE KEY-----`

	prikey, err := key.NewPrivateKeyProvider(types.ECDSA_K1, prik) //  eth.LoadPrivateKey([]byte(prik))
	if err != nil {
		t.Fatal(err)
	}

	txd, _ := tx.SignData(prikey)
	s2 := hexutil.Encode(txd)

	fmt.Println(s2)

	tx2 := &Transaction{}

	bb2, _ := hexutil.Decode(s2)
	rlp.DecodeBytes(bb2, tx2)

	fmt.Println(strings.ToLower(tx2.data.Recipient.Hex()))

	mm, err := abi.MethodById(tx2.data.Payload)

	fmt.Println(mm.Name)
}

func Test2(t *testing.T) {

	var list []interface{}

	list = append(list, 1)
	list = append(list, "abc")

	by, _ := json.Marshal(list)

	fmt.Println(string(by))

	s := "123456"

	h := sha256.New()

	h.Write([]byte(s))
	hash := h.Sum(nil)

	fmt.Println(hex.EncodeToString(hash))

}

func Test3(t *testing.T) {
	abiString := "[{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"string\"},{\"name\":\"trade_hash\",\"type\":\"string\"}],\"name\":\"isExisted\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userId\",\"type\":\"string\"}],\"name\":\"select\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userId\",\"type\":\"string\"},{\"name\":\"trade_hash\",\"type\":\"string\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ret\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"userId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"userHash\",\"type\":\"string\"}],\"name\":\"eventForInsert\",\"type\":\"event\"}]"
	abi, _ := abi.JSON(strings.NewReader(abiString))

	m := abi.Methods["insert"]

	fmt.Println(m.Sig)

}
func Test4(t *testing.T) {
	abiString := `[{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"},{"name":"baseValue","type":"bytes"}],"name":"update","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"index","type":"uint256"}],"name":"keyAtIndex","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"}],"name":"remove","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"baseKey","type":"bytes32"},{"name":"baseValue","type":"bytes"}],"name":"insert","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"baseKey","type":"bytes32"}],"name":"retrieve","outputs":[{"name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		t.Fatal(err)
	}
	funcName := "insert"

	var list []interface{}

	//s:= fmt.Sprintf("%064s",hex.EncodeToString([]byte("test1102")))
	list = append(list, getbyte32("test1102"))
	list = append(list, []byte("test1102"))

	p, err := abi.Pack(funcName, list...)

	if err != nil {
		t.Fatal(err)
	}

	hex1 := hexutil.Encode(p)

	fmt.Println(hex1)
}

func getbyte32(str string) [32]byte {
	var a [32]byte

	a1 := []byte(str)
	l := len(a1)
	if l > 32 {

	}

	for i := 0; i < l; i++ {
		a[32-(l-i)] = a1[i]
	}
	return a
}

func TestByte(t *testing.T) {
	var a [32]byte

	a1 := []byte("test1102")
	l := len(a1)
	if l > 32 {

	}

	for i := 0; i < l; i++ {
		a[32-(l-i)] = a1[i]
	}

	fmt.Println(a)
	fmt.Println(a1)
}

func Test5(t *testing.T) {
	abiString := `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		t.Fatal(err)
	}
	funcName := "select"

	var args []interface{}
	args = append(args, "string")

	p, err := abi.Pack(funcName, args...)

	if err != nil {
		t.Fatal(err)
	}

	hex1 := hexutil.Encode(p)

	fmt.Println(hex1)
}

func TestParesData(t *testing.T) {
	abiString := `[{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"update","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"}],"name":"remove","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"base_id","type":"string"},{"name":"base_key","type":"int256"},{"name":"base_value","type":"string"}],"name":"insert","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"base_id","type":"string"}],"name":"select","outputs":[{"name":"","type":"string[]"},{"name":"","type":"int256[]"},{"name":"","type":"string[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`
	funcName := "insert"
	var args []interface{}
	pr := new(big.Int).SetInt64(101)
	args = append(args, "abc123")
	args = append(args, pr)
	args = append(args, "abc12345")

	data, err := ParesData(abiString, funcName, args, true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)
}
