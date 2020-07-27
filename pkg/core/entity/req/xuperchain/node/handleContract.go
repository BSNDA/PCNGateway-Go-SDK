package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type HandleContractReqData struct {
	base.BaseReqModel
	Body HandleContractReqDataReqDataBody `json:"body"` // 消息体
}

type HandleContractReqDataReqDataBody struct {
	UserId       string `json:"userId"` //用户ID必须由6~100位字母或数字组成！
	UserAddr     string `json:"userAddr"`
	ContractName string `json:"contractName"`
	FuncName     string `json:"funcName"`
	FuncParam    string `json:"funcParam"`
}

func (m *HandleContractReqData) GetEncryptionValue() string {

	return m.GetBaseEncryptionValue() + m.Body.UserId + m.Body.UserAddr + m.Body.ContractName + m.Body.FuncName + m.Body.FuncParam
}
