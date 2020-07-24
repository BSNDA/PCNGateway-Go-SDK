package user

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RegisterUserReqData struct {
	base.BaseReqModel
	Body RegisterUserReqDataBody `json:"body"` // 消息体
}

type RegisterUserReqDataBody struct {
	UserId string `json:"userId"`
}

func (m *RegisterUserReqData) GetEncryptionValue() string {

	return m.GetBaseEncryptionValue() + m.Body.UserId
}
