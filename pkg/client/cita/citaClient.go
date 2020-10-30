package cita

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/wonderivan/logger"
	"math/big"
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

//func (c *CitaClient) getBlockLimit() (*big.Int, error) {
//	res, err := c.GetBlockHeight()
//	if err != nil {
//		return nil, err
//	}
//	if res.Header.Code != 0 {
//		return nil, errors.New(res.Header.Msg)
//	}
//
//	height, err := strconv.ParseInt(res.Body.Data, 10, 64)
//	if err != nil {
//		return nil, errors.New("ledger height has error")
//	}
//
//	height = height + 100
//	return new(big.Int).SetInt64(height), nil
//}

func (c *CitaClient) getGroupId() (*big.Int, error) {
	groupId, err := strconv.ParseInt(c.Config.GetAppInfo().ChannelId, 10, 64)

	if err != nil {
		return nil, err
	}
	return new(big.Int).SetInt64(groupId), nil
}
