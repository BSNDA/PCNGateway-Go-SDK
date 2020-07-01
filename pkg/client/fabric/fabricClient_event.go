package fabric

import (
	resBase "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/event"
	eventres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/event"

	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *FabricClient) EventRegister(body eventreq.RegisterReqDataBody) (*eventres.RegisterResData, error) {

	url := c.GetURL("/api/fabric/v1/chainCode/event/register")

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

	res := &eventres.RegisterResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) EventQuery() (*eventres.QueryResData, error) {

	url := c.GetURL("/api/fabric/v1/chainCode/event/query")

	data := &eventreq.QueryReqData{}
	data.Header = c.GetHeader()
	//data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &eventres.QueryResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) EventRemove(body eventreq.RemoveReqDataBody) (*resBase.BaseResModel, error) {

	url := c.GetURL("/api/fabric/v1/chainCode/event/remove")

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

	res := &resBase.BaseResModel{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
