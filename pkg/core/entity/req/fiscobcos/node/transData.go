package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type TransData struct {
	Contract ContractData
	FuncName string
	Args     []interface{}
	UserName string
}

type ContractData struct {
	ContractName    string
	ContractAbi     string
	ContractAddress string
}

type KeyTransReqData struct {
	base.BaseReqModel
	Body KeyTransReqDataBody `json:"body"` //消息体
}

type KeyTransReqDataBody struct {
	ContractName string `json:"contractName"` //合约名称
	TransData    string `json:"transData"`
}

func (f *KeyTransReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.ContractName + f.Body.TransData

}
