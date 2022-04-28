package client

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/sign"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	crySign "github.com/BSNDA/bsn-sdk-crypto/sign"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
)

type Client struct {
	Config *config.Config
	sign   sign.Crypto
}

func (c *Client) SetAlgorithm(algorithmType enum.App_AlgorithmType, puk, pri string) error {

	kt := algorithmType.ToKeyType()

	privKey, err := key.NewPrivateKeyProvider(kt, pri)
	if err != nil {
		return errors.WithMessagef(err, "new [%s] private key provider key has error", kt.String())
	}

	pubKey, err := key.NewPublicProvider(kt, puk)
	if err != nil {
		return errors.WithMessagef(err, "new [%s] public key provider key has error", kt.String())
	}

	sh, err := crySign.NewSignProvider(crySign.WithPrivateKey(privKey), crySign.WithPublicKey(pubKey))
	if err != nil {
		return errors.WithMessagef(err, "new [%s] sign provider has error", kt.String())
	}

	c.sign = sign.NewCrypto(sh)
	return nil

}

func (c *Client) GetHeader() base.ReqHeader {
	return c.Config.GetReqHeader()
}

func (c *Client) GetURL(url string) string {
	return c.Config.GetNodeApi() + url
}

func (c *Client) Sign(data string) string {

	mac, err := c.sign.Sign(data)

	if err != nil {
		logger.Error("Exception in signature : %s", err.Error())
	}

	return mac
}

func (c *Client) Verify(mac, data string) bool {
	if mac == "" {
		return true
	}
	return c.sign.Verify(mac, data)

}
