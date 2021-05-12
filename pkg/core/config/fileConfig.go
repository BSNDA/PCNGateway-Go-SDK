package config

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/file"
)

const (
	_DefaultConfigPath = "./conf/config.json"
	_DefaultMSPPath    = "./msp"
)

type fileConfig struct {
	NodeApi        string `json:"nodeApi"`
	UserCode       string `json:"userCode"`
	AppCode        string `json:"appCode"`
	UserPrivateKey string `json:"userPrivateKey"`
	BSNPublicKey   string `json:"bsnPublicKey"`
	MspPath        string `json:"mspPath"`
}

// NewConfigFormFile get config from file, if path is empty ,use the ./conf/config.json file in the root directory
func NewConfigFormFile(path string) (*Config, error) {
	if path == "" {
		path = _DefaultConfigPath
	}

	fileBytes, err := file.ReadFile(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read config file faild:%s", err.Error()))
	}

	fileConf := &fileConfig{}
	err = json.Unmarshal(fileBytes, fileConf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("read config file faild:%s", err.Error()))
	}

	if fileConf.MspPath == "" {
		fileConf.MspPath = _DefaultMSPPath
	}

	return NewConfig2(fileConf.NodeApi,
		fileConf.UserCode,
		fileConf.AppCode,
		fileConf.UserPrivateKey,
		fileConf.MspPath)
}
