package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type QueryEventReqData struct {
	base.BaseReqModel
	Body interface{} `json:"body"`
}

func (f *QueryEventReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
