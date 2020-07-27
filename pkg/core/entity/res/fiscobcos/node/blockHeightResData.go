/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  blockHeightResData
 * @Version: 1.0.0
 * @Date: 2020/6/9 14:32
 */

package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type BlockHeightResData struct {
	base.BaseResModel
	Body BlockHeightResDataBody `json:"body"`
}

type BlockHeightResDataBody struct {
	Data string `json:"data"`
}

func (f *BlockHeightResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.Data

}
