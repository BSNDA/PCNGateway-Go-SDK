package cita

import (
	"encoding/json"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/event"
	eventres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/event"
	"github.com/wonderivan/logger"
)

//register cita event listeners
func (c *CitaClient) EventRegister(body eventreq.RegisterReqDataBody) (*eventres.RegisterEventResData, error) {
	url := c.GetURL("/api/cita/v1/event/register")

	data := &eventreq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

//query event list
func (c *CitaClient) EventQuery() (*eventres.QueryEventResData, error) {
	url := c.GetURL("/api/cita/v1/event/query")

	data := &eventreq.QueryReqData{}
	data.Header = c.GetHeader()

	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

//remove event
func (c *CitaClient) EventRemove(body eventreq.RemoveReqDataBody) (*eventres.RemoveEventResData, error) {
	url := c.GetURL("/api/cita/v1/event/remove")

	data := &eventreq.RemoveReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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
