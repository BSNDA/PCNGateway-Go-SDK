package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type TxTransReqData struct {
	base.BaseReqModel
	Body TxTransReqDataBody `json:"body"`
}

type TxTransReqDataBody struct {
	TxId string `json:"txHash"`
}

func (f *TxTransReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.TxId
}
