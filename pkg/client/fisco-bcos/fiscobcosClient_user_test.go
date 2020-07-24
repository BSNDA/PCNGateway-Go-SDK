package fisco_bcos

import (
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/user"
	"testing"
)

func TestFiscoBcosClient_RegisterUser(t *testing.T) {

	config, err := config2.NewMockTestFiscoSMConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterReqDataBody{
		UserId: "test0623",
	}

	res, err := fabricClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}

func TestFiscoBcosClient_RegisterUserK1(t *testing.T) {

	config, err := config2.NewMockTestFiscoK1Config()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := NewFiscoBcosClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterReqDataBody{
		UserId: "test0611",
	}

	res, err := fabricClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
