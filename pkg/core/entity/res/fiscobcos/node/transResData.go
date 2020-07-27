package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type TransResData struct {
	base.BaseResModel
	Body *TransResDataDataBody `json:"body"`
}

type TransResDataDataBody struct {
	Constant bool `json:"constant"`

	QueryInfo string `json:"queryInfo"`

	TxId        string `json:"txId"`
	BlockHash   string `json:"blockHash"`
	BlockNumber int64  `json:"blockNumber"`

	GasUsed int64 `json:"gasUsed"`

	Status string `json:"status"`

	From string `json:"from"`
	To   string `json:"to"`

	Input     string `json:"input"`
	Output    string `json:"output"`
	EventLogs string `json:"logs"`
}

func (f *TransResData) getBlockValue() string {

	if f.Body == nil {
		return ""
	}

	fb := ""
	if f.Body.Constant {
		fb = fb + "true"
	} else {
		fb = fb + "false"
	}
	fb = fb + f.Body.QueryInfo
	fb = fb + f.Body.TxId
	fb = fb + f.Body.BlockHash
	fb = fb + strconv.FormatInt(f.Body.BlockNumber, 10)
	fb = fb + strconv.FormatInt(f.Body.GasUsed, 10)
	fb = fb + f.Body.Status
	fb = fb + f.Body.From
	fb = fb + f.Body.To
	fb = fb + f.Body.Input
	fb = fb + f.Body.Output
	fb = fb + f.Body.EventLogs

	return fb
}

func (f *TransResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.getBlockValue()

}
