package fisco_bcos

import (
	"encoding/json"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/event"
	eventres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fiscobcos/event"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

// EventRegister register fisco bcos transaction or block event  listeners
func (c *FiscoBcosClient) EventRegister(body eventreq.RegisterReqDataBody) (*eventres.RegisterEventResData, error) {
	url := c.GetURL("/api/fiscobcos/v1/event/register")

	data := &eventreq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &eventres.RegisterEventResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

// EventQuery query fisco bcos event list
func (c *FiscoBcosClient) EventQuery() (*eventres.QueryEventResData, error) {
	url := c.GetURL("/api/fiscobcos/v1/event/query")

	data := &eventreq.QueryReqData{}
	data.Header = c.GetHeader()

	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &eventres.QueryEventResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

// EventRemove remove fisco bcos event
func (c *FiscoBcosClient) EventRemove(body eventreq.RemoveReqDataBody) (*eventres.RemoveEventResData, error) {
	url := c.GetURL("/api/fiscobcos/v1/event/remove")

	data := &eventreq.RemoveReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &eventres.RemoveEventResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
