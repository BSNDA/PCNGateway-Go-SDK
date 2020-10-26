package xuperchain

import (
	"bytes"
	"encoding/json"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain/account"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain/common"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain/crypto"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/trans/xuperchain/pb"
	"time"
)

//GenerateInvokeIR
func GenerateInvokeIR(contractName, methodName string, args map[string][]byte) *pb.InvokeRequest {
	return &pb.InvokeRequest{
		ModuleName:   "wasm",
		MethodName:   methodName,
		ContractName: contractName,
		Args:         args,
	}
}

//GenerateTransaction
func GenerateTransaction(initiator string, response *pb.InvokeResponse, privatekey, publickey string) (*pb.Transaction, error) {
	var err error
	tx := &pb.Transaction{
		Desc:      []byte("sdk contract transaction"),
		Version:   common.TxVersion,
		Coinbase:  false,
		Timestamp: time.Now().UnixNano(),
		Initiator: initiator,
	}
	tx.TxInputsExt = response.GetInputs()
	tx.TxOutputsExt = response.GetOutputs()
	tx.ContractRequests = response.GetRequests()

	err = common.SetSeed()
	if err != nil {
		return nil, err
	}
	tx.Nonce = common.GetNonce()

	digestHash, err := makeTxDigestHash(tx)
	if err != nil {
		return nil, err
	}
	pKey, err := account.GetEcdsaPrivateKey(privatekey)
	if err != nil {
		return nil, err
	}
	cryptoClient := crypto.GetGmCryptoClient()
	sign, err := cryptoClient.SignECDSA(pKey, digestHash)
	if err != nil {
		return nil, err
	}
	pukJsonStr, err := account.GetEcdsaPublicKeyJsonFormatFromPublicKey(publickey)
	if err != nil {
		return nil, err
	}
	signatureInfo := &pb.SignatureInfo{
		PublicKey: pukJsonStr,
		Sign:      sign,
	}
	var signatureInfos []*pb.SignatureInfo
	signatureInfos = append(signatureInfos, signatureInfo)
	tx.InitiatorSigns = signatureInfos
	tx.Txid, _ = makeTransactionID(tx)
	return tx, nil
}

// makeTxDigestHash
func makeTxDigestHash(tx *pb.Transaction) ([]byte, error) {
	coreData, err := encodeTxData(tx, false)
	if err != nil {
		return nil, err
	}

	cryptoClient := crypto.GetXchainCryptoClient()
	txHash := cryptoClient.HashUsingDoubleSha256(coreData)

	return txHash, nil
}

// makeTransactionID
func makeTransactionID(tx *pb.Transaction) ([]byte, error) {
	coreData, err := encodeTxData(tx, true)
	if err != nil {
		return nil, err
	}

	cryptoClient := crypto.GetXchainCryptoClient()
	txid := cryptoClient.HashUsingDoubleSha256(coreData)

	return txid, nil
}

// encodeTxData
func encodeTxData(tx *pb.Transaction, includeSigns bool) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	for _, txInput := range tx.TxInputs {
		if len(txInput.RefTxid) > 0 {
			err := encoder.Encode(txInput.RefTxid)
			if err != nil {
				return nil, err
			}
		}
		err := encoder.Encode(txInput.RefOffset)
		if err != nil {
			return nil, err
		}
		if len(txInput.FromAddr) > 0 {
			err = encoder.Encode(txInput.FromAddr)
			if err != nil {
				return nil, err
			}
		}
		if len(txInput.Amount) > 0 {
			err = encoder.Encode(txInput.Amount)
			if err != nil {
				return nil, err
			}
		}
		err = encoder.Encode(txInput.FrozenHeight)
		if err != nil {
			return nil, err
		}
	}
	err := encoder.Encode(tx.TxOutputs)
	if err != nil {
		return nil, err
	}
	if len(tx.Desc) > 0 {
		err = encoder.Encode(tx.Desc)
		if err != nil {
			return nil, err
		}
	}
	err = encoder.Encode(tx.Nonce)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(tx.Timestamp)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(tx.Version)
	if err != nil {
		return nil, err
	}
	for _, txInputExt := range tx.TxInputsExt {
		if err = encoder.Encode(txInputExt.Bucket); err != nil {
			return nil, err
		}
		if len(txInputExt.Key) > 0 {
			if err = encoder.Encode(txInputExt.Key); err != nil {
				return nil, err
			}
		}
		if len(txInputExt.RefTxid) > 0 {
			if err = encoder.Encode(txInputExt.RefTxid); err != nil {
				return nil, err
			}
		}
		if err = encoder.Encode(txInputExt.RefOffset); err != nil {
			return nil, err
		}
	}
	for _, txOutputExt := range tx.TxOutputsExt {
		if err = encoder.Encode(txOutputExt.Bucket); err != nil {
			return nil, err
		}
		if len(txOutputExt.Key) > 0 {
			if err = encoder.Encode(txOutputExt.Key); err != nil {
				return nil, err
			}
		}
		if len(txOutputExt.Value) > 0 {
			if err = encoder.Encode(txOutputExt.Value); err != nil {
				return nil, err
			}
		}
	}
	if err = encoder.Encode(tx.ContractRequests); err != nil {
		return nil, err
	}
	if err = encoder.Encode(tx.Initiator); err != nil {
		return nil, err
	}
	if err = encoder.Encode(tx.AuthRequire); err != nil {
		return nil, err
	}
	if includeSigns {
		if err = encoder.Encode(tx.InitiatorSigns); err != nil {
			return nil, err
		}
		if err = encoder.Encode(tx.AuthRequireSigns); err != nil {
			return nil, err
		}
		if tx.GetXuperSign() != nil {
			err = encoder.Encode(tx.AuthRequireSigns)
			if err != nil {
				return nil, err
			}
		}
	}
	if err = encoder.Encode(tx.Coinbase); err != nil {
		return nil, err
	}
	if err = encoder.Encode(tx.Autogen); err != nil {
		return nil, err
	}

	if tx.Version >= 2 {
		if err = encoder.Encode(tx.HDInfo); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
