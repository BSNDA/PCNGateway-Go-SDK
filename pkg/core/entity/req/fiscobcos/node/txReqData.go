package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type TxReqData struct {
	base.BaseReqModel
	Body TxReqDataBody `json:"body"`
}
type TxReqDataBody struct {
	TxHash string `json:"txHash"`
}

func (f *TxReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.TxHash

}
