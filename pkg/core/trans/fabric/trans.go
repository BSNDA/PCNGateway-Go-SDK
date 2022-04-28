package fabric

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/pkg/errors"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric/fab"
	protos_utils "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/fabric/utils"
	cryptocomm "github.com/BSNDA/bsn-sdk-crypto/common"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-config/protolator"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strings"
)

func CreateRequest(user *msp.UserData, request *TransRequest) (data string, proposal *fab.TransactionProposal, err error) {

	txh, err := NewTranHeader(request.ChannelId, user)

	if err != nil {
		return "", nil, err
	}

	proposal, err = createInvokerProposal(txh, request.GetRequest(), request.ChaincodeId)

	if err != nil {
		return "", nil, err
	}

	pb_proposal, err := signProposal(proposal.Proposal, user)

	if err != nil {
		return "", nil, err
	}

	data = getRequestData(pb_proposal)

	return data, proposal, nil

}

func ComputeTxnID(nonce, creator []byte, hash key.HashProvider) (string, error) {

	if hash == nil {
		hash = &key.SHA256Hash{}
	}

	b := append(nonce, creator...)
	digest := hash.Hash(b)
	id := hex.EncodeToString(digest)
	return id, nil
}

func NewTranHeader(channelID string, usre *msp.UserData) (fab.TransactionHeader, error) {

	nonce, err := cryptocomm.GetRandomNonce()

	creator, err := usre.Serialize()

	id, err := ComputeTxnID(nonce, creator, usre.Hash())

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

func GetProposalTxId(data []byte) (string, error) {

	header := &common.Header{}

	err := proto.Unmarshal(data, header)

	if err != nil {
		return "", err
	}

	channelHeader := &common.ChannelHeader{}

	err = proto.Unmarshal(header.ChannelHeader, channelHeader)

	if err != nil {
		return "", err
	}

	txId := channelHeader.TxId

	return txId, nil

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

func NewEnvelope(user *msp.UserData, proposal *fab.TransactionProposal, resps []*fab.TransactionProposalResponse) (*common.Envelope, error) {

	tx, err := NewTransaction(proposal, resps)
	if err != nil {
		return nil, err
	}

	txBytes, err := protos_utils.GetBytesTransaction(tx)
	if err != nil {
		return nil, err
	}

	hdr, err := protos_utils.GetHeader(proposal.Proposal.Header)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal proposal header failed")
	}

	payload := common.Payload{Header: hdr, Data: txBytes}

	payloadBytes, err := proto.Marshal(&payload)
	if err != nil {
		return nil, errors.WithMessage(err, "marshaling of payload failed")
	}

	var signature []byte
	if user != nil {
		signature, err = user.Sign(payloadBytes)
		if err != nil {
			return nil, errors.WithMessage(err, "envelope sign failed")
		}
	}

	return &common.Envelope{Payload: payloadBytes, Signature: signature}, nil
}

func NewTransaction(proposal *fab.TransactionProposal, resps []*fab.TransactionProposalResponse) (*peer.Transaction, error) {

	hdr, err := protos_utils.GetHeader(proposal.Header)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal proposal header failed")
	}

	// the original payload
	pPayl, err := protos_utils.GetChaincodeProposalPayload(proposal.Payload)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal proposal payload failed")
	}

	// get header extensions so we have the visibility field
	//hdrExt, err := protos_utils.GetChaincodeHeaderExtension(hdr)
	//if err != nil {
	//	return nil, err
	//}

	responsePayload := resps[0].ProposalResponse.Payload
	if vprErr := validateProposalResponses(resps); vprErr != nil {
		return nil, vprErr
	}

	// fill endorsements
	endorsements := make([]*pb.Endorsement, len(resps))
	for n, r := range resps {
		endorsements[n] = r.ProposalResponse.Endorsement
	}

	// create ChaincodeEndorsedAction
	cea := &pb.ChaincodeEndorsedAction{ProposalResponsePayload: responsePayload, Endorsements: endorsements}

	// obtain the bytes of the proposal payload that will go to the transaction
	propPayloadBytes, err := protos_utils.GetBytesProposalPayloadForTx(pPayl, nil)
	if err != nil {
		return nil, err
	}

	// serialize the chaincode action payload
	cap := &pb.ChaincodeActionPayload{ChaincodeProposalPayload: propPayloadBytes, Action: cea}
	capBytes, err := protos_utils.GetBytesChaincodeActionPayload(cap)
	if err != nil {
		return nil, err
	}

	// create a transaction
	taa := &pb.TransactionAction{Header: hdr.SignatureHeader, Payload: capBytes}
	taas := make([]*pb.TransactionAction, 1)
	taas[0] = taa

	return &peer.Transaction{Actions: taas}, nil
}

func validateProposalResponses(responses []*fab.TransactionProposalResponse) error {
	for _, r := range responses {
		if r.ProposalResponse.Response.Status < int32(common.Status_SUCCESS) || r.ProposalResponse.Response.Status >= int32(common.Status_BAD_REQUEST) {
			return errors.Errorf("proposal response was not successful, error code %d, msg %s", r.ProposalResponse.Response.Status, r.ProposalResponse.Response.Message)
		}
	}
	return nil
}
