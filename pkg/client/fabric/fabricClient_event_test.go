package fabric

import (
	"fmt"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/event"
	"testing"
)

func TestFabricClient_EventRegister(t *testing.T) {

	fabricClient := getFabricClient(t)

	body := req.RegisterReqDataBody{
		ChainCode:  "cc_app0001202007291443281737652_01",
		EventKey:   "test",
		NotifyUrl:  "http://192.168.6.85:18080/api/fisco/test",
		AttachArgs: "a=1",
	}

	res, _ := fabricClient.EventRegister(body)

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
		EventId: "eaf3f0ca28f0455db4fc9fa2b8b0c8d3",
	}

	res, _ := fabricClient.EventRemove(body)

	fmt.Println(res)
}
