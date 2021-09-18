package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type TranDataRes struct {
	base.BaseResModel
	Body *TranDataResBody `json:"body"`
}

type TranDataResBody struct {
	TxId      string `json:"blockInfo"`
	TransData string `json:"transData"`
}

func (f *TranDataRes) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.TxId
	fp = fp + f.Body.TransData
	return fp
}
