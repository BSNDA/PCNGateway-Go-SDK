/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  blockTxResData
 * @Version: 1.0.0
 * @Date: 2020/6/5 14:35
 */

package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockTxResData struct {
	base.BaseResModel
	Body TransactionData `json:"body"`
}

type TransactionData struct {
	TxId        string `json:"txId"`
	BlockHash   string `json:"blockHash"`
	BlockNumber uint64 `json:"blockNumber"`

	GasUsed int64 `json:"gasUsed"`

	From  string `json:"from"`
	To    string `json:"to"`
	Value int64  `json:"value"`

	Input string `json:"input"`
}

func (f *TransactionData) GetEncryptionValue() string {

	fb := ""

	fb = fb + f.TxId
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatUint(f.BlockNumber, 10)
	fb = fb + strconv.FormatInt(f.GasUsed, 10)
	fb = fb + f.From
	fb = fb + f.To
	fb = fb + strconv.FormatInt(f.Value, 10)
	fb = fb + f.Input

	return fb
}

func (f *BlockTxResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
