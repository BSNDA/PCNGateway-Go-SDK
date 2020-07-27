// @Title  xuperchainClient
// @Description
// @Author  zxl  2020/7/22 19:31
// @Version 1.0.0
// @Update  2020/7/22 19:31
package xuperchain

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/wonderivan/logger"
)

func NewXuperChainClient(config *config.Config) (*XuperChainClient, error) {
	client := client.Client{
		Config: config,
	}

	xuperChainClient := &XuperChainClient{
		client,
	}

	err := xuperChainClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed")
		return nil, err
	}

	return xuperChainClient, nil
}

type XuperChainClient struct {
	client.Client
}

func (f *XuperChainClient) isSM() bool {
	return f.Config.GetAppInfo().AlgorithmType == enum.AppAlgorithmType_SM2
}
