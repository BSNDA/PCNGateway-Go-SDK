package fiscobcos

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tjfoc/gmsm/sm3"

	"bytes"
	"math/big"
	"strings"
)

var (
	GAS_PRICE = new(big.Int).SetInt64(22000000000)
	GAS_LIMIT = new(big.Int).SetInt64(4300000)

	ChainId = new(big.Int).SetInt64(1)
)

func ParesData(contractabi, funcName string, args []interface{}, smcrypto bool) (string, error) {
	abi, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return "", err
	}

	funcData, err := pack(&abi, funcName, args, smcrypto)
	if err != nil {
		return "", err
	}

	return hexutil.Encode(funcData), nil

}

func TransData(contractabi, contractAddress string, funcName string, args []interface{}, groupId, blockLimit *big.Int, extraData []byte, smcrypto bool, privKey interface{}) (string, bool, error) {

	abi, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return "", false, err
	}

	method, ok := abi.Methods[funcName]
	if !ok {
		return "", false, errors.New("contract not has function")
	}

	funcData, err := pack(&abi, funcName, args, smcrypto)
	if err != nil {
		return "", false, err
	}

	if method.IsConstant() {
		return hexutil.Encode(funcData), true, nil
	}
	toAddress := common.HexToAddress(contractAddress)

	nonce, _ := crypto.GetRandomBigInt()

	if err != nil {
		return "", false, err
	}

	tx := NewTransaction(nonce, toAddress, nil, GAS_LIMIT, GAS_PRICE, blockLimit, funcData, ChainId, groupId, extraData, smcrypto)

	dataByte, err := tx.SignData(privKey)

	if err != nil {
		return "", false, err
	}

	return hexutil.Encode(dataByte), false, nil

}

func pack(abi *abi.ABI, funcName string, args []interface{}, sm bool) ([]byte, error) {
	if sm {

		if funcName == "" {
			// constructor
			arguments, err := abi.Constructor.Inputs.Pack(args...)
			if err != nil {
				return nil, err
			}
			return arguments, nil
		}

		method, ok := abi.Methods[funcName]

		if !ok {
			return nil, errors.New("contract not has function")
		}

		id := getMethodId(method)

		arguments, err := method.Inputs.Pack(args...)
		if err != nil {
			return nil, err
		}

		return append(id, arguments...), nil

	} else {
		return abi.Pack(funcName, args...)
	}

}

func getMethodId(method abi.Method) []byte {
	digest := []byte(method.Sig)

	h := sm3.New()
	h.Write(digest)
	hash := h.Sum(nil)
	return hash[:4]
}

func ParseTrans(data string, contractAbi string, sm bool) (contractAddress, funcName string, constant bool, err error) {

	err = nil

	dataByte, err := hexutil.Decode(data)
	if err != nil {
		return
	}

	sign := &txsign{}

	err = rlp.DecodeBytes(dataByte, sign)
	if err != nil {
		funcName, constant, err = ParseMethod(dataByte, contractAbi, sm)
		return
	}

	contractAddress = strings.ToLower(sign.Recipient.String())
	funcName, constant, err = ParseMethod(sign.Payload, contractAbi, sm)

	return

}

func ParseMethod(data []byte, contractabi string, sm bool) (funcName string, constant bool, err error) {
	abiObj, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return
	}

	if sm {
		sig := data[:4]
		for _, method := range abiObj.Methods {
			id := getMethodId(method)
			if bytes.Equal(id, sig) {

				funcName = method.Name
				constant = method.IsConstant()
				return

			}
		}

	} else {
		var method *abi.Method
		method, err = abiObj.MethodById(data)
		if err != nil {
			return
		}

		funcName = method.Name
		constant = method.IsConstant()
		return

	}

	err = errors.New("func is not found")
	return

}
