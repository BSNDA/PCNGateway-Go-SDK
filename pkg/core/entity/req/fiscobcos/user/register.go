package user

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"`
}

type RegisterReqDataBody struct {
	UserId string `json:"userId"`
}

func (f *RegisterReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.UserId

}
