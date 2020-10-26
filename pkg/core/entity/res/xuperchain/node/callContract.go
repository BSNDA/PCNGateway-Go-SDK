package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type CallContractResData struct {
	base.BaseResModel
	Body *CallContractResDataBody `json:"body"`
}

type CallContractResDataBody struct {
	TxId      string `json:"txId"`
	QueryInfo string `json:"queryInfo"`
}

func (f *CallContractResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.TxId + f.Body.QueryInfo
	return fp
}
