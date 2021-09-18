package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockDataRes struct {
	base.BaseResModel
	Body *BlockDataResBody `json:"body"`
}

type BlockDataResBody struct {
	BlockHash    string `json:"blockHash"`
	BlockNumber  uint64 `json:"blockNumber"`
	PreBlockHash string `json:"preBlockHash"`
	BlockData    string `json:"blockData"`
}

func (f *BlockDataRes) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockHash
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.PreBlockHash
	fp = fp + f.Body.BlockData

	return fp
}
