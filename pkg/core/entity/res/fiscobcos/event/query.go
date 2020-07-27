/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:15
 */

package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type QueryEventResData struct {
	base.BaseResModel
	Body EventList `json:"body"`
}

func (f *QueryEventResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.GetEncryptionValue()

}
