package user

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterResData struct {
	base.BaseResModel
	Body *RegisterResDataBody `json:"body"`
}

type RegisterResDataBody struct {
	Name   string `json:"name"` //Less than 20 bits
	Secret string `json:"secret"`
}

func (f *RegisterResData) GetEncryptionValue() string {

	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Name
	fp = fp + f.Body.Secret
	return fp
}
