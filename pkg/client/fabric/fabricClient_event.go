package fabric

import (
	resBase "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	eventreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/event"
	eventres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/event"
	"github.com/pkg/errors"
)

const (
	EventRegister      = "chainCode/event/register"
	BlockEventRegister = "chainCode/event/blockRegister"
	EventQuery         = "chainCode/event/query"
	EventRemove        = "chainCode/event/remove"
)

// EventRegister register fabric transaction event listeners
func (c *FabricClient) EventRegister(body eventreq.RegisterReqDataBody) (*eventres.RegisterResData, error) {

	req := &eventreq.RegisterReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.RegisterResData{}

	err := c.Call(EventRegister, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventRegister)
	}
	return res, nil
}

// BlockEventRegister register fabric block event listeners
func (c *FabricClient) BlockEventRegister(body eventreq.RegisterReqDataBody) (*eventres.RegisterResData, error) {

	req := &eventreq.RegisterReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &eventres.RegisterResData{}

	err := c.Call(BlockEventRegister, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", BlockEventRegister)
	}
	return res, nil
}

// EventQuery query fabric event list
func (c *FabricClient) EventQuery() (*eventres.QueryResData, error) {

	req := &eventreq.QueryReqData{}
	req.Header = c.GetHeader()

	res := &eventres.QueryResData{}

	err := c.Call(EventQuery, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventQuery)
	}
	return res, nil
}

// EventRemove remove fabric event
func (c *FabricClient) EventRemove(body eventreq.RemoveReqDataBody) (*resBase.BaseResModel, error) {

	req := &eventreq.RemoveReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &resBase.BaseResModel{}

	err := c.Call(EventRemove, req, res)

	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EventRemove)
	}
	return res, nil
}
