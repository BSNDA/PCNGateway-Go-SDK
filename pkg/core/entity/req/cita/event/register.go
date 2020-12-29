package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type RegisterReqData struct {
	base.BaseReqModel
	Body RegisterReqDataBody `json:"body"`
}

type RegisterReqDataBody struct {
	EventType int `json:"eventType"` //1 block 2 contract

	//if EventType is 1，ContractAddress and ContractName Can be empty
	//if EventType is 2，ContractAddress and ContractName Cannot be empty at the same time， ContractAddress first
	ContractAddress string `json:"contractAddress"`
	ContractName    string `json:"contractName"`

	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
}

func (f *RegisterReqDataBody) GetEncryptionValue() string {
	fp := strconv.Itoa(f.EventType)
	fp = fp + f.ContractAddress
	fp = fp + f.ContractName
	fp = fp + f.NotifyUrl
	fp = fp + f.AttachArgs
	return fp
}

func (f *RegisterReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
