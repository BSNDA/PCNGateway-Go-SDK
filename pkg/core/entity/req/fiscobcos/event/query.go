package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type QueryReqData struct {
	base.BaseReqModel
	Body interface{} `json:"body"`
}

func (f *QueryReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
