package xuperchain

import (
	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/user"
	res "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/xuperchain/user"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *XuperChainClient) RegisterUser(body req.RegisterUserReqDataBody) (*res.RegisterUserResData, error) {
	url := c.GetURL("/api/xuperchain/v1/user/register")

	data := &req.RegisterUserReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	resData := &res.RegisterUserResData{}

	err = json.Unmarshal(resBytes, resData)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}
	if !c.Verify(resData.Mac, resData.GetEncryptionValue()) {
		return nil, errors.New("sign has error")
	}

	return resData, nil
}
