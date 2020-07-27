/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  remove
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:17
 */

package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RemoveEventResData struct {
	base.BaseResModel
	Body interface{} `json:"body"`
}

func (f *RemoveEventResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
