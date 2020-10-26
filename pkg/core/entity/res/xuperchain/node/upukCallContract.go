package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type UPukCallContractResData struct {
	base.BaseResModel
	Body *UPukCallContractResDataBody `json:"body"`
}

type UPukCallContractResDataBody struct {
	TxId       string `json:"txId"`
	QueryInfo  string `json:"queryInfo"`
	PreExecRes string `json:"preExecRes"`
}

func (f *UPukCallContractResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()
	fp += f.Body.TxId + f.Body.QueryInfo + f.Body.PreExecRes
	return fp
}
