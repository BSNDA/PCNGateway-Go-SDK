/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  transReqDataBody
 * @Version: 1.0.0
 * @Date: 2020/6/4 17:37
 */

package node

import "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

type TransReqData struct {
	base.BaseReqModel
	Body TransReqDataBody `json:"body"`
}

type TransReqDataBody struct {
	UserId       string `json:"userId"`
	ContractName string `json:"contractName"`
	FuncName     string `json:"funcName"`
	FuncParam    string `json:"funcParam"`
}

func (f *TransReqData) GetEncryptionValue() string {

	fp := f.Body.FuncParam

	return f.GetBaseEncryptionValue() + f.Body.UserId + f.Body.ContractName + f.Body.FuncName + fp

}
