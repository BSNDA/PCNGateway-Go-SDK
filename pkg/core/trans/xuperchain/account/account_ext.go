package account

import (
	"crypto/ecdsa"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/account"
)

//根据公钥生成address
func GetAddressFromPublicKey(pubkey string) (string, error) {
	puk, err := sm2.ReadPublicKeyFromMem([]byte(pubkey), nil)
	if err != nil {
		return "", err
	}
	var key = &ecdsa.PublicKey{}
	key.Y = puk.Y
	key.X = puk.X
	key.Curve = puk.Curve
	address, err := account.GetAddressFromPublicKey(key)
	return address, err
}

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

//获取json格式的公钥信息
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
