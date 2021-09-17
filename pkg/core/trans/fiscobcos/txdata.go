package fiscobcos

import (
	"bytes"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/eth"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/sm"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/sha3"

	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/tjfoc/gmsm/sm2"
	"io"
	"math/big"
	"sync/atomic"
)

type Transaction struct {
	data txdata

	//sign txsign

	// caches
	hash     atomic.Value
	size     atomic.Value
	from     atomic.Value
	smcrypto bool
}

type txdata struct {
	AccountNonce *big.Int        `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice"   gencodec:"required"`
	GasLimit     *big.Int        `json:"gas"        gencodec:"required"`
	BlockLimit   *big.Int        `json:"blocklimit" gencodec:"required"`
	Recipient    *common.Address `json:"to"         rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"      gencodec:"required"`
	Payload      []byte          `json:"input"      gencodec:"required"`
	// for fisco bcos 2.0
	ChainID   *big.Int `json:"chainId"    gencodec:"required"`
	GroupID   *big.Int `json:"groupId"    gencodec:"required"`
	ExtraData []byte   `json:"extraData"  gencodec:"required"` // rlp:"nil"

	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

// NewTransaction returns a new transaction
func NewTransaction(nonce *big.Int, to common.Address, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, chainId *big.Int, groupId *big.Int, extraData []byte, smcrypto bool) *Transaction {
	return newTransaction(nonce, &to, amount, gasLimit, gasPrice, blockLimit, data, chainId, groupId, extraData, smcrypto)
}
func newTransaction(nonce *big.Int, to *common.Address, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, blockLimit *big.Int, data []byte, chainId *big.Int, groupId *big.Int, extraData []byte, smcrypto bool) *Transaction {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := txdata{
		AccountNonce: nonce,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		GasLimit:     gasLimit,
		BlockLimit:   blockLimit,
		Price:        new(big.Int),
		ChainID:      new(big.Int),
		GroupID:      new(big.Int),
		ExtraData:    extraData,

		V: new(big.Int),
		R: new(big.Int),
		S: new(big.Int),
	}
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}
	if chainId != nil {
		d.ChainID.Set(chainId)
	}
	if groupId != nil {
		d.GroupID.Set(groupId)
	}
	if extraData != nil {
		d.ExtraData = extraData
	}

	return &Transaction{data: d, smcrypto: smcrypto}
}

// EncodeRLP implements rlp.Encoder
func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}

// DecodeRLP implements rlp.Decoder
func (tx *Transaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.data)
	if err == nil {
		tx.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (tx *Transaction) Sign(priKey *ecdsa.PrivateKey, isSM bool) []byte {
	txb, _ := rlp.EncodeToBytes(tx.data)

	if isSM {

	} else {

		hash, _ := eth.Hash(txb)

		r, s, v, _, _ := eth.SignData(priKey, hash)

		tx.data.V = v
		tx.data.R = r
		tx.data.S = s
	}

	txd, _ := rlp.EncodeToBytes(tx.data)
	return txd

}

func (tx *Transaction) SignData(priKey interface{}) ([]byte, error) {
	//txb, _ := rlp.EncodeToBytes(tx.data)

	if tx.smcrypto {

		txb := SM3HashNonSig(tx).Bytes()
		pk := priKey.(*sm2.PrivateKey)
		r, s, pub, err := sm.SignData(pk, txb)
		if err != nil {
			return nil, err
		}

		tx.data.V = pub
		tx.data.R = r
		tx.data.S = s
	} else {
		txb := Hash(tx).Bytes()
		//hash, _ := eth.Hash(txb)
		r, s, v, _, err := eth.SignData(priKey.(*ecdsa.PrivateKey), txb)

		if err != nil {
			return nil, err
		}

		tx.data.V = v
		tx.data.R = r
		tx.data.S = s
	}

	txd, _ := rlp.EncodeToBytes(tx.data)
	return txd, nil

}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func Hash(tx *Transaction) common.Hash {
	return rlpHash([]interface{}{
		tx.data.AccountNonce,
		tx.data.Price,
		tx.data.GasLimit,
		tx.data.BlockLimit,
		tx.data.Recipient,
		tx.data.Amount,
		tx.data.Payload,
		tx.data.ChainID,
		tx.data.GroupID,
		tx.data.ExtraData,
	})
}

func SM3HashNonSig(tx *Transaction) (h common.Hash) {
	var src []byte
	buf := bytes.NewBuffer(src)
	rlp.Encode(buf, []interface{}{
		tx.data.AccountNonce,
		tx.data.Price,
		tx.data.GasLimit,
		tx.data.BlockLimit,
		tx.data.Recipient,
		tx.data.Amount,
		tx.data.Payload,
		tx.data.ChainID,
		tx.data.GroupID,
		tx.data.ExtraData,
	})

	h3 := sm3.New()
	h3.Write(buf.Bytes())
	hash := h3.Sum(nil)
	copy(h[:], hash)
	return
}
