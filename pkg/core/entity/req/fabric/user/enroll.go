package user

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type EnrollReqData struct {
	base.BaseReqModel
	Body EnrollReqDataBody `json:"body"`
}

type EnrollReqDataBody struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
	CsrPem string `json:"csrPem"`
}

func (f *EnrollReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Name
	fp = fp + f.Body.Secret
	fp = fp + f.Body.CsrPem

	return fp

}
