package fisco_bcos

import (
	"fmt"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/event"
	"testing"
)

func TestFiscoBcosClient_EventRegister(t *testing.T) {

	fiscoClient := getK1Client(t)

	body := eventreq.RegisterReqDataBody{
		EventType:       1,
		ContractAddress: "0x866aefc204b8f8fdc3e45b908fd43d76667d7f76",
		ContractName:    "BsnBaseContractk1",
		NotifyUrl:       "http://192.168.6.85:18080/api/fiscobcos/test",
		AttachArgs:      "abc=123",
	}

	res, err := fiscoClient.EventRegister(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}

func TestFiscoBcosClient_EventQuery(t *testing.T) {
	fiscoClient := getK1Client(t)

	res, err := fiscoClient.EventQuery()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestFiscoBcosClient_EventRemove(t *testing.T) {
	fiscoClient := getK1Client(t)

	body := eventreq.RemoveReqDataBody{
		EventId: "764d6a2d8c0e44dd824bbc196830f282",
	}

	res, err := fiscoClient.EventRemove(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
