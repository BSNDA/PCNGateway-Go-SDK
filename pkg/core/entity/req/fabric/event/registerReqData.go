package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"`
}

type RegisterReqDataBody struct {
	ChainCode  string `json:"chainCode"`
	EventKey   string `json:"eventKey"`
	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
}

func (f *RegisterReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.ChainCode
	fp = fp + f.Body.EventKey
	fp = fp + f.Body.NotifyUrl
	fp = fp + f.Body.AttachArgs

	return fp

}
