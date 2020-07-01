package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockResData struct {
	base.BaseResModel
	Body *BlockResDataBody `json:"body"`
}

type BlockResDataBody struct {
	BlockHash    string `json:"blockHash"`
	BlockNumber  uint64 `json:"blockNumber"`
	PreBlockHash string `json:"preBlockHash"`
	BlockSize    uint64 `json:"blockSize"`
	BlockTxCount int    `json:"blockTxCount"`

	Transactions []*TransactionData `json:"transactions"`
}

type TransactionData struct {
	TxId         string `json:"txId"`
	Status       int    `json:"status"`
	CreateName   string `json:"createName"`
	TimeSpanSec  int64  `json:"timeSpanSec"`
	TimeSpanNsec int64  `json:"timeSpanNsec"`
}

func (f *BlockResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockHash
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + f.Body.PreBlockHash
	fp = fp + strconv.FormatUint(f.Body.BlockSize, 10)
	fp = fp + strconv.Itoa(f.Body.BlockTxCount)

	for _, t := range f.Body.Transactions {

		fp = fp + t.TxId
		fp = fp + strconv.Itoa(t.Status)
		fp = fp + t.CreateName
		fp = fp + strconv.FormatInt(t.TimeSpanSec, 10)
		fp = fp + strconv.FormatInt(t.TimeSpanNsec, 10)

	}

	return fp
}
