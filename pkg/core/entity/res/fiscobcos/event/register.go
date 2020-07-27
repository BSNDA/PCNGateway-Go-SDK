/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  register
 * @Version: 1.0.0
 * @Date: 2020/6/9 15:17
 */

package event

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type RegisterEventResData struct {
	base.BaseResModel
	Body RegisterEvent `json:"body"`
}

type RegisterEvent struct {
	EventId string `json:"eventId"`
}

func (f *RegisterEventResData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue() + f.Body.EventId

}
