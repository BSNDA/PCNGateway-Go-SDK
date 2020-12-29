package cita

import (
	"fmt"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/event"
	"testing"
)

func TestCitaClient_EventRegister(t *testing.T) {

	citaClient := getCitaClient(t)

	body := eventreq.RegisterReqDataBody{
		EventType:       2,
		ContractAddress: "0xb4e5fa1f9f65908e04d322d6d0bb89702e88c986",
		//ContractName:    "GetTopic",
		NotifyUrl:  "http://192.168.1.172:58011/v1/fabric/test",
		AttachArgs: "abc=123",
	}

	res, err := citaClient.EventRegister(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}

func TestCitaClient_EventQuery(t *testing.T) {
	citaClient := getCitaClient(t)

	res, err := citaClient.EventQuery()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}

func TestCitaClient_EventRemove(t *testing.T) {
	citaClient := getCitaClient(t)

	body := eventreq.RemoveReqDataBody{
		EventId: "41d7d52e604e110365e4046072be88da970c96c533452ceeca77aa4e1cab54c3",
	}
	res, err := citaClient.EventRemove(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
