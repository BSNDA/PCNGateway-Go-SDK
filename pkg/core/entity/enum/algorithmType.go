package enum

import "github.com/BSNDA/bsn-sdk-crypto/types"

type App_AlgorithmType int

func (a App_AlgorithmType) ToKeyType() types.KeyType {

	if a == AppAlgorithmType_SM2 {
		return types.SM
	} else if a == AppAlgorithmType_K1 {
		return types.ECDSA_K1
	} else {
		return types.ECDSA_R1
	}

}

const (
	AppAlgorithmType_Not App_AlgorithmType = 0

	AppAlgorithmType_SM2 App_AlgorithmType = 1
	AppAlgorithmType_R1  App_AlgorithmType = 2
	AppAlgorithmType_K1  App_AlgorithmType = 3
)
