package fisco_bcos

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/wonderivan/logger"
	"math/big"
	"strconv"
)

func NewFiscoBcosClient(config *config.Config) (*FiscoBcosClient, error) {

	fiscoBcosClient := &FiscoBcosClient{
		Client: client.Client{
			Config: config,
		},
	}

	err := fiscoBcosClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed")
		return nil, err
	}

	return fiscoBcosClient, nil
}

type FiscoBcosClient struct {
	client.Client
}

func (f *FiscoBcosClient) isSM() bool {
	return f.Config.GetAppInfo().AlgorithmType == enum.AppAlgorithmType_SM2
}

func (c *FiscoBcosClient) getBlockLimit() (*big.Int, error) {
	res, err := c.GetBlockHeight()
	if err != nil {
		return nil, err
	}
	if res.Header.Code != 0 {
		return nil, errors.New(res.Header.Msg)
	}

	height, err := strconv.ParseInt(res.Body.Data, 10, 64)
	if err != nil {
		return nil, errors.New("ledger height has error")
	}

	height = height + 100
	return new(big.Int).SetInt64(height), nil
}

func (c *FiscoBcosClient) getGroupId() (*big.Int, error) {
	groupId, err := strconv.ParseInt(c.Config.GetAppInfo().ChannelId, 10, 64)

	if err != nil {
		return nil, err
	}
	return new(big.Int).SetInt64(groupId), nil
}
