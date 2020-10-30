package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type BlockReqData struct {
	base.BaseReqModel
	Body BlockReqDataBody `json:"body"`
}
type BlockReqDataBody struct {
	BlockNumber string `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
}

func (f *BlockReqData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.BlockNumber + f.Body.BlockHash

}
