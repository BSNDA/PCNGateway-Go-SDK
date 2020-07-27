package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type CallContractReqData struct {
	base.BaseReqModel
	Body CallContractReqDataReqDataBody `json:"body"` // 消息体
}

type CallContractReqDataReqDataBody struct {
	UserId       string `json:"userId"` //用户ID必须由6~100位字母或数字组成！
	UserAddr     string `json:"userAddr"`
	ContractName string `json:"contractName"`
	FuncName     string `json:"funcName"`
	FuncParam    string `json:"funcParam"`
}

func (m *CallContractReqData) GetEncryptionValue() string {

	return m.GetBaseEncryptionValue() + m.Body.UserId + m.Body.UserAddr + m.Body.ContractName + m.Body.FuncName + m.Body.FuncParam
}
