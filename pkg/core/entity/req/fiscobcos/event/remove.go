/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  remove
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:14
 */

package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RemoveReqData struct {
	base.BaseReqModel
	Body RemoveReqDataBody `json:"body"`
}

type RemoveReqDataBody struct {
	EventId string `json:"eventId"`
}

func (f *RemoveReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.EventId

}
