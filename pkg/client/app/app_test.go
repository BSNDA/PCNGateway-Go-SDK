package app

import (
	"fmt"
	"testing"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req"
)

func TestGetAppInfo(t *testing.T) {

	api := "http://192.168.1.43:17502"

	reqData := req.AppInfoReqData{}

	header := base.ReqHeader{
		UserCode: "USER0001202004161009309407413",
		AppCode:  "app0001202004161017141233920",
	}

	reqData.Header = header

	res, _ := GetAppInfo(&reqData, api, "")

	if res.Header.Code != 0 {
		fmt.Println(res.Header.Msg)
	} else {
		fmt.Println(res.Body)
	}
}

func TestName(t *testing.T) {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = a + b

	fmt.Println(c)
	fmt.Println(0.1 + 0.2)
	fmt.Println(0.1+0.2 == 0.3)
}
