/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:09
 */

package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type QueryReqData struct {
	base.BaseReqModel
	Body interface{} `json:"body"`
}

func (f *QueryReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
