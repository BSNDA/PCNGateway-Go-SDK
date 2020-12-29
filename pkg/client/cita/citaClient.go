package cita

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/wonderivan/logger"
	"strconv"
)

func NewCitaClient(config *config.Config) (*CitaClient, error) {
	client := client.Client{
		Config: config,
	}

	citaClient := &CitaClient{
		client,
	}

	err := citaClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed")
		return nil, err
	}

	return citaClient, nil
}

type CitaClient struct {
	client.Client
}

func (f *CitaClient) isSM() bool {
	return f.Config.GetAppInfo().AlgorithmType == enum.AppAlgorithmType_SM2
}

func (c *CitaClient) getBlockLimit() (uint64, error) {
	res, err := c.GetBlockHeight()
	if err != nil {
		return 0, err
	}
	if res.Header.Code != 0 {
		return 0, errors.New(res.Header.Msg)
	}

	height, err := strconv.ParseUint(res.Body.Data, 10, 64)
	if err != nil {
		return 0, errors.New("ledger height has error")
	}

	height = height + 80
	return height, nil
}

func (c *CitaClient) getBaseInfo() (string, uint32, error) {

	chainId := c.Config.GetAppInfo().ChannelId
	if !has0xPrefix(chainId) {
		chainId = "0x" + chainId
	}

	//chainId:="0x675"
	version, err := strconv.ParseUint(c.Config.GetAppInfo().Version, 10, 64)
	if err != nil {
		return "", 0, err
	}
	return chainId, uint32(version), nil
}
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
