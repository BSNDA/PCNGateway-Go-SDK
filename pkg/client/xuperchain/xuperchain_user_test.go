package xuperchain

import (
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/user"
	"testing"
)

func TestXuperChainClient_RegisterUser(t *testing.T) {
	client := getXuperChainClient(t)
	body := req.RegisterUserReqDataBody{
		UserId: "zxl072201708",
	}
	res, err := client.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}
	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
}
