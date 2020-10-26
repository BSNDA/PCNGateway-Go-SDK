package xuperchain

import (
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/event"
	res "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/xuperchain/event"

	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *XuperChainClient) RegisterEvent(body req.RegisterEventReqDataBody) (*res.QueryEventResData, error) {

	url := c.GetURL("/api/xuperchain/v1/event/register")

	data := &req.RegisterEventReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &res.QueryEventResData{}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *XuperChainClient) QueryEvent() (*res.QueryEventResData, error) {

	url := c.GetURL("/api/xuperchain/v1/event/query")

	data := &req.QueryEventReqData{}
	data.Header = c.GetHeader()
	//data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &res.QueryEventResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *XuperChainClient) RemoveEvent(body req.RemoveEventReqDataBody) (*res.RemoveEventResData, error) {

	url := c.GetURL("/api/xuperchain/v1/event/remove")

	data := &req.RemoveEventReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &res.RemoveEventResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
