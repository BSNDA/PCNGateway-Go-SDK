package keystore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/secp256r1"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/sm"
	"github.com/BSNDA/bsn-sdk-crypto/keystore/cert"
	"github.com/BSNDA/bsn-sdk-crypto/keystore/key"
	"github.com/wonderivan/logger"

	ksecdsa "github.com/BSNDA/bsn-sdk-crypto/keystore/ecdsa"
	kssm "github.com/BSNDA/bsn-sdk-crypto/keystore/sm"
)

func LoadKey(user *msp.UserData, ks key.KeyStore, algorithmType enum.App_AlgorithmType) error {

	key, err := cert.ImportCert(user.EnrollmentCertificate)
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
		pk := ksecdsa.GetECDSAPrivateKey(prikey)
		if pk != nil {
			user.PrivateKey = pk
			user.SetSignHandle(secp256r1.NewTransUserR1Handle(pk))
			return nil
		}
	}

	if algorithmType == enum.AppAlgorithmType_SM2 {
		pk := kssm.GetSmPrivateKey(prikey)
		if pk != nil {
			user.SetSignHandle(sm.NewTransUserSMHandle(pk))
			return nil
		}
	}

	return errors.New("algorithm type failed")
}
