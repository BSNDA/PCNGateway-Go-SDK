package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterEventResData struct {
	base.BaseResModel
	Body RegisterEventResDataBody `json:"body"`
}

type RegisterEventResDataBody struct {
	EventId string `json:"eventId"`
}

func (f *RegisterEventResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue() + f.Body.EventId
}
