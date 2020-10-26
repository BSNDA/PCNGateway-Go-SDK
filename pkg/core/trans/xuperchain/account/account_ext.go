// @Title  account_ext
// @Description
// @Author  zxl  2020/7/1 16:39
// @Version 1.0.0
// @Update  2020/7/1 16:39
package account

import (
	"crypto/ecdsa"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/account"
)

func GetEcdsaPrivateKey(prikey string) (*ecdsa.PrivateKey, error) {
	pkey, err := sm2.ReadPrivateKeyFromMem([]byte(prikey), nil)
	if err != nil {
		return nil, err
	}
	var key = &ecdsa.PrivateKey{}
	key.Y = pkey.Y
	key.X = pkey.X
	key.Curve = pkey.Curve
	key.D = pkey.D
	return key, nil
}

func GetEcdsaPublicKeyJsonFormatFromPublicKey(pubkey string) (string, error) {
	puk, err := sm2.ReadPublicKeyFromMem([]byte(pubkey), nil)
	if err != nil {
		return "", err
	}
	var key = &ecdsa.PublicKey{}
	key.Y = puk.Y
	key.X = puk.X
	key.Curve = puk.Curve
	publicKeyJsonStr, err := account.GetEcdsaPublicKeyJsonFormatFromPublicKey(key)
	return publicKeyJsonStr, err
}
