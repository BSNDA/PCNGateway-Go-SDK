package user

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RegisterUserResData struct {
	base.BaseResModel
	Body *RegisterUserResDataBody `json:"body"` // 消息体
}

type RegisterUserResDataBody struct {
	AccAddr string
}

func (f *RegisterUserResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.AccAddr
	return fp
}
