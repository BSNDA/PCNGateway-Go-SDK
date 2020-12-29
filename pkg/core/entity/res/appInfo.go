package res

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type AppInfoResData struct {
	base.BaseResModel
	Body AppInfoResDataBody `json:"body"`
}

type AppInfoResDataBody struct {
	AppName       string `json:"appName"`
	AppType       string `json:"appType"`
	CaType        int    `json:"caType"`
	AlgorithmType int    `json:"algorithmType"`
	MspId         string `json:"mspId"`
	ChannelId     string `json:"channelId"`
	Version       string `json:"version"`
}

func (f *AppInfoResData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.AppName
	fp = fp + f.Body.AppType
	fp = fp + strconv.Itoa(f.Body.CaType)
	fp = fp + strconv.Itoa(f.Body.AlgorithmType)
	fp = fp + f.Body.MspId
	fp = fp + f.Body.ChannelId
	return fp
}
