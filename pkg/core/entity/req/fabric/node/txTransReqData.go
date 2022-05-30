package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type TxTransReqData struct {
	base.BaseReqModel
	Body TxTransReqDataBody `json:"body"`
}

type TxTransReqDataBody struct {
	TxId string `json:"txId"`

	// DataType Options as json
	DataType string `json:"dataType,omitempty"`
}

func (f *TxTransReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.TxId + f.Body.DataType
}
