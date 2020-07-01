package req

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type AppInfoReqData struct {
	base.BaseReqModel
	Body AppInfoReqDataBody `json:"body"`
}

type AppInfoReqDataBody struct {
}

func (f *AppInfoReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
