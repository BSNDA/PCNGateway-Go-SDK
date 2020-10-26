package xuperchain

import (
	"fmt"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/event"
	"testing"
)

func TestXuperChainClient_EventRegister(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.RegisterEventReqDataBody{
		ContractName: "cc_appxc_01",
		EventKey:     "increase_event",
		NotifyUrl:    "http://192.168.7.141:8088/revNotify1",
		AttachArgs:   "abc=123",
	}

	res, err := client.RegisterEvent(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
func TestXuperChainClient_EventQuery(t *testing.T) {
	client := getXuperChainClient(t)

	res, err := client.QueryEvent()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)

}
func TestXuperChainClient_EventRemove(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.RemoveEventReqDataBody{
		EventId: "f4c6912f5b0540399ff080ef798763fa",
	}

	res, err := client.RemoveEvent(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
