package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type SdkTransReqData struct {
	base.BaseReqModel
	Body SdkTransReqDataBody `json:"body"`
}

type SdkTransReqDataBody struct {
	TransData string `json:"transData"`
}

func (f *SdkTransReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.TransData

}
