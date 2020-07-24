package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterEventResData struct {
	base.BaseResModel
	Body RegisterEvent `json:"body"`
}

type RegisterEvent struct {
	EventId string `json:"eventId"`
}

func (f *RegisterEventResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.EventId

}
