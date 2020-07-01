package base

import "encoding/json"

type ReqHeader struct {
	UserCode string `json:"userCode"` // User unique identification
	AppCode  string `json:"appCode"`  // App unique identification
	TId      string `json:"tId"`
}

type BaseReqModel struct {
	Header ReqHeader `json:"header"`

	Mac string `json:"mac"`
}

func (b *BaseReqModel) GetBaseEncryptionValue() string {

	return b.Header.UserCode + b.Header.AppCode
}

func (b *BaseReqModel) ToJson() []byte {

	jsonBytes, _ := json.Marshal(b)
	return jsonBytes
}
