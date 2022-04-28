package app

import (
	"encoding/json"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res"
)

// GetAppInfo query basic information of Dapp
func GetAppInfo(data *req.AppInfoReqData, baseApi string, cert string) (*res.AppInfoResData, error) {

	url := baseApi + "/api/app/getAppInfo"

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)

	if err != nil {
		return nil, err
	}

	resData := &res.AppInfoResData{}

	err = json.Unmarshal(resBytes, resData)

	if err != nil {
		return nil, err
	}

	return resData, nil

}
