package fabric

import (
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/user"
	"testing"
)

func TestFabricClient_RegisterUser(t *testing.T) {
	fabricClient := getFabricClient(t)
	body := req.RegisterReqDataBody{
		Name:             "user20220421",
		Secret:           "123456",
		ExtendProperties: "{'key1':'abc'}", //用户拓展属性，json格式，非必填
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

	fabricClient := getFabricClient(t)

	body := req.RegisterReqDataBody{
		Name:   "user20220421",
		Secret: "123456",
	}

	_, err := fabricClient.EnrollUser(body)

	if err != nil {
		t.Fatal(err.Error())
	}
}
