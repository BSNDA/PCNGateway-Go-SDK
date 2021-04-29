package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type QueryResData struct {
	base.BaseResModel
	Body []EventQueryBody `json:"body"`
}

type EventQueryBody struct {
	EventKey   string `json:"eventKey,omitempty"`
	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
	EventId    string `json:"eventId"`
	CreateTime string `json:"createTime"`
	OrgCode    string `json:"orgCode"`
	UserCode   string `json:"userCode"`
	AppCode    string `json:"appCode"`
	ChainCode  string `json:"chainCode,omitempty"`
	EventType  string `json:"eventType,omitempty"`
}

func (f *QueryResData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()
	for _, task := range f.Body {
		fp = fp + task.EventKey
		fp = fp + task.NotifyUrl
		fp = fp + task.AttachArgs
		fp = fp + task.EventId
		fp = fp + task.CreateTime
		fp = fp + task.OrgCode
		fp = fp + task.UserCode
		fp = fp + task.AppCode
		fp = fp + task.ChainCode
		fp = fp + task.EventType
	}

	return fp
}
