package keystore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"github.com/wonderivan/logger"
)

func LoadKey(user *msp.UserData, ks bccsp.KeyStore) error {

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
	pk := GetECDSAPrivateKey(prikey)
	user.PrivateKey = pk

	return nil
}
