package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type UPukCallContractReqData struct {
	base.BaseReqModel
	Body UPukCallContractReqDataReqDataBody `json:"body"` // 消息体
}

type UPukCallContractReqDataReqDataBody struct {
	Initiator string `json:"initiator"`
	TransData string `json:"transData"`
	Flag      int    `json:"flag"` // 0：预执行	1：执行
}

func (m *UPukCallContractReqData) GetEncryptionValue() string {
	return m.GetBaseEncryptionValue() + m.Body.Initiator + m.Body.TransData + strconv.Itoa(m.Body.Flag)
}
