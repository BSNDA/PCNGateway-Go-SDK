/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  event
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:15
 */

package event

type EventList struct {
	BlockEvent    []BlockEvent    `json:"blockEvent"`
	ContractEvent []ContractEvent `json:"contractEvent"`
}

type Event struct {
	EventId    string `json:"eventId"`
	AppCode    string `json:"appCode"`
	UserCode   string `json:"userCode"`
	NotifyUrl  string `json:"notifyUrl"`
	AttachArgs string `json:"attachArgs"`
	CreateTime string `json:"createTime"`
}

type BlockEvent struct {
	Event
}

type ContractEvent struct {
	Event
	ContractAddress string `json:"contractAddress"`
}

func (e *EventList) GetEncryptionValue() string {

	fp := ""
	for _, b := range e.BlockEvent {
		fp = fp + b.EventId
		fp = fp + b.AppCode
		fp = fp + b.UserCode
		fp = fp + b.NotifyUrl
		fp = fp + b.AttachArgs
		fp = fp + b.CreateTime
	}

	for _, b := range e.ContractEvent {
		fp = fp + b.EventId
		fp = fp + b.AppCode
		fp = fp + b.UserCode
		fp = fp + b.NotifyUrl
		fp = fp + b.AttachArgs
		fp = fp + b.CreateTime
		fp = fp + b.ContractAddress
	}

	return fp
}
