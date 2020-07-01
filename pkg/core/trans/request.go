package trans

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type TransRequest struct {
	ChannelId    string
	ChaincodeId  string
	Fcn          string
	Args         [][]byte
	TransientMap map[string][]byte
}

func (t *TransRequest) GetRequest() *fab.ChaincodeInvokeRequest {

	request := &fab.ChaincodeInvokeRequest{
		ChaincodeID:  t.ChaincodeId,
		TransientMap: t.TransientMap,
		Fcn:          t.Fcn,
		Args:         t.Args,
	}

	return request

}
