package app

import (
	"encoding/json"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
)

func GetAppInfo(data *req.AppInfoReqData, baseApi string, cert string) (*res.AppInfoResData, error) {

	url := baseApi + "/api/app/getAppInfo"

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, cert)

	if err != nil {
		return nil, err
	}

	res := &res.AppInfoResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		return nil, err
	}

	return res, nil

}
