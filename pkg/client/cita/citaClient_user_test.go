package cita

import (
	"fmt"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/user"
	"testing"
)

func TestCitaClient_RegisterUser(t *testing.T) {

	citaClient := getCitaClient(t)

	body := req.RegisterReqDataBody{
		UserId: "test10281116",
	}

	res, err := citaClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}
	fmt.Println(res)
	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))
}
