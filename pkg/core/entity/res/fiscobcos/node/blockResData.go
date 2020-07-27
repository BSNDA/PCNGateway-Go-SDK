/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  blockResData
 * @Version: 1.0.0
 * @Date: 2020/6/9 14:21
 */

package node

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"strconv"
)

type BlockResData struct {
	base.BaseResModel
	Body *BlockData `json:"body"`
}

type BlockData struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber int64  `json:"blockNumber"`

	ParentBlockHash string `json:"parentBlockHash"`

	BlockSize int64 `json:"blockSize"`

	BlockTime int64 `json:"blockTime"`

	Author string `json:"author"`

	Transactions []*TransactionData `json:"transactions"`
}

func (f *BlockData) GetEncryptionValue() string {

	fb := ""
	fb = fb + f.BlockHash
	fb = fb + strconv.FormatInt(f.BlockNumber, 10)
	fb = fb + f.ParentBlockHash
	fb = fb + strconv.FormatInt(f.BlockSize, 10)
	fb = fb + strconv.FormatInt(f.BlockTime, 10)
	fb = fb + f.Author

	for _, tx := range f.Transactions {
		fb = fb + tx.GetEncryptionValue()
	}

	return fb
}

func (f *BlockResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
