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

// saveKey delegate method for saving private key
type KeyOpts func(rawPem []byte, alias string) error
type ClientOpts func(*FabricClient) error

type FabricClient struct {
	appInfo  config.AppInfo
	userCode string

	nodeClients     map[string]*node.NodeClient
	defaultNodeName string
	keyOpts         keystore.KeyStore
	userOpts        keystore.UserCertStore

	users map[string]*msp.UserData
}

func WithUserOpts(userOpts keystore.UserCertStore) ClientOpts {
	return func(client *FabricClient) error {
		client.userOpts = userOpts
		return nil
	}
}

func WithKeyOpts(keyOpts keystore.KeyStore) ClientOpts {
	return func(client *FabricClient) error {
		client.keyOpts = keyOpts
		return nil
	}
}

func WithDefaultNodeName(nodeName string) ClientOpts {
	return func(client *FabricClient) error {
		client.defaultNodeName = nodeName
		return nil
	}
}

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
