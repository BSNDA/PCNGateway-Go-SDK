package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RegisterEventReqData struct {
	base.BaseReqModel
	Body RegisterEventReqDataBody `json:"body"`
}

type RegisterEventReqDataBody struct {
	ContractName string `json:"contractName"`
	EventKey     string `json:"eventKey"`
	NotifyUrl    string `json:"notifyUrl"`
	AttachArgs   string `json:"attachArgs"`
}

func (f *RegisterEventReqData) GetEncryptionValue() string {
	fp := f.Body.ContractName
	fp = fp + f.Body.EventKey
	fp = fp + f.Body.NotifyUrl
	fp = fp + f.Body.AttachArgs
	return f.GetBaseEncryptionValue() + fp
}
