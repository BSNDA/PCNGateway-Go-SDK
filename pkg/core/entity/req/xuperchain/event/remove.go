package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RemoveEventReqData struct {
	base.BaseReqModel
	Body RemoveEventReqDataBody `json:"body"`
}

type RemoveEventReqDataBody struct {
	EventId string `json:"eventId"`
}

func (f *RemoveEventReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.EventId
}
