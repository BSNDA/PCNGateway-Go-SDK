package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockTxResData struct {
	base.BaseResModel
	Body TransactionData `json:"body"`
}

type TransactionData struct {
	TxHash          string `json:"txHash"`
	BlockHash       string `json:"blockHash"`
	BlockNumber     int64  `json:"blockNumber"`
	Data            string `json:"data"`
	ChainId         string `json:"chainId"`
	Quota           string `json:"quota"`
	From            string `json:"from"`
	To              string `json:"to"`
	Nonce           string `json:"nonce"`
	ValidUntilBlock string `json:"validUntilBlock"`
	Version         string `json:"version"`
}

func (f *TransactionData) GetEncryptionValue() string {

	fp := ""

	fp = fp + f.TxHash
	fp = fp + f.BlockHash
	fp = fp + strconv.FormatInt(f.BlockNumber, 10)
	fp = fp + f.Data
	fp = fp + f.ChainId
	fp = fp + f.Quota
	fp = fp + f.From
	fp = fp + f.To
	fp = fp + f.Nonce
	fp = fp + f.ValidUntilBlock
	fp = fp + f.Version

	return fp
}

func (f *BlockTxResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
