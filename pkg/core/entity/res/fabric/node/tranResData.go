package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type TranResData struct {
	base.BaseResModel
	Body *TranResDataBody `json:"body"`
}

type TranResDataBody struct {
	BlockInfo BlockInfo `json:"blockInfo"`

	CCRes CCRes `json:"ccRes"`
}

type BlockInfo struct {
	TxId string `json:"txId"`

	BlockHash string `json:"blockHash"`

	Status int `json:"status"`
}

type CCRes struct {
	CCCode int `json:"ccCode"`

	CCData string `json:"ccData"`
}

func (f *TranResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	fp := f.GetBaseEncryptionValue()
	fp = fp + f.Body.BlockInfo.TxId
	fp = fp + f.Body.BlockInfo.BlockHash
	fp = fp + strconv.Itoa(f.Body.BlockInfo.Status)

	fp = fp + strconv.Itoa(f.Body.CCRes.CCCode)
	fp = fp + f.Body.CCRes.CCData

	return fp
}
