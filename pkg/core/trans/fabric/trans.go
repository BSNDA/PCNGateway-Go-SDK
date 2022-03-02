package fabric

import (
	"encoding/base64"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	protos_utils "github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/protos/utils"
	cryptocomm "github.com/BSNDA/bsn-sdk-crypto/common"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"strings"
)

func CreateRequest(user *msp.UserData, request *TransRequest) (data string, txId string, err error) {

	txh, err := NewTranHeader(request.ChannelId, user)

	if err != nil {
		return "", "", err
	}

	proposal, err := createInvokerProposal(txh, request.GetRequest(), request.ChaincodeId)

	if err != nil {
		return "", "", err
	}

	pb_proposal, err := signProposal(proposal.Proposal, user)

	if err != nil {
		return "", "", err
	}

	data = getRequestData(pb_proposal)

	return data, string(txh.TransactionID()), nil

}

func NewTranHeader(channelID string, usre *msp.UserData) (fab.TransactionHeader, error) {

	nonce, err := cryptocomm.GetRandomNonce()

	creator, err := usre.Serialize()

	id, err := cryptocomm.ComputeTxnID(nonce, creator)

	txnID := TransactionHeader{
		id:        fab.TransactionID(id),
		creator:   creator,
		nonce:     nonce,
		channelID: channelID,
	}

	return &txnID, err

}

func createInvokerProposal(txh fab.TransactionHeader, request *fab.ChaincodeInvokeRequest, chaincodeID string) (*fab.TransactionProposal, error) {

	argsArray := make([][]byte, len(request.Args)+1)
	argsArray[0] = []byte(request.Fcn)
	for i, arg := range request.Args {
		argsArray[i+1] = arg
	}
	ccis := &pb.ChaincodeInvocationSpec{ChaincodeSpec: &pb.ChaincodeSpec{
		Type: pb.ChaincodeSpec_GOLANG, ChaincodeId: &pb.ChaincodeID{Name: chaincodeID},
		Input: &pb.ChaincodeInput{Args: argsArray}}}

	proposal, _, err := protos_utils.CreateChaincodeProposalWithTxIDNonceAndTransient(string(txh.TransactionID()), common.HeaderType_ENDORSER_TRANSACTION, txh.ChannelID(), ccis, txh.Nonce(), txh.Creator(), request.TransientMap)
	if err != nil {
		return nil, errors.New("failed to create chaincode proposal")
	}

	tp := fab.TransactionProposal{
		TxnID:    txh.TransactionID(),
		Proposal: proposal,
	}

	return &tp, nil
}

func signProposal(proposal *pb.Proposal, user *msp.UserData) (*pb.SignedProposal, error) {

	proposalBytes, err := proto.Marshal(proposal)

	if err != nil {
		return nil, errors.New("mashal proposal failed")
	}

	//digest, err := crypto.GetHash(proposalBytes)

	signature, err := user.Sign(proposalBytes) // secp256r1.SignECDSA(user.PrivateKey, digest)

	return &pb.SignedProposal{ProposalBytes: proposalBytes, Signature: signature}, nil

}

func getRequestData(signedProposal *pb.SignedProposal) string {

	proposalBytes, _ := proto.Marshal(signedProposal)

	base64 := base64.StdEncoding.EncodeToString(proposalBytes)

	return base64
}

func getRequest(signedProposal *pb.SignedProposal) fab.ProcessProposalRequest {

	request := fab.ProcessProposalRequest{SignedProposal: signedProposal}

	return request

}

func ParseRequest(data string) error {
	by, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return err
	}

	signproposal := &pb.SignedProposal{}

	err = proto.Unmarshal(by, signproposal)

	if err != nil {
		return err
	}

	proposal := &pb.Proposal{}

	err = proto.Unmarshal(signproposal.ProposalBytes, proposal)

	if err != nil {
		return err
	}

	getProposalHeader(proposal.Header)
	getProposalPayload(proposal.Payload)

	return nil
}

func getProposalHeader(data []byte) error {

	header := &common.Header{}

	err := proto.Unmarshal(data, header)

	if err != nil {
		return err
	}

	channelHeader := &common.ChannelHeader{}

	err = proto.Unmarshal(header.ChannelHeader, channelHeader)

	if err != nil {
		return err
	}

	fmt.Println("ChannelId:", channelHeader.ChannelId)
	fmt.Println("TxId:", channelHeader.TxId)
	//fmt.Println("TxId:",channelHeader.TxId)

	return err

}

func getProposalPayload(data []byte) error {

	ccPropPayload := &peer.ChaincodeProposalPayload{}

	err := proto.Unmarshal(data, ccPropPayload)

	if err != nil {
		return err
	}

	ccis := &peer.ChaincodeInvocationSpec{}

	err = proto.Unmarshal(ccPropPayload.Input, ccis)

	if err != nil {
		return err
	}

	fmt.Println("ChaincodeId:", ccis.ChaincodeSpec.ChaincodeId.Name)
	fmt.Println("Fcn:", string(ccis.ChaincodeSpec.Input.Args[0]))

	return nil
}

func ConvertToTran(transData string) (*pb.ProcessedTransaction, error) {
	transBytes, err := base64.StdEncoding.DecodeString(transData)
	if err != nil {
		return nil, errors.New("convert trans data has error")
	}

	trans := &pb.ProcessedTransaction{}

	err = proto.Unmarshal(transBytes, trans)
	if err != nil {
		return nil, errors.New("convert trans bytes has error")
	}

	return trans, nil
}

func ConvertTransToJson(tx *pb.ProcessedTransaction) (string, error) {

	var sb strings.Builder

	err := protolator.DeepMarshalJSON(&sb, tx)
	if err != nil {
		return "", err
	}

	return sb.String(), err

}
