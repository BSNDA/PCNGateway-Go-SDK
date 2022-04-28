package fisco_bcos

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/node"
	noderes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fiscobcos/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fiscobcos"
	"github.com/wonderivan/logger"
	"strconv"
)

// ReqChainCode Dapp transaction in public key trust mode
func (c *FiscoBcosClient) ReqChainCode(body nodereq.TransReqDataBody) (*noderes.TransResData, error) {
	url := c.GetURL("/api/fiscobcos/v1/node/reqChainCode")

	data := &nodereq.TransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

// GetBlockInfo query fisco bcos block data
func (c *FiscoBcosClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getBlockInfo")
	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

// GetBlockHeight query fisco bcos block height
func (c *FiscoBcosClient) GetBlockHeight() (*noderes.BlockHeightResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getBlockHeight")
	data := &nodereq.LedgerReqData{}
	data.Header = c.GetHeader()

	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

// GetTxCount query fisco bcos total number of transactions in Dapp
func (c *FiscoBcosClient) GetTxCount() (*noderes.BlockHeightResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getTxCount")
	data := &nodereq.LedgerReqData{}
	data.Header = c.GetHeader()

	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

// GetTxCountByBlockNumber query fisco bcos total number of transactions in block
func (c *FiscoBcosClient) GetTxCountByBlockNumber(blockNumber int64) (*noderes.BlockHeightResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getTxCountByBlockNumber")
	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()
	data.Body = nodereq.BlockReqDataBody{
		BlockNumber: strconv.FormatInt(blockNumber, 10),
	}
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)
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

// GetTxCountByBlockNumber query fisco bcos transactions receipt
func (c *FiscoBcosClient) GetTxReceiptByTxHash(body nodereq.TxReqDataBody) (*noderes.BlockTxReceiptResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getTxReceiptByTxHash")

	data := &nodereq.TxReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)

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

// GetTxCountByBlockNumber query fisco bcos transactions info
func (c *FiscoBcosClient) GetTxInfoByTxHash(body nodereq.TxReqDataBody) (*noderes.BlockTxResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/node/getTxInfoByTxHash")

	data := &nodereq.TxReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)

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

func (c *FiscoBcosClient) getTransDataAndConstant(data nodereq.TransData) (string, bool, error) {
	blockLimit, err := c.getBlockLimit()
	if err != nil {
		return "", false, err
	}
	groupId, err := c.getGroupId()
	if err != nil {
		return "", false, err
	}

	key, err := c.getUser(data.UserName)
	if err != nil {
		return "", false, err
	}

	tx, constant, err := fiscobcos.TransData(data.Contract.ContractAbi, data.Contract.ContractAddress, data.FuncName, data.Args, groupId, blockLimit, nil, c.isSM(), key)

	if err != nil {
		return "", false, err
	} else {
		return tx, constant, nil
	}

}

func (c *FiscoBcosClient) getTransData(data nodereq.TransData) (string, error) {
	blockLimit, err := c.getBlockLimit()
	if err != nil {
		return "", err
	}
	groupId, err := c.getGroupId()
	if err != nil {
		return "", err
	}

	key, err := c.getUser(data.UserName)
	if err != nil {
		return "", err
	}

	tx, _, err := fiscobcos.TransData(data.Contract.ContractAbi, data.Contract.ContractAddress, data.FuncName, data.Args, groupId, blockLimit, nil, c.isSM(), key)

	if err != nil {
		return "", err
	} else {
		return tx, nil
	}

}

// Trans Dapp transaction in key upload mode
func (c *FiscoBcosClient) Trans(data nodereq.TransData) (*noderes.TransResData, error) {
	url := c.GetURL("/api/fiscobcos/v1/node/trans")

	tx, constant, err := c.getTransDataAndConstant(data)
	if err != nil {
		return nil, err
	}

	reqData := &nodereq.KeyTransReqData{}
	reqData.Header = c.GetHeader()
	reqData.Body = nodereq.KeyTransReqDataBody{
		ContractName: data.Contract.ContractName,
		TransData:    tx,
	}

	if constant {
		reqData.Body.ContractAbi = data.Contract.ContractAbi
		reqData.Body.ContractAddress = data.Contract.ContractAddress
	}

	reqData.Mac = c.Sign(reqData.GetEncryptionValue())

	reqBytes, _ := json.Marshal(reqData)

	resBytes, err := http.SendPost(reqBytes, url)
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
