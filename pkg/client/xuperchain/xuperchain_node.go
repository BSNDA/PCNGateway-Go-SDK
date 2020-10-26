// @Title  xuperchain_node
// @Description
// @Author  zxl  2020/7/22 19:32
// @Version 1.0.0
// @Update  2020/7/22 19:32
package xuperchain

import (
	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	req "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/xuperchain/node"
	res "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/xuperchain/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
)

func (c *XuperChainClient) SdkTran(body req.UPukCallContractReqDataReqDataBody) (*res.UPukCallContractResData, error) {
	url := c.GetURL("/api/xuperchain/v1/node/trans")

	data := &req.UPukCallContractReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	resData := &res.UPukCallContractResData{}

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
func (c *XuperChainClient) ReqChainCode(body req.CallContractReqDataReqDataBody) (*res.CallContractResData, error) {
	url := c.GetURL("/api/xuperchain/v1/node/reqChainCode")

	data := &req.CallContractReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	resData := &res.CallContractResData{}

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

func (c *XuperChainClient) GetTxInfo(body req.GetTxInfoReqDataBody) (*res.GetTxInfoResData, error) {
	url := c.GetURL("/api/xuperchain/v1/node/getTxInfoByTxHash")

	data := &req.GetTxInfoReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	resData := &res.GetTxInfoResData{}

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

func (c *XuperChainClient) GetBlockInfo(body req.GetBlockInfoReqDataBody) (*res.GetBlockInfoResData, error) {
	url := c.GetURL("/api/xuperchain/v1/node/getBlockInfo")

	data := &req.GetBlockInfoReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	resData := &res.GetBlockInfoResData{}

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
