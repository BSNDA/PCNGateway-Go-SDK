package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric"
)

type TransReqData struct {
	base.BaseReqModel
	Body TransReqDataBody `json:"body"`
}

type TransReqDataBody struct {
	UserName string `json:"userName"`

	// Nonce
	Nonce string `json:"nonce"`

	ChainCode    string            `json:"chainCode"`
	FuncName     string            `json:"funcName"`
	Args         []string          `json:"args"`
	TransientMap map[string]string `json:"transientData"`
}

func (f *TransReqData) GetEncryptionValue() string {

	fp := f.GetBaseEncryptionValue()

	fp = fp + f.Body.UserName
	fp = fp + f.Body.Nonce
	fp = fp + f.Body.ChainCode
	fp = fp + f.Body.FuncName

	for _, a := range f.Body.Args {
		fp = fp + a
	}

	for k, v := range f.Body.TransientMap {
		fp = fp + k + v
	}

	return fp

}

func (t *TransReqDataBody) GetTransRequest(channelId string) *fabric.TransRequest {
	request := &fabric.TransRequest{
		ChannelId:   channelId,
		ChaincodeId: t.ChainCode,
		Fcn:         t.FuncName,
	}

	for _, a := range t.Args {
		request.Args = append(request.Args, []byte(a))
	}
	request.TransientMap = make(map[string][]byte)
	for k, v := range t.TransientMap {
		request.TransientMap[k] = []byte(v)
	}

	return request
}
