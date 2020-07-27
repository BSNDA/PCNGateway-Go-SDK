/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  ledgerReqData
 * @Version: 1.0.0
 * @Date: 2020/6/9 14:31
 */

package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type LedgerReqData struct {
	base.BaseReqModel
	Body LedgerReqDataBody `json:"body"`
}

type LedgerReqDataBody struct {
}

func (f *LedgerReqData) GetEncryptionValue() string {

	return f.GetBaseEncryptionValue()

}
