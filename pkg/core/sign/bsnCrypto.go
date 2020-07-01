package sign

import "encoding/base64"

type Crypto interface {
	Sign(value string) (string, error)
	Verify(mac, value string) bool
}

func NewCrypto(sign SignHandle) Crypto {

	return &bsnCrypto{
		sign: sign,
	}

}

type bsnCrypto struct {
	sign SignHandle
}

func (g *bsnCrypto) Sign(value string) (string, error) {

	digest, _ := g.sign.Hash([]byte(value))

	sign, err := g.sign.Sign(digest)
	if err != nil {
		return "", err
	}

	mac := base64.StdEncoding.EncodeToString(sign)

	return mac, nil
}

func (g *bsnCrypto) Verify(mac, value string) bool {

	sign, err := base64.StdEncoding.DecodeString(mac)
	if err != nil {
		return false
	}

	digest, _ := g.sign.Hash([]byte(value))

	falg, err := g.sign.Verify(sign, digest)

	if err != nil {
		return false
	}

	return falg
}
