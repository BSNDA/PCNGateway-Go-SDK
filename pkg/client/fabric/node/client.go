package node

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	crySign "github.com/BSNDA/bsn-sdk-crypto/sign"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/sign"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
)

const (
	fabricApiVersion = "v1"
	fabricApiFormat  = "%s/api/fabric/%s/%s"
)

func NewNodeCli(url string, algorithmType enum.App_AlgorithmType, pri string) (*NodeClient, error) {

	cli := &NodeClient{gateWayURL: url}
	err := cli.SetAlgorithm(algorithmType, pri)
	if err != nil {
		return nil, err
	}
	return cli, nil

}

type NodeClient struct {
	gateWayURL string
	sign       sign.Crypto
	mspId      string
}

func (n *NodeClient) URL() string {

	return n.gateWayURL
}

func (c *NodeClient) SetAlgorithm(algorithmType enum.App_AlgorithmType, pri string) error {

	kt := algorithmType.ToKeyType()

	privKey, err := key.NewPrivateKeyProvider(kt, pri)
	if err != nil {
		return errors.WithMessagef(err, "new [%s] private key provider key has error", kt.String())
	}

	puk := config.GetGatewayPublicKey(algorithmType)

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

func (c *NodeClient) Sign(data string) (string, error) {
	mac, err := c.sign.Sign(data)
	if err != nil {
		return "", errors.WithMessage(err, "Exception in signature")
	}
	return mac, nil
}

func (c *NodeClient) Verify(mac, data string) bool {
	if mac == "" {
		return true
	}
	return c.sign.Verify(mac, data)

}

func (c *NodeClient) methodUrl(method string) string {
	return fmt.Sprintf(fabricApiFormat, c.gateWayURL, fabricApiVersion, method)
}

func (c *NodeClient) Call(method string, req base.ReqInterface, res base.ResInterface) error {
	url := c.methodUrl(method)
	mac, err := c.Sign(req.GetEncryptionValue())
	if err != nil {
		return err
	}
	req.SetMac(mac)

	reqBytes, err := json.Marshal(req)
	if err != nil {
		logger.Error("request parameter serialization failed：", err)
		return errors.WithMessage(err, "request parameter serialization failed")
	}

	resBytes, err := http.SendPost(reqBytes, url)
	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return errors.WithMessage(err, "send post has error")
	}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return errors.WithMessage(err, "return parameter serialization failed")
	}

	//v :=c.Verify(res.GetMac(),res.GetEncryptionValue())

	//fmt.Println(v)

	return nil
}
