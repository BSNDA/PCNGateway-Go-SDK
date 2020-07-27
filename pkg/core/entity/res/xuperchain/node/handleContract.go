package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type HandleContractResData struct {
	base.BaseResModel
	Body *HandleContractResDataBody `json:"body"` // 消息体
}

type HandleContractResDataBody struct {
	TxId      string `json:"txId"`
	QueryInfo string `json:"queryInfo"`
}

func (f *HandleContractResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.TxId + f.Body.QueryInfo
	return fp
}
