package cita

import (
	"encoding/json"
	"fmt"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/node"
	noderes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/cita"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

//DApp transaction in public key trust mode
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

//get block information
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

//get block height
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

//get transaction Receipt
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

//get transaction information
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

//DApp transaction in public key upload mode
func (c *CitaClient) Trans(data nodereq.TransData) (*noderes.TransResData, error) {
	if c.Config.GetAppInfo().CAType == enum.AppCaType_Trust {
		return nil, errors.New("the trusteeship application cannot call the api")
	}
	url := c.GetURL("/api/cita/v1/node/trans")

	tx, err := c.getTransData(data)
	if err != nil {
		return nil, err
	}

	reqData := &nodereq.KeyTransReqData{}
	reqData.Header = c.GetHeader()
	reqData.Body = nodereq.KeyTransReqDataBody{
		ContractName: data.Contract.ContractName,
		TransData:    tx,
	}
	reqData.Mac = c.Sign(reqData.GetEncryptionValue())

	reqBytes, _ := json.Marshal(reqData)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}
	res := &noderes.TransResData{}

	err = json.Unmarshal(resBytes, res)

	fmt.Println(c.Verify(res.Mac, res.GetEncryptionValue()))

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

//splicing key upload mode transaction string
func (c *CitaClient) getTransData(data nodereq.TransData) (string, error) {
	blockLimit, err := c.getBlockLimit()
	if err != nil {
		return "", err
	}
	chainId, version, err := c.getBaseInfo()
	if err != nil {
		return "", err
	}

	key, err := c.getUser(data.UserName)
	if err != nil {
		return "", err
	}

	tx, _, err := cita.TransData(data.Contract.ContractAbi, data.Contract.ContractAddress, data.FuncName, data.Args, blockLimit, chainId, version, c.isSM(), key)

	if err != nil {
		return "", err
	} else {
		return tx, nil
	}

}
