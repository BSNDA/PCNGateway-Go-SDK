package fabric

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/keystore"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/userstore"
	"github.com/BSNDA/PCNGateway-Go-SDK/third_party/github.com/hyperledger/fabric/bccsp"
	"github.com/wonderivan/logger"
)

//initialize client requested by fabric
func InitFabricClient(config *config.Config) (*FabricClient, error) {

	//initialize configuration information
	if err := config.Init(); err != nil {
		logger.Error("Configuration initialization failed")
		return nil, err
	}
	//generate a private key handler
	ks, err := keystore.NewFileBasedKeyStore(nil, config.GetKSPath(), false)

	if err != nil {
		logger.Error("keystore initialization failed")
		return nil, err
	}
	//generate a cert handler
	us := userstore.NewUserStore(config.GetUSPath())

	client := client.Client{
		Config: config,
	}

	fabricClient := &FabricClient{
		Client: client,
		ks:     ks,
		us:     us,
		users:  make(map[string]*msp.UserData),
	}
	//configure the algorithm type of user signature and generate a unified signature verification handler
	err = fabricClient.SetAlgorithm(config.GetAppInfo().AlgorithmType, config.GetAppCert().AppPublicCert, config.GetAppCert().UserAppPrivateCert)

	if err != nil {
		logger.Error("signHandle initialization failed:%v", err)
		return nil, err
	}
	//load the client's info generated locally
	fabricClient.LoadUser()

	return fabricClient, nil
}

type FabricClient struct {
	client.Client
	ks    bccsp.KeyStore
	us    userstore.UserStore
	users map[string]*msp.UserData
}

func (c *FabricClient) GetCertName(name string) string {
	return name + "@" + c.Config.GetAppInfo().AppCode
}

func (c *FabricClient) LoadUser() {

	users := c.us.LoadAll(c.Config.GetAppInfo().AppCode)

	for i, _ := range users {

		user := users[i]
		user.MspId = c.Config.GetAppInfo().MspId

		err := keystore.LoadKey(user, c.ks, c.Config.GetAppInfo().AlgorithmType)

		if err == nil {
			c.users[user.UserName] = user
		}
	}

}

func (c *FabricClient) GetUser(name string) (*msp.UserData, error) {
	user, ok := c.users[name]
	if ok {
		return user, nil
	} else {
		return nil, errors.New("user does not exist")
	}

}

func (c *FabricClient) LoadLocalUser(name string) (*msp.UserData, error) {

	user := &msp.UserData{
		UserName: name,
		AppCode:  c.Config.GetAppInfo().AppCode,
	}

	err := c.us.Load(user)
	if err != nil {
		return nil, err
	}
	err = keystore.LoadKey(user, c.ks, c.Config.GetAppInfo().AlgorithmType)

	if err != nil {
		return nil, err

	} else {
		c.users[user.UserName] = user
		return user, nil

	}

}
