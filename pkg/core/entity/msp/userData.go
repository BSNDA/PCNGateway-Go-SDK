package msp

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/golang/protobuf/proto"
	pb_msp "github.com/hyperledger/fabric-protos-go/msp"
)

type UserData struct {
	UserName              string
	AppCode               string
	MspId                 string
	EnrollmentCertificate []byte

	PrivateKey key.PrivateKeyProvider

	TxHash key.HashProvider
}

func (u *UserData) Hash() key.HashProvider {
	if u.TxHash == nil {
		return u.PrivateKey
	} else {
		return u.TxHash
	}
}

//func (u *UserData) SetSignHandle(s sign.SignProvider) {
//	u.sign = s
//}

func (u *UserData) Sign(digest []byte) (signature []byte, err error) {
	hash := u.PrivateKey.Hash(digest)
	return u.PrivateKey.Sign(hash)
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
