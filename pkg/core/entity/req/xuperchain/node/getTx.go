package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type GetTxInfoReqData struct {
	base.BaseReqModel
	Body GetTxInfoReqDataBody `json:"body"` // 消息体
}

type GetTxInfoReqDataBody struct {
	TxHash string `json:"txHash"`
}

func (m *GetTxInfoReqData) GetEncryptionValue() string {

	return m.GetBaseEncryptionValue() + m.Body.TxHash
}
