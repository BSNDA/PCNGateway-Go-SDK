package config

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/app"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req"
	"path"
)

const (
	_KeyStore = "keystore"
)

// Create a profile information
// api: address of the node gateway
// userCode: user's code
// appCode: DApp code
// puk : public key cert of node gateway
// prk : private key of DApp cert
// cert:https cert path
func NewConfig(api, userCode, appCode, puk, prk, mspDir, cert string) (*Config, error) {

	config := &Config{
		nodeApi:  api,
		mspDir:   mspDir,
		httpCert: cert,
		appCert:  certInfo{AppPublicCert: puk, UserAppPrivateCert: prk},
		user:     userInfo{UserCode: userCode},
		app:      appInfo{AppCode: appCode},
	}
	err := config.Init()
	return config, err
}

//Create a profile information
// api: address of the node gateway
// userCode: user's code
// appCode: DApp code
// prk : private key of DApp cert
func NewConfig2(api, userCode, appCode, prk, mspDir string) (*Config, error) {

	config := &Config{
		nodeApi:  api,
		mspDir:   mspDir,
		httpCert: "",
		appCert:  certInfo{AppPublicCert: "", UserAppPrivateCert: prk},
		user:     userInfo{UserCode: userCode},
		app:      appInfo{AppCode: appCode},
	}
	err := config.Init()
	return config, err
}

type Config struct {
	nodeApi string
	mspDir  string

	user userInfo
	app  appInfo

	//DApp cert【public key of bsn node gateway and private key of user's DApp】
	appCert certInfo

	//https connection cert
	httpCert string

	isInit bool
}

func (c *Config) GetAppInfo() appInfo {
	return c.app
}

func (c *Config) GetCert() string {
	return c.httpCert
}

func (c *Config) GetAppCert() certInfo {
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
		res, err := app.GetAppInfo(&reqData, c.nodeApi, c.httpCert)

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
			c.appCert.AppPublicCert = getGatewayPublicKey(c.app.AlgorithmType)
		}

		if c.appCert.AppPublicCert == "" {
			return errors.New("gateway public key not setting")
		}

		c.app.MspId = res.Body.MspId

		c.app.ChannelId = res.Body.ChannelId
		c.app.Version = res.Body.Version
		c.isInit = true
	}

	return nil
}

type certInfo struct {
	//public key cert of DApp
	AppPublicCert string

	autoPublicKey bool

	//Private key cert of user
	UserAppPrivateCert string
}

type appInfo struct {
	AppCode string
	AppType string

	CAType        enum.App_CaType
	AlgorithmType enum.App_AlgorithmType

	//AppCertPuk string

	MspId     string
	ChannelId string
	Version   string
}

type userInfo struct {
	UserCode string
}

type orgConfig struct {
	OrgCode string
	MspId   string

	NodeApi string
}
