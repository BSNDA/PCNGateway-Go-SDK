package user

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"`
}

type RegisterReqDataBody struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func (f *RegisterReqData) GetEncryptionValue() string {
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Name
	fp = fp + f.Body.Secret

	return fp

}
