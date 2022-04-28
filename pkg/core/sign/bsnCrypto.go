package sign

import (
	"encoding/base64"
	"github.com/BSNDA/bsn-sdk-crypto/sign"
)

type Crypto interface {
	Sign(value string) (string, error)
	Verify(mac, value string) bool
}

func NewCrypto(sign sign.SignProvider) Crypto {

	return &gatewayCrypto{
		sign: sign,
	}

}

type gatewayCrypto struct {
	sign sign.SignProvider
}

func (g *gatewayCrypto) Sign(value string) (string, error) {

	digest := g.sign.Hash([]byte(value))

	signBytes, err := g.sign.Sign(digest)
	if err != nil {
		return "", err
	}

	mac := base64.StdEncoding.EncodeToString(signBytes)

	return mac, nil
}

func (g *gatewayCrypto) Verify(mac, value string) bool {

	signBytes, err := base64.StdEncoding.DecodeString(mac)
	if err != nil {
		return false
	}

	digest := g.sign.Hash([]byte(value))

	falg, err := g.sign.Verify(signBytes, digest)

	if err != nil {
		return false
	}

	return falg
}
