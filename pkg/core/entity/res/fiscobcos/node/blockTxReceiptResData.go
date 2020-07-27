package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockTxReceiptResData struct {
	base.BaseResModel
	Body *TxReceiptData `json:"body"`
}

type TxReceiptData struct {
	TxId        string `json:"txId"`
	BlockHash   string `json:"blockHash"`
	BlockNumber int64  `json:"blockNumber"`

	GasUsed int64 `json:"gasUsed"`

	From string `json:"from"`
	To   string `json:"to"`

	ContractAddress string `json:"contractAddress"`
	Logs            string `json:"logs"`
}

func (f *TxReceiptData) GetEncryptionValue() string {

	fb := ""

	fb = fb + f.TxId
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatInt(f.BlockNumber, 10)
	fb = fb + strconv.FormatInt(f.GasUsed, 10)
	fb = fb + f.From
	fb = fb + f.To

	fb = fb + f.ContractAddress
	fb = fb + f.Logs
	return fb
}

func (f *BlockTxReceiptResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.GetEncryptionValue()

}
