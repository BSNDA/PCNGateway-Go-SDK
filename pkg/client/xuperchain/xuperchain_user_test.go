package xuperchain

import (
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/user"
	"testing"
)

func TestXuperChainClient_RegisterUser(t *testing.T) {
	config, err := config2.NewMockTestFiscoSMConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	client, err := NewXuperChainClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterUserReqDataBody{
		UserId: "test0623",
	}

	res, err := client.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
