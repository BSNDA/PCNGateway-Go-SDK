package cita

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/cita/pb"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/golang/protobuf/proto"
	"math/big"
	"strconv"
)

func NewTransaction(nonce string, to common.Address, quota uint64, validUntilBlock uint64, value []byte, chainId string, data []byte, version uint32) *pb.Transaction {
	return newTransaction(nonce, to, quota, validUntilBlock, value, chainId, data, version)
}
func newTransaction(nonce string, to common.Address, quota uint64, validUntilBlock uint64, value []byte, chainId string, data []byte, version uint32) *pb.Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := pb.Transaction{
		Nonce:           nonce,
		Quota:           quota,
		ValidUntilBlock: validUntilBlock,
		Data:            data,
		Version:         version,
	}
	valuebyte, err := ToBytes32(value)
	if err != nil {
		return nil
	}
	d.Value = valuebyte

	if version == 0 {
		chain, _ := strconv.ParseUint(chainId, 10, 64)
		d.ChainId = uint32(chain)
		d.To = string(to[:])

	} else {

		bigchainId, _ := hexutil.DecodeBig(chainId)
		chainidbyte, err := ToBytes32(bigchainId.Bytes())
		if err != nil {
			return nil
		}
		d.ChainIdV1 = chainidbyte
		d.ToV1 = to.Bytes()
	}
	return &d
}

func ToBytes32(b []byte) ([]byte, error) {
	bl := len(b)
	if bl > 32 {
		return nil, errors.New("len out of limit")
	} else if bl == 32 {
		return b, nil
	} else {
		var a [32]byte
		for i := 0; i < bl; i++ {
			a[32-(bl-i)] = b[i]
		}
		var bb []byte
		for i := 0; i < len(a); i++ {
			bb = append(bb, a[i])
		}
		return bb, nil
	}
}

func SignData(tx *pb.Transaction, priKey key.PrivateKeyProvider, isSM bool) ([]byte, error) {

	txb, err := proto.Marshal(tx)
	if err != nil {
		return nil, err
	}
	if isSM {

		//pk := priKey.Key().(*sm2.PrivateKey)

		h := priKey.Hash(txb)

		r, s, _, err := priKey.SignTx(h)

		if err != nil {
			return nil, err
		}
		//publicKeyStr := fmt.Sprintf("%s%s", fillStr64(String16(pk.X)), fillStr64(String16(pk.Y)))
		sign := fmt.Sprintf("%s%s", fillStr64(String16(r)), fillStr64(String16(s)))
		signBytes, err := hex.DecodeString(sign)
		if err != nil {
			return nil, err
		}
		publicKeyBytes := priKey.PublicKey().Bytes()[1:]
		//if err != nil {
		//	return nil, err
		//}
		return BytesCombine(signBytes, publicKeyBytes), nil
	}

	return nil, nil

}
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}
func fillStr64(str string) string {
	if len(str) >= 64 {
		return str
	} else {
		s := fmt.Sprintf("%064s", str)
		return s
	}
}
func String16(x *big.Int) string {
	return x.Text(16)
}
func serializeUnverifiedTransaction(tx *pb.Transaction, sig []byte) (string, error) {
	trans := pb.UnverifiedTransaction{
		Signature:   sig,
		Transaction: tx,
		Crypto:      pb.Crypto_DEFAULT,
	}
	t, err := proto.Marshal(&trans)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(t), nil
}
