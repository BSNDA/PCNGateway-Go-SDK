package base

import "strconv"

const Header_error_code int = -1

const Header_success_code int = 0

type ResHeader struct {
	Code int `json:"code"`

	Msg string `json:"msg"`
}

type BaseResModel struct {
	Header *ResHeader `json:"header"`

	Mac string `json:"mac"` // mac
}

func (b *BaseResModel) GetMac() string {
	return b.Mac
}

func (b *BaseResModel) GetBaseEncryptionValue() string {

	return strconv.Itoa(b.Header.Code) + b.Header.Msg
}
