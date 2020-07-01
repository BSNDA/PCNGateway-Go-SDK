package fabric

import (
	"fmt"
	config2 "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/user"
	"testing"
)

func TestFabricClient_RegisterUser(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterReqDataBody{
		Name:   "user01",
		Secret: "123456",
	}

	res, err := fabricClient.RegisterUser(body)
	if err != nil {
		t.Fatal(err)
	}

	if res.Header.Code != 0 {
		t.Fatal(res.Header.Msg)
	}

}

func TestFabricClient_EnrollUser(t *testing.T) {

	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	body := req.RegisterReqDataBody{
		Name:   "user01",
		Secret: "123456",
	}

	res := fabricClient.EnrollUser(body)

	if res != nil {
		t.Fatal(res.Error())
	}
}

func TestFabricClient_LoadUser(t *testing.T) {
	config, err := config2.NewMockConfig()

	if err != nil {
		t.Fatal(err.Error())
	}

	fabricClient, err := InitFabricClient(config)

	if err != nil {
		t.Fatal(err.Error())
	}

	fmt.Println(*fabricClient.Users["abcde"])
}
