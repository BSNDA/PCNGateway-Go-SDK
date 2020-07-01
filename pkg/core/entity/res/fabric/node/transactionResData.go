package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type TransactionResData struct {
	base.BaseResModel
	Body *TransactionResDataBody `json:"body"`
}

type TransactionResDataBody struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber uint64 `json:"blockNumber"`

	Status       int    `json:"status"`
	CreateName   string `json:"createName"`
	TimeSpanSec  int64  `json:"timeSpanSec"`
	TimeSpanNsec int64  `json:"timeSpanNsec"`
}

func (f *TransactionResData) GetEncryptionValue() string {

	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockHash
	fp = fp + strconv.FormatUint(f.Body.BlockNumber, 10)
	fp = fp + strconv.Itoa(f.Body.Status)
	fp = fp + f.Body.CreateName
	fp = fp + strconv.FormatInt(f.Body.TimeSpanSec, 10)
	fp = fp + strconv.FormatInt(f.Body.TimeSpanNsec, 10)

	return fp
}
