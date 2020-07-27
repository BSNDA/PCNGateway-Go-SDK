package user

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterResData struct {
	base.BaseResModel
	Body RegisterResDataBody `json:"body"`
}

type RegisterResDataBody struct {
	UserId      string `json:"userId"`
	UserAddress string `json:"userAddress"`
}

func (f *RegisterResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.UserId + f.Body.UserAddress
}
