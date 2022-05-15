package fabric

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric/node"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
)

//initialize client requested by fabric
func InitFabricClient(config *config.Config, opts ...ClientOpts) (*FabricClient, error) {

	//initialize configuration information
	if err := config.Init(); err != nil {
		logger.Error("Configuration initialization failed")
		return nil, err
	}

	defCli, err := node.NewNodeCli(config.GetNodeApi(), config.GetAppInfo().AlgorithmType, config.GetAppCert().UserAppPrivateCert)
	if err != nil {
		return nil, err
	}

	fabricClient := &FabricClient{
		appInfo:     config.GetAppInfo(),
		userCode:    config.GetUserCode(),
		nodeClients: make(map[string]*node.NodeClient),
		users:       make(map[string]*msp.UserData),
	}

	for _, option := range opts {
		err := option(fabricClient)
		if err != nil {
			return nil, err
		}
	}

	if fabricClient.keyOpts == nil {
		fabricClient.keyOpts = keystore.NewFileKeyStore(config.GetKSPath())
	}

	if fabricClient.userOpts == nil {
		fabricClient.userOpts = keystore.NewUserCertStore(config.GetUSPath())
	}

	if fabricClient.defaultNodeName == "" {
		fabricClient.defaultNodeName = fabricClient.appInfo.MspId
	}
	fabricClient.nodeClients[fabricClient.defaultNodeName] = defCli

	return fabricClient, nil
}

type FabricClient struct {
	appInfo  config.AppInfo
	userCode string

	nodeClients     map[string]*node.NodeClient
	defaultNodeName string
	keyOpts         keystore.KeyStore
	userOpts        keystore.UserCertStore

	users map[string]*msp.UserData
}

type ClientOpts func(*FabricClient) error

// WithUserOpts 指定的用户证书存储对象，可由调用者自己按照接口 keystore.UserCertStore 实现
func WithUserOpts(userOpts keystore.UserCertStore) ClientOpts {
	return func(client *FabricClient) error {
		client.userOpts = userOpts
		return nil
	}
}

// WithUserOpts 指定的用户私钥存储对象，可由调用者自己按照接口 keystore.KeyStore 实现
func WithKeyOpts(keyOpts keystore.KeyStore) ClientOpts {
	return func(client *FabricClient) error {
		client.keyOpts = keyOpts
		return nil
	}
}

// WithDefaultNodeName 默认的客户端配置的城市名称，如果不指定，则为默认配置城市的MSPID
func WithDefaultNodeName(nodeName string) ClientOpts {
	return func(client *FabricClient) error {
		client.defaultNodeName = nodeName
		return nil
	}
}

// AddCityNode 增加一个城市接入配置
//  - nodeName 指定的城市名称标识，可以在执行多节点背书交易时，按照nodeName指定背书的城市列表
//  - gateWayUrl 网关地址
//  - privateKey 应用在该城市的接入私钥
func (c *FabricClient) AddCityNode(nodeName, gateWayUrl, privateKey string) error {

	_, ok := c.nodeClients[nodeName]
	if ok {
		return errors.New("node already exists")
	}

	Cli, err := node.NewNodeCli(gateWayUrl, c.appInfo.AlgorithmType, privateKey)
	if err != nil {
		return err
	}

	c.nodeClients[nodeName] = Cli

	return nil
}

func (c *FabricClient) GetHeader() base.ReqHeader {
	header := base.ReqHeader{
		UserCode: c.userCode,
		AppCode:  c.appInfo.AppCode,
	}

	return header
}

// subUserCertName sub user name in certificate
func (c *FabricClient) subUserCertName(name string) string {
	return name + "@" + c.appInfo.AppCode
}

func (c *FabricClient) Call(method string, req base.ReqInterface, res base.ResInterface) error {
	//todo random Call

	return c.nodeClients[c.defaultNodeName].Call(method, req, res)
}
