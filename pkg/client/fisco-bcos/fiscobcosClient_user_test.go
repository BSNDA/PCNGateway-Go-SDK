package fisco_bcos

import (
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/user"
	"testing"
)

func TestFiscoBcosClient_RegisterUser(t *testing.T) {

	fiscoClient := getSMClient(t)

	body := req.RegisterReqDataBody{
		UserId: "test0927",
	}

	res, err := fiscoClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestFiscoBcosClient_RegisterUserK1(t *testing.T) {

	fiscoClient := getK1Client(t)

	body := req.RegisterReqDataBody{
		UserId: "test0611",
	}

	res, err := fiscoClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
