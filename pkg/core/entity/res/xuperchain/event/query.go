package event

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
)

type QueryEventResData struct {
	base.BaseResModel
	Body QueryEventResDataBody `json:"body"` // 消息体
}

type QueryEventResDataBody struct {
	BlockEvent []QueryEventResData_BlockEvent `json:"blockEvent"`
}

type QueryEventResData_BlockEvent struct {
	EventId      string `json:"eventId"`
	EventKey     string `json:"eventKey"`
	NotifyUrl    string `json:"notifyUrl"`
	AttachArgs   string `json:"attachArgs"`
	CreateTime   string `json:"createTime"`
	UserCode     string `json:"userCode"`
	AppCode      string `json:"appCode"`
	ContractName string `json:"contractName"`
}

func (f *QueryEventResData) GetEncryptionValue() string {
	fp := ""
	for _, b := range f.Body.BlockEvent {
		fp += b.EventId
		fp += b.EventKey
		fp += b.AppCode
		fp += b.UserCode
		fp += b.ContractName
		fp += b.NotifyUrl
		fp += b.AttachArgs
		fp += b.CreateTime
	}

	return f.GetBaseEncryptionValue() + fp
}
