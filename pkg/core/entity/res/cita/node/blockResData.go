package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockResData struct {
	base.BaseResModel
	Body *BlockData `json:"body"`
}

type BlockData struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      int64  `json:"blockNumber"`
	PrevBlockHash    string `json:"prevBlockHash"`
	BlockTime        string `json:"blockTime"`
	QuotaUsed        string `json:"quotaUsed"`
	TransactionsRoot string `json:"transactionsRoot"`
	StateRoot        string `json:"stateRoot"`
	ReceiptsRoot     string `json:"receiptsRoot"`

	Transactions []*TransactionData `json:"transactions"`
}

func (f *BlockData) GetEncryptionValue() string {
	fp := ""

	fp = fp + f.BlockHash
	fp = fp + strconv.FormatInt(f.BlockNumber, 10)
	fp = fp + f.PrevBlockHash
	fp = fp + f.BlockTime
	fp = fp + f.QuotaUsed
	fp = fp + f.TransactionsRoot
	fp = fp + f.StateRoot
	fp = fp + f.ReceiptsRoot

	for _, tx := range f.Transactions {
		fp = fp + tx.GetEncryptionValue()
	}

	return fp
}

func (f *BlockResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
