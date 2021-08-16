package client

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/sign"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/secp256k1"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/secp256r1"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/sm"
	"github.com/wonderivan/logger"
)

type Client struct {
	Config *config.Config
	sign   sign.Crypto
}

func (c *Client) SetAlgorithm(algorithmType enum.App_AlgorithmType, puk, pri string) error {

	var sh sign.SignHandle
	var err error

	switch algorithmType {
	case enum.AppAlgorithmType_SM2:
		sh, err = sm.NewSM2Handle(puk, pri)
	case enum.AppAlgorithmType_R1:
		sh, err = secp256r1.NewEcdsaR1Handle(puk, pri)
	case enum.AppAlgorithmType_K1:
		sh, err = secp256k1.NewEcdsaK1Handle(puk, pri)
	default:
		return errors.New("Invalid certificate type")
	}

	if err != nil {
		return err
	} else {
		c.sign = sign.NewCrypto(sh)
		return nil
	}

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
		logger.Error("Exception in signature")
	}

	return mac
}

func (c *Client) Verify(mac, data string) bool {
	if mac == "" {
		return true
	}
	return c.sign.Verify(mac, data)

}
