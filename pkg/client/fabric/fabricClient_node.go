package fabric

import (
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	noderes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/node"
	"github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/wonderivan/logger"

	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric"
	blockconvert "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
)

func (c *FabricClient) SdkTran(body nodereq.TransReqDataBody) (*noderes.TranResData, error) {

	url := c.GetURL("/api/fabric/v1/node/trans")

	user, err := c.GetUser(body.UserName)
	if err != nil {
		return nil, err
	}

	request := body.GetTransRequest(c.Config.GetAppInfo().ChannelId)

	transData, _, err := fabric.CreateRequest(user, request)

	if err != nil {
		return nil, err
	}

	data := &nodereq.SdkTransReqData{}
	data.Header = c.GetHeader()
	data.Body = nodereq.SdkTransReqDataBody{
		TransData: transData,
	}
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.TranResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil

}

func (c *FabricClient) ReqChainCode(body nodereq.TransReqDataBody) (*noderes.TranResData, error) {
	url := c.GetURL("/api/fabric/v1/node/reqChainCode")

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
	res := &noderes.TranResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}

func (c *FabricClient) GetTransInfo(body nodereq.TxTransReqDataBody) (*noderes.TransactionResData, error) {
	url := c.GetURL("/api/fabric/v1/node/getTransaction")

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
	res := &noderes.TransactionResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
func (c *FabricClient) GetTransData(body nodereq.TxTransReqDataBody) (*noderes.TranDataRes, *pb.ProcessedTransaction, error) {
	url := c.GetURL("/api/fabric/v1/node/getTransdata")

	data := &nodereq.TxTransReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, nil, err
	}
	res := &noderes.TranDataRes{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, nil, err
	}
	trans := &pb.ProcessedTransaction{}
	trans, err = blockconvert.ConvertToTran(res.Body.TransData)
	if err != nil {
		logger.Error("convertTrans failed,errmessage：", err)
		return res, nil, err
	}
	return res, trans, nil
}

func (c *FabricClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {

	url := c.GetURL("/api/fabric/v1/node/getBlockInfo")

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
func (c *FabricClient) GetBlockData(body nodereq.BlockReqDataBody) (*noderes.BlockDataRes, *common.Block, error) {

	url := c.GetURL("/api/fabric/v1/node/getBlockData")

	data := &nodereq.BlockReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, nil, err
	}

	res := &noderes.BlockDataRes{}
	block := &common.Block{}
	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, nil, err
	}

	if len(res.Body.BlockData) > 0 {
		block, err = blockconvert.ConvertToBlock(res.Body.BlockData)
		if err != nil {
			logger.Error("convertBlock failed,errmessage：", err)
			return res, nil, err
		}
	}

	return res, block, nil
}

func (c *FabricClient) GetLedgerInfo() (*noderes.LedgerResData, error) {

	url := c.GetURL("/api/fabric/v1/node/getLedgerInfo")

	data := &nodereq.LedgerReqData{}
	data.Header = c.GetHeader()
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &noderes.LedgerResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil
}
