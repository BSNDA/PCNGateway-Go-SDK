package cita

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/uuid"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"bytes"
	"strings"
)

var (
	QUOTA = uint64(10000000)
	Value = []byte{0x0}
)

func TransData(contractabi, contractAddress string, funcName string, args []interface{}, blockLimit uint64, chainId string, version uint32, smcrypto bool, privKey key.PrivateKeyProvider) (string, bool, error) {

	abi, err := abi.JSON(strings.NewReader(contractabi))
	if err != nil {
		return "", false, err
	}

	_, ok := abi.Methods[funcName]
	if !ok {
		return "", false, errors.New("contract not has function")
	}

	funcData, err := pack(&abi, funcName, args, false)
	if err != nil {
		return "", false, err
	}

	//if method.IsConstant() {
	//	return hexutil.Encode(funcData), true, nil
	//}
	toAddress := common.HexToAddress(contractAddress)

	nonce := uuid.GetUUID()

	tx := NewTransaction(nonce, toAddress, QUOTA, blockLimit, Value, chainId, funcData, version)

	dataByte, err := SignData(tx, privKey, smcrypto)

	if err != nil {
		return "", false, err
	}

	trans, err := serializeUnverifiedTransaction(tx, dataByte)
	if err != nil {
		return "", false, err
	}
	return trans, false, nil

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
	h := &key.SM3Hash{}
	hash := h.Hash(digest)
	return hash[:4]
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
