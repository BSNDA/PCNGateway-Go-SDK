package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type GetTxInfoResData struct {
	base.BaseResModel
	Body *Transaction `json:"body"` // 消息体
}

func (f *GetTxInfoResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}
	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.Txid + f.Body.Blockid + strconv.Itoa(int(f.Body.Version))

	for _, v := range f.Body.ContractRequests {
		fp += v.ContractName
		fp += v.MethodName
		fp += v.Args
	}
	fp += strconv.FormatInt(f.Body.ReceivedTimestamp, 10)
	return fp
}
