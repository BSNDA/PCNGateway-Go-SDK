package cita

import (
	"encoding/json"
	userReq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/user"
	userRes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/user"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *CitaClient) RegisterUser(body userReq.RegisterReqDataBody) (*userRes.RegisterResData, error) {

	url := c.GetURL("/api/cita/v1/user/register")

	data := &userReq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userRes.RegisterResData{}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}
	//if !c.Verify(res.Mac, res.GetEncryptionValue()) {
	//	return nil, errors.New("sign has error")
	//}

	return res, nil
}
