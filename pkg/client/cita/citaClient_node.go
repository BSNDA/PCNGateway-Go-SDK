package cita

import (
	"encoding/json"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/node"
	noderes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *CitaClient) ReqChainCode(body nodereq.TransReqDataBody) (*noderes.TransResData, error) {
	url := c.GetURL("/api/cita/v1/node/reqChainCode")

	data := &nodereq.TransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.TransResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *CitaClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {

	url := c.GetURL("/api/cita/v1/node/getBlockInfo")
	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.BlockResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *CitaClient) GetBlockHeight() (*noderes.BlockHeightResData, error) {

	url := c.GetURL("/api/cita/v1/node/getBlockHeight")
	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()

	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.BlockHeightResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *CitaClient) GetTxReceiptByTxHash(body nodereq.TxTransReqDataBody) (*noderes.BlockTxReceiptResData, error) {

	url := c.GetURL("/api/cita/v1/node/getTxReceiptByTxHash")

	data := &nodereq.TxTransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.BlockTxReceiptResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *CitaClient) GetTxInfoByTxHash(body nodereq.TxTransReqDataBody) (*noderes.BlockTxResData, error) {

	url := c.GetURL("/api/cita/v1/node/getTxInfoByTxHash")

	data := &nodereq.TxTransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.BlockTxResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
