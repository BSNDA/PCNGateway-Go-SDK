package msp

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/sign"
	pb_msp "github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/protos/msp"
	"github.com/golang/protobuf/proto"
)

type UserData struct {
	UserName              string
	AppCode               string
	MspId                 string
	EnrollmentCertificate []byte

	PrivateKey interface{}

	sign sign.SignHandle
}

func (u *UserData) SetSignHandle(s sign.SignHandle) {
	u.sign = s
}

func (u *UserData) Sign(digest []byte) (signature []byte, err error) {
	hash, err := u.sign.Hash(digest)
	if err != nil {
		return nil, errors.New("data hash failed")
	}
	return u.sign.Sign(hash)
}

func (u *UserData) Serialize() ([]byte, error) {
	serializedIdentity := &pb_msp.SerializedIdentity{
		Mspid:   u.MspId,
		IdBytes: u.EnrollmentCertificate,
	}
	identity, err := proto.Marshal(serializedIdentity)
	if err != nil {
		return nil, errors.New("marshal serializedIdentity failed")
	}
	return identity, nil
}
