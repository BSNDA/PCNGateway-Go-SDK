package fabric

import (
	"encoding/json"
	"fmt"

	"github.com/wonderivan/logger"

	"github.com/BSNDA/bsn-sdk-crypto/crypto/secp256r1"
	"github.com/BSNDA/bsn-sdk-crypto/crypto/sm"
	ksecdsa "github.com/BSNDA/bsn-sdk-crypto/keystore/ecdsa"
	kssm "github.com/BSNDA/bsn-sdk-crypto/keystore/sm"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/cert"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/enum"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	userreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/user"
	userres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/user"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/sign"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/util/http"
)

func (c *FabricClient) RegisterUser(body userreq.RegisterReqDataBody) (*userres.RegisterResData, error) {

	url := c.GetURL("/api/fabric/v1/user/register")

	data := &userreq.RegisterReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userres.RegisterResData{}

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

func (c *FabricClient) EnrollUser(body userreq.RegisterReqDataBody) error {

	enrollBody := userreq.EnrollReqDataBody{
		Name:   body.Name,
		Secret: body.Secret,
	}

	csr, key, err := cert.GetCSRPEM(c.GetCertName(enrollBody.Name), c.Config.GetAppInfo().AlgorithmType, c.ks)
	if err != nil {
		fmt.Println(err)
		return err
	}

	enrollBody.CsrPem = csr

	res, err := c.enroll(enrollBody)

	if err != nil {
		return err
	}

	if res.Header.Code == base.Header_success_code {

		var pk interface{}
		var sh sign.SignHandle

		if c.Config.GetAppInfo().AlgorithmType == enum.AppAlgorithmType_SM2 {
			pk = kssm.GetSmPrivateKey(key)
			sh = sm.NewTransUserSMHandle(kssm.GetSmPrivateKey(key))

		} else {
			pk = ksecdsa.GetECDSAPrivateKey(key)
			sh = secp256r1.NewTransUserR1Handle(ksecdsa.GetECDSAPrivateKey(key))
		}
		user := &msp.UserData{
			UserName:              enrollBody.Name,
			AppCode:               c.Config.GetAppInfo().AppCode,
			MspId:                 c.Config.GetAppInfo().MspId,
			EnrollmentCertificate: []byte(res.Body.Cert),
			PrivateKey:            pk,
		}
		user.SetSignHandle(sh)

		c.us.Store(user)
		c.Users[user.UserName] = user
	} else {
		return errors.New(res.Header.Msg)
	}

	return nil

}

func (c *FabricClient) enroll(body userreq.EnrollReqDataBody) (*userres.EnrollResData, error) {

	url := c.GetURL("/api/fabric/v1/user/enroll")

	data := &userreq.EnrollReqData{}
	data.Header = c.GetHeader()
	data.Body = body
	data.Mac = c.Sign(data.GetEncryptionValue())

	reqBytes, _ := json.Marshal(data)

	resBytes, err := http.SendPost(reqBytes, url, c.Config.GetCert())

	if err != nil {
		logger.Error("gateway interface call failed：", err)
		return nil, err
	}

	res := &userres.EnrollResData{}

	err = json.Unmarshal(resBytes, res)

	if err != nil {
		logger.Error("return parameter serialization failed：", err)
		return nil, err
	}

	return res, nil

}
