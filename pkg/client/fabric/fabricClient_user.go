package fabric

import (
	"encoding/hex"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/base"

	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
	"github.com/BSNDA/bsn-sdk-crypto/key"
	"io/ioutil"
	"path"

	"github.com/pkg/errors"
	"github.com/wonderivan/logger"

	userreq "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/user"
	userres "github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/res/fabric/user"
)

const (
	RegisterUser = "user/register"
	EnrollUser   = "user/enroll"
)

// RegisterUser register sub user
func (c *FabricClient) RegisterUser(body userreq.RegisterReqDataBody) (*userres.RegisterResData, error) {

	req := &userreq.RegisterReqData{}
	req.Header = c.GetHeader()
	req.Body = body
	res := &userres.RegisterResData{}

	err := c.Call(RegisterUser, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", RegisterUser)
	}
	return res, nil
}

// DefaultSaveKey Default delegate method for saving private key
func DefaultSaveKey(dir string) KeyOpts {
	return func(rawPem []byte, alias string) error {
		keyFile := path.Join(dir, alias+"_sk")
		err := ioutil.WriteFile(keyFile, rawPem, 0600)
		if err != nil {
			logger.Debug("Failed storing private key [%s]: [%s]", keyFile, err)
			return err
		}
		return nil
	}
}

// EnrollUser enroll sub user certificate and store to local folder and FabricClient.Users
func (c *FabricClient) EnrollUser(body userreq.RegisterReqDataBody) (*userres.EnrollResData, error) {

	enrollBody := userreq.EnrollReqDataBody{
		Name:   body.Name,
		Secret: body.Secret,
	}

	privKey, err := key.NewPrivateKeyByGen(c.appInfo.AlgorithmType.ToKeyType())
	if err != nil {
		return nil, errors.WithMessage(err, "Generate private key has error")
	}

	csrReq := key.NewCertificateRequest(c.subUserCertName(enrollBody.Name))

	csrBytes, err := privKey.GenCSR(csrReq)
	if err != nil {
		return nil, errors.WithMessage(err, "Generate csr has error")
	}

	rawKey, err := privKey.KeyPEM()
	if err != nil {
		return nil, errors.WithMessage(err, "privateKey to PEM has error")
	}

	// save key
	err = c.keyOpts.StoreKey(rawKey, hex.EncodeToString(privKey.SKI()))
	if err != nil {
		return nil, errors.WithMessage(err, "save privateKey has error")
	}

	user := c.newUser(body.Name)
	user.PrivateKey = privKey

	enrollBody.CsrPem = string(csrBytes)
	res, err := c.enroll(enrollBody)
	if err != nil {
		return nil, err
	}

	if res.Header.Code == base.Header_error_code {
		return nil, errors.New(res.Header.Msg)
	}

	user.EnrollmentCertificate = []byte(res.Body.Cert)
	err = c.userOpts.Store(user)
	if err != nil {
		logger.Warn("user [%s] cert store has error : %s", user.UserName, err.Error())
	}

	c.users[body.Name] = user

	return res, nil
}

// enroll enroll sub user certificate
func (c *FabricClient) enroll(body userreq.EnrollReqDataBody) (*userres.EnrollResData, error) {

	req := &userreq.EnrollReqData{}
	req.Header = c.GetHeader()
	req.Body = body

	res := &userres.EnrollResData{}

	err := c.Call(EnrollUser, req, res)
	if err != nil {
		return nil, errors.WithMessagef(err, "call %s has error", EnrollUser)
	}
	return res, nil

}

func (c *FabricClient) newUser(userName string) *msp.UserData {
	user := &msp.UserData{
		UserName: userName,
		AppCode:  c.appInfo.AppCode,
		MspId:    c.appInfo.MspId,
		TxHash:   c.appInfo.TxHash(),
	}

	return user
}

// LoadUser load user from local store , before, the cache is checked from the client users
func (c *FabricClient) LoadUser(userName string) (*msp.UserData, error) {

	user, ok := c.users[userName]
	if ok {
		return user, nil
	}

	user = c.newUser(userName)
	err := c.userOpts.Load(user)
	if err != nil {
		return nil, err
	}

	puk, err := key.NewPublicProvider(c.appInfo.AlgorithmType.ToKeyType(), string(user.EnrollmentCertificate))
	if err != nil {
		return nil, errors.WithMessage(err, "cert pem load has error")
	}

	alias := hex.EncodeToString(puk.SKI())

	rawkey, err := c.keyOpts.LoadKey(alias)
	if err != nil {
		return nil, errors.WithMessage(err, "load private key has error")
	}

	k, err := key.NewPrivateKeyProvider(c.appInfo.AlgorithmType.ToKeyType(), string(rawkey))
	if err != nil {
		return nil, errors.WithMessage(err, "new private key provider has error")
	}

	user.PrivateKey = k
	c.users[userName] = user
	return user, nil
}
