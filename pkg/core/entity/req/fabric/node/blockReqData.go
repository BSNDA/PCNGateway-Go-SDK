package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockReqData struct {
	base.BaseReqModel
	Body BlockReqDataBody `json:"body"`
}

type BlockReqDataBody struct {
	BlockNumber uint64 `json:"blockNumber"`
	BlockHash   string `json:"blockHash"`
	TxId        string `json:"txId"`
	// DataType Options as json
	DataType string `json:"dataType,omitempty"`
}

func (f *BlockReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.BlockHash
	fp = fp + f.Body.TxId
	fp = fp + f.Body.DataType
	return fp

}
