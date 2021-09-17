package cita

import (
	"fmt"
	"testing"

	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/user"
)

func TestCitaClient_RegisterUser(t *testing.T) {

	citaClient := getCitaClient(t)

	body := req.RegisterReqDataBody{
		UserId: "testcurel",
	}

	res, err := citaClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println(citaClient.Verify(res.Mac, res.GetEncryptionValue()))
}
