package config

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/app"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"path"
)

const (
	_KeyStore = "keystore"
)

//Create a profile information
// api: address of the node gateway
// userCode: user's code
// appCode: DApp code
// prk : private key of DApp cert
func NewConfig(api, userCode, appCode, prk, mspDir string) (*Config, error) {

	config := &Config{
		nodeApi: api,
		mspDir:  mspDir,
		appCert: CertInfo{AppPublicCert: "", UserAppPrivateCert: prk},
		user:    userInfo{UserCode: userCode},
		app:     AppInfo{AppCode: appCode},
	}
	err := config.Init()
	return config, err
}

type Config struct {
	nodeApi string
	mspDir  string

	user userInfo
	app  AppInfo

	//DApp cert【public key of bsn node gateway and private key of user's DApp】
	appCert CertInfo

	isInit bool
}

func (c *Config) GetAppInfo() AppInfo {
	return c.app
}

func (c *Config) GetUserCode() string {
	return c.user.UserCode
}

func (c *Config) GetAppCert() CertInfo {
	return c.appCert
}

func (c *Config) GetKSPath() string {
	return path.Join(c.mspDir, _KeyStore)
}

func (c *Config) GetUSPath() string {
	return c.mspDir
}

func (c *Config) GetNodeApi() string {
	return c.nodeApi
}

func (c *Config) GetReqHeader() base.ReqHeader {
	header := base.ReqHeader{
		UserCode: c.user.UserCode,
		AppCode:  c.app.AppCode,
	}

	return header
}

func (c *Config) Init() error {
	if !c.isInit {
		reqData := req.AppInfoReqData{}

		reqData.Header = c.GetReqHeader()

		//reqData.Body = req.AppInfoReqDataBody{}
		res, err := app.GetAppInfo(&reqData, c.nodeApi, "")

		if err != nil {
			return err
		}

		if res.Header.Code != 0 {
			return errors.New("get app info failed ：" + res.Header.Msg)
		}

		c.app.AppType = res.Body.AppType

		c.app.CAType = enum.App_CaType(res.Body.CaType)
		c.app.AlgorithmType = enum.App_AlgorithmType(res.Body.AlgorithmType)

		if c.appCert.AppPublicCert == "" {
			c.appCert.AppPublicCert = GetGatewayPublicKey(c.app.AlgorithmType)
		}

		if c.appCert.AppPublicCert == "" {
			return errors.New("gateway public key not setting")
		}

		c.app.MspId = res.Body.MspId

		c.app.ChannelId = res.Body.ChannelId
		c.app.Version = res.Body.FabricVersion
		c.isInit = true
	}

	return nil
}

type CertInfo struct {
	//public key cert of DApp
	AppPublicCert string

	autoPublicKey bool

	//Private key cert of user
	UserAppPrivateCert string
}

type AppInfo struct {
	AppCode string
	AppType string

	CAType        enum.App_CaType
	AlgorithmType enum.App_AlgorithmType

	//AppCertPuk string

	MspId     string
	ChannelId string
	Version   string
}

func (a *AppInfo) TxHash() key.HashProvider {
	if a.AlgorithmType == enum.AppAlgorithmType_SM2 && a.Version == "2.2.1" {
		return &key.SM3Hash{}
	} else {
		return &key.SHA256Hash{}
	}
}

type userInfo struct {
	UserCode string
}
