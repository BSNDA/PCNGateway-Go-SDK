package fabric

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	nodereq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/node"
	noderes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric"
	"github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"

	blockconvert "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric"
)

const (
	ReqChainCode  = "node/reqChainCode"
	SDKTran       = "node/trans"
	GetTransInfo  = "node/getTransaction"
	GetTransData  = "node/getTransdata"
	GetBlockInfo  = "node/getBlockInfo"
	GetBlockData  = "node/getBlockData"
	GetLedgerInfo = "node/getLedgerInfo"
)

// SdkTran Dapp transaction in key upload mode
func (c *FabricClient) SdkTran(body nodereq.TransReqDataBody, user *msp.UserData) (*noderes.TranResData, error) {

	var err error
	if user == nil {
		user, err = c.LoadUser(body.UserName)
		if err != nil {
			return nil, errors.WithMessagef(err, "user [%s] load failed", body.UserName)
		}
	}

	request := body.GetTransRequest(c.appInfo.ChannelId)
	transData, _, err := fabric.CreateRequest(user, request)

	if err != nil {
		return nil, err
	}

	data := &nodereq.SdkTransReqData{}
	data.Header = c.GetHeader()
	data.Body = nodereq.SdkTransReqDataBody{
		TransData: transData,
	}
	res := &noderes.TranResData{}
	err = c.Call(SDKTran, data, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil

}

// ReqChainCode Dapp transaction in public key trust mode
func (c *FabricClient) ReqChainCode(body nodereq.TransReqDataBody) (*noderes.TranResData, error) {

	req := &nodereq.TransReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.TranResData{}

	err := c.Call(ReqChainCode, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", ReqChainCode)
	}
	return res, nil
}

// GetTransInfo query fabric transaction
func (c *FabricClient) GetTransInfo(body nodereq.TxTransReqDataBody) (*noderes.TransactionResData, error) {

	req := &nodereq.TxTransReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.TransactionResData{}

	err := c.Call(GetTransInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetTransInfo)
	}
	return res, nil
}

// GetTransData  query fabric transaction,
// but the return data is "peer.ProcessedTransaction" serialized bytes after Base64 encoding .
func (c *FabricClient) GetTransData(body nodereq.TxTransReqDataBody) (*noderes.TranDataRes, *pb.ProcessedTransaction, error) {

	req := &nodereq.TxTransReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.TranDataRes{}

	err := c.Call(GetTransData, req, res)
	if err != nil {
		return nil, nil, errors.WithMessagef(err, "call %s has error", GetTransData)
	}

	trans := &pb.ProcessedTransaction{}
	trans, err = blockconvert.ConvertToTran(res.Body.TransData)
	if err != nil {
		logger.Error("convertTrans failed,errmessage：", err)
		return res, nil, err
	}
	return res, trans, nil
}

// GetBlockInfo query fabric block data
func (c *FabricClient) GetBlockInfo(body nodereq.BlockReqDataBody) (*noderes.BlockResData, error) {

	req := &nodereq.BlockReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.BlockResData{}

	err := c.Call(GetBlockInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetBlockInfo)
	}
	return res, nil
}

// GetBlockData query fabric block data
// but the return data is "common.Block" serialized bytes after Base64 encoding .
func (c *FabricClient) GetBlockData(body nodereq.BlockReqDataBody) (*noderes.BlockDataRes, *common.Block, error) {

	req := &nodereq.BlockReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &noderes.BlockDataRes{}

	err := c.Call(GetBlockData, req, res)
	if err != nil {
		return nil, nil, errors.WithMessagef(err, "call %s has error", GetBlockData)
	}

	block := &common.Block{}
	if len(res.Body.BlockData) > 0 {
		block, err = blockconvert.ConvertToBlock(res.Body.BlockData)
		if err != nil {
			logger.Error("convertBlock failed,errmessage：", err)
			return res, nil, err
		}
	}

	return res, block, nil
}

// GetBlockData query fabric ledger data
func (c *FabricClient) GetLedgerInfo() (*noderes.LedgerResData, error) {

	req := &nodereq.LedgerReqData{}
	req.Header = c.GetHeader()

	res := &noderes.LedgerResData{}

	err := c.Call(GetLedgerInfo, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", GetLedgerInfo)
	}

	return res, nil
}
