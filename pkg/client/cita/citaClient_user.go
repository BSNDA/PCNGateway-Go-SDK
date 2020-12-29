package cita

import (
	"encoding/json"
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/cert"
	userReq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/cita/user"
	userRes "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/cita/user"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"path"
)

func (c *CitaClient) RegisterUser(body userReq.RegisterReqDataBody) (*userRes.RegisterResData, error) {

	url := c.GetURL("/api/cita/v1/user/register")

	data := &userReq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

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

func (c *CitaClient) getUser(userName string) (interface{}, error) {
	userPath := c.Config.GetKSPath()

	fileName := getKeyFileName(userName, c.Config.GetAppInfo().AppCode)

	filePath := path.Join(userPath, fileName)

	if _, err1 := os.Stat(filePath); os.IsNotExist(err1) {
		//create key
		key, keyBytes, err := cert.NewUser(c.Config.GetAppInfo().AlgorithmType)
		if err != nil {
			return nil, err
		}
		err = storeKey(keyBytes, filePath)
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

		return cert.GetUserKey(bytes, c.Config.GetAppInfo().AlgorithmType)

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
