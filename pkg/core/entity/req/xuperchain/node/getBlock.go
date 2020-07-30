package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type GetBlockInfoReqData struct {
	base.BaseReqModel
	Body GetBlockInfoReqDataBody `json:"body"` // 消息体
}

type GetBlockInfoReqDataBody struct {
	BlockHeight int64  `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
}

func (m *GetBlockInfoReqData) GetEncryptionValue() string {

	return m.GetBaseEncryptionValue() + strconv.FormatInt(m.Body.BlockHeight, 10) + m.Body.BlockHash
}
