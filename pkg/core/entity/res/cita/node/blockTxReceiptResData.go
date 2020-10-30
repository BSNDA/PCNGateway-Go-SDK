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
	TransactionHash        string               `json:"transactionHash"`
	TransactionIndex       int64                `json:"transactionIndex"`
	BlockHash              string               `json:"blockHash"`
	BlockNumber            uint64               `json:"blockNumber"`
	CumulativeGasUsed      string               `json:"cumulativeGasUsed"`
	CumulativeQuotaUsed    uint64               `json:"cumulativeQuotaUsed"`
	GasUsed                string               `json:"gasUsed"`
	QuotaUsed              uint64               `json:"quotaUsed"`
	ContractAddress        string               `json:"contractAddress"`
	Root                   string               `json:"root"`
	Status                 string               `json:"status"`
	From                   string               `json:"from"`
	To                     string               `json:"to"`
	LogsBloom              string               `json:"logsBloom"`
	ErrorMessage           string               `json:"errorMessage"`
	TransactionIndexRaw    string               `json:"transactionIndexRaw"`
	BlockNumberRaw         string               `json:"blockNumberRaw"`
	CumulativeGasUsedRaw   string               `json:"cumulativeGasUsedRaw"`
	CumulativeQuotaUsedRaw string               `json:"cumulativeQuotaUsedRaw"`
	GasUsedRaw             string               `json:"gasUsedRaw"`
	QuotaUsedRaw           string               `json:"quotaUsedRaw"`
	Logs                   []TranReceiptLogData `json:"logs"`
}
type TranReceiptLogData struct {
	Removed             bool     `json:"removed"`
	LogIndex            uint64   `json:"logIndex"`
	TransactionIndex    uint64   `json:"transactionIndex"`
	TransactionHash     string   `json:"transactionHash"`
	BlockHash           string   `json:"blockHash"`
	BlockNumber         uint64   `json:"blockNumber"`
	Address             string   `json:"address"`
	Data                string   `json:"data"`
	TransactionLogIndex string   `json:"transactionLogIndex"`
	TransactionIndexRaw string   `json:"transactionIndexRaw"`
	BlockNumberRaw      string   `json:"blockNumberRaw"`
	LogIndexRaw         string   `json:"logIndexRaw"`
	Topics              []string `json:"topics"`
}

func (f *TxReceiptData) GetEncryptionValue() string {

	if f == nil {
		return ""
	}

	fb := ""

	fb = fb + f.TransactionHash
	fb = fb + strconv.FormatInt(f.TransactionIndex, 10)
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatUint(f.BlockNumber, 10)
	fb = fb + f.CumulativeGasUsed
	fb = fb + strconv.FormatUint(f.CumulativeQuotaUsed, 10)
	fb = fb + f.GasUsed
	fb = fb + strconv.FormatUint(f.QuotaUsed, 10)
	fb = fb + f.ContractAddress
	fb = fb + f.Root
	fb = fb + f.Status
	fb = fb + f.From
	fb = fb + f.To
	fb = fb + f.LogsBloom
	fb = fb + f.ErrorMessage
	fb = fb + f.TransactionIndexRaw
	fb = fb + f.BlockNumberRaw
	fb = fb + f.CumulativeGasUsedRaw
	fb = fb + f.CumulativeQuotaUsedRaw
	fb = fb + f.GasUsedRaw
	fb = fb + f.QuotaUsedRaw
	for _, t := range f.Logs {
		fb = fb + strconv.FormatBool(t.Removed)
		fb = fb + strconv.FormatUint(t.LogIndex, 10)
		fb = fb + strconv.FormatUint(t.TransactionIndex, 10)
		fb = fb + t.TransactionHash
		fb = fb + t.BlockHash
		fb = fb + strconv.FormatUint(t.BlockNumber, 10)
		fb = fb + t.Address
		fb = fb + t.Data
		fb = fb + t.TransactionLogIndex
		fb = fb + t.TransactionIndexRaw
		fb = fb + t.BlockNumberRaw
		fb = fb + t.LogIndexRaw
		for _, s := range t.Topics {
			fb = fb + s
		}
	}

	return fb
}

func (f *BlockTxReceiptResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.GetEncryptionValue()

}
