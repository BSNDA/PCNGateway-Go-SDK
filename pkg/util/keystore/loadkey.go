package keystore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto/secp256r1"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/crypto/sm"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"github.com/wonderivan/logger"
)

func LoadKey(user *msp.UserData, ks bccsp.KeyStore, algorithmType enum.App_AlgorithmType) error {

	key, err := ImportCert(user.EnrollmentCertificate)
	if err != nil {
		logger.Error("get cert has error:", err.Error())
		return err
	}

	prikey, err := ks.GetKey(key.SKI())
	if err != nil {
		logger.Error("get cert has error:，", err.Error())
		return err
	}

	if algorithmType == enum.AppAlgorithmType_R1 {
		pk := GetECDSAPrivateKey(prikey)
		if pk != nil {
			user.PrivateKey = pk
			user.SetSignHandle(secp256r1.NewTransUserR1Handle(pk))
			return nil
		}
	}

	if algorithmType == enum.AppAlgorithmType_SM2 {
		pk := GetSmPrivateKey(prikey)
		if pk != nil {
			user.SetSignHandle(sm.NewTransUserSMHandle(pk))
			return nil
		}
	}

	return errors.New("algorithm type failed")
}
