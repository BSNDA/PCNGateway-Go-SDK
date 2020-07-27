package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type GetBlockInfoResData struct {
	base.BaseResModel
	Body *Block `json:"body"` // 消息体
}

func (f *GetBlockInfoResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()
	fp = fp + strconv.Itoa(int(f.Body.Version)) + f.Body.Blockid + f.Body.PreHash + strconv.FormatInt(f.Body.Height, 10) + strconv.FormatInt(f.Body.Timestamp, 10)
	for _, t := range f.Body.Transactions {
		fp = fp + t.Txid + t.Blockid + strconv.Itoa(int(t.Version))
		for _, v := range t.ContractRequests {
			fp += v.ContractName
			fp += v.MethodName
			fp += v.Args
		}
		fp += strconv.FormatInt(t.ReceivedTimestamp, 10)
	}
	fp += strconv.Itoa(int(f.Body.TxCount)) + f.Body.NextHash
	return fp
}
