package account

import (
	"crypto/ecdsa"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/sm"
	"github.com/xuperchain/crypto/gm/account"
)

//根据公钥生成address
func GetAddressFromPublicKey(pubkey string) (string, error) {
	key, err := sm.ConvertSMPublicKey(pubkey)
	if err != nil {
		return "", err
	}

	address, err := account.GetAddressFromPublicKey(key)
	return address, err
}

func GetEcdsaPrivateKey(prikey string) (*ecdsa.PrivateKey, error) {
	return sm.ConvertSMPrivateKey(prikey)
}

//获取json格式的公钥信息
func GetEcdsaPublicKeyJsonFormatFromPublicKey(pubkey string) (string, error) {
	key, err := sm.ConvertSMPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	publicKeyJsonStr, err := account.GetEcdsaPublicKeyJsonFormatFromPublicKey(key)
	return publicKeyJsonStr, err
}
