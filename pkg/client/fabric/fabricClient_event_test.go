package fabric

import (
	"fmt"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/event"
	"testing"
)

func TestFabricClient_EventRegister(t *testing.T) {

	fabricClient := getFabricClient(t)

	body := req.RegisterReqDataBody{
		ChainCode:  "cc_app0001202012111600499234472_02",
		EventKey:   "test",
		NotifyUrl:  "http://192.168.1.172:58011/v1/fabric/test",
		AttachArgs: "a=1",
	}

	res, _ := fabricClient.EventRegister(body)

	fmt.Println(res)

}

func TestFabricClient_BlockEventRegister(t *testing.T) {

	fabricClient := getFabricClient(t)

	body := req.RegisterReqDataBody{
		NotifyUrl:  "http://192.168.1.172:58011/v1/fabric/test",
		AttachArgs: "a=1",
	}

	res, _ := fabricClient.BlockEventRegister(body)

	fmt.Println(res)

}

func TestFabricClient_EventQuery(t *testing.T) {
	fabricClient := getFabricClient(t)

	res, _ := fabricClient.EventQuery()

	fmt.Println(res)
}

func TestFabricClient_EventRemove(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.RemoveReqDataBody{
		EventId: "c5d10b8df8f64700a8bccaf46d390dd8",
	}

	res, _ := fabricClient.EventRemove(body)

	fmt.Println(res)
}
