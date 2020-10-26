package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RemoveEventResData struct {
	base.BaseResModel
	Body interface{} `json:"body"`
}

func (f *RemoveEventResData) GetEncryptionValue() string {
	return f.GetBaseEncryptionValue()
}
