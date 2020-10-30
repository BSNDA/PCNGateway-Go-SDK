package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type TransResData struct {
	base.BaseResModel
	Body *TransResDataDataBody `json:"body"`
}

type TransResDataDataBody struct {
	TxId string `json:"txId"`

	Status string `json:"status"`

	Data string `json:"data"`
}

func (f *TransResData) getBlockValue() string {

	if f.Body == nil {
		return ""
	}

	fb := ""

	fb = fb + f.Body.TxId

	fb = fb + f.Body.Status
	fb = fb + f.Body.Data

	return fb
}

func (f *TransResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.getBlockValue()

}
