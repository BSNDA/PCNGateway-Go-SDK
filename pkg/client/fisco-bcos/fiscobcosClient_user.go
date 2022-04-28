package fisco_bcos

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	userReq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fiscobcos/user"
	userRes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fiscobcos/user"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"path"
)

// RegisterUser register sub user
func (c *FiscoBcosClient) RegisterUser(body userReq.RegisterReqDataBody) (*userRes.RegisterResData, error) {

	url := c.GetURL("/api/fiscobcos/v1/user/register")

	data := &userReq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url)

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userRes.RegisterResData{}

	err = json.Unmarshal(resBytes, res)
	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}
	if !c.Verify(res.Mac, res.GetEncryptionValue()) {
		return nil, errors.New("sign has error")
	}

	return res, nil
}

func (c *FiscoBcosClient) getUser(userName string) (key.PrivateKeyProvider, error) {
	userPath := c.Config.GetKSPath()

	fileName := getKeyFileName(userName, c.Config.GetAppInfo().AppCode)

	filePath := path.Join(userPath, fileName)

	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		//create key

		privKey, err := key.NewPrivateKeyByGen(c.Config.GetAppInfo().AlgorithmType.ToKeyType())

		if err != nil {
			return nil, err
		}
		keyBytes, err := privKey.KeyPEM()
		if err != nil {
			return nil, err
		}
		err = storeKey(keyBytes, filePath)
		if err != nil {
			return nil, err
		}
		return privKey, nil

	} else {
		bytes, err := ioutil.ReadFile(filePath) // nolint: gas
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, errors.New("user key error")
		}

		return key.NewPrivateKeyProvider(c.Config.GetAppInfo().AlgorithmType.ToKeyType(), string(bytes))

	}

}

func storeKey(keyBytes []byte, filePath string) error {

	err := os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, keyBytes, 0600)
}

func getKeyFileName(name string, appCode string) string {
	return fmt.Sprintf("%s@%s_pk", name, appCode)
}
