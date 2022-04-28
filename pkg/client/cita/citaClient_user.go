package cita

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"io/ioutil"
	"os"
	"path"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/http"
	userReq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/user"
	userRes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/user"
	"github.com/wonderivan/logger"
)

//register user
func (c *CitaClient) RegisterUser(body userReq.RegisterReqDataBody) (*userRes.RegisterResData, error) {

	url := c.GetURL("/api/cita/v1/user/register")

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
	//if !c.Verify(res.Mac, res.GetEncryptionValue()) {
	//	return nil, errors.New("sign has error")
	//}

	return res, nil
}

func (c *CitaClient) getUser(userName string) (key.PrivateKeyProvider, error) {
	userPath := c.Config.GetKSPath()

	fileName := getKeyFileName(userName, c.Config.GetAppInfo().AppCode)

	filePath := path.Join(userPath, fileName)

	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		//create key
		key, err := key.NewPrivateKeyByGen(c.Config.GetAppInfo().AlgorithmType.ToKeyType())
		if err != nil {
			return nil, err
		}

		ketBytes, err := key.KeyPEM()
		if err != nil {
			return nil, err
		}
		err = storeKey(ketBytes, filePath)
		if err != nil {
			return nil, err
		}
		return key, nil

	} else {
		bytes, err := ioutil.ReadFile(filePath) // nolint: gas
		if err != nil {
			return nil, err
		}
		if bytes == nil {
			return nil, errors.New("user key error")
		}

		k, err := key.NewPrivateKeyProvider(c.Config.GetAppInfo().AlgorithmType.ToKeyType(), string(bytes))

		return k, err

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
