package config

const (
	_MspDir    = "D:/test/bsn-sdk-go/msp"
	_HttpsCert = "D:/test/bsn-sdk-go/cert/bsn_gateway_https.crt"
)

func Nt() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig(api, userCode, appCode, PubK_R1, privK, _MspDir, _HttpsCert)
}

// sm Trust
func NewFabricSMTrustConfig() (*Config, error) {
	api := "http://192.168.1.41:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig(api, userCode, appCode, PubK_SM_test, privK, _MspDir, _HttpsCert)
}

func NewFabricSMNoTrustConfig() (*Config, error) {
	api := "http://192.168.1.43:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007271538152051987"

	privK := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcBE//7DVT/JKe2KV
ZFpgM0sFgCXAKU4+FXFC1lTXizahRANCAATfygwwI4ssWL35HF0qZq/iyO1rnuEp
XTusvlrY9lWksKY96jw1nF/v0SeIpr2ThPYIvWq9egrTAG95ypznYxJC
-----END PRIVATE KEY-----`

	return NewConfig(api, userCode, appCode, PubK_SM_test, privK, _MspDir, _HttpsCert)
}

func NewMockFabricR1Config() (*Config, error) {

	api := "http://192.168.1.43:17502"
	userCode := "USER0001202004151958010871292"
	appCode := "app0001202004161020152918451"

	privK := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcRniHqapLZ4dwFpJ
Zo3ExKJfnRrYlOzHtLgWYEtiOy2hRANCAAQfFo0cjWXm9Fe1F/vKeYZM+5xIGAa8
pgvb1+c+s8bRqw+9xWvSoQv8AuP2TFJe4iTxZJE1tLxHVsREfH0mOH1p
-----END PRIVATE KEY-----`
	return NewConfig(api, userCode, appCode, PubK_R1_test, privK, _MspDir, _HttpsCert)
}

func NewMockFabricSMConfig() (*Config, error) {
	api := "http://192.168.1.43:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007221352525405550"
	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgHOfxK/qJwIpj81VO
VXXPS5rJgVbB1kMhKChnk2/5ihigCgYIKoEcz1UBgi2hRANCAARkJeY6ywWCM5ch
kHjqi3XMBrabyIHqikq4Y6CgcNc8AEZGjS9rBKexOc8Yvm3OEzSnShBodp9G/HYn
r7IWHv/E
-----END PRIVATE KEY-----`
	return NewConfig(api, userCode, appCode, PubK_SM_test, privK, _MspDir, _HttpsCert)
}

func NewMockT1FiscoConfig() (*Config, error) {
	api := "http://192.168.7.231:17505"
	userCode := "USER0001202006231440381776531"
	appCode := "app0001202006231731066689760"
	privK := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT
Vyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==
-----END PUBLIC KEY-----`
	return NewConfig(api, userCode, appCode, PubK_SM, privK, _MspDir, _HttpsCert)

}

func NewMockTestFiscoK1Config() (*Config, error) {

	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202006042321579692440",
		"app0001202006042323057101002",
		Pubk_K1_test,
		`-----BEGIN PRIVATE KEY-----
MIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQgs9DOx+bq2PlWVFRESHAM
VBKjDU9co5TIUzY203/utIugBwYFK4EEAAqhRANCAAR2T4i+jP7Tw1kFcHwGttKT
OMD7p1OHVE/evqTNlHRkYgDxEKBFE5Yoc/SsgStHhn9P9Isdz1xXYoiIzvPm9cFQ
-----END PRIVATE KEY-----`,
		_MspDir,
		_HttpsCert,
	)
}

func NewMockTestFiscoSMConfig() (*Config, error) {

	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202005281426464614357",
		"app0001202006221045063821068",
		PubK_SM_test,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg3ail5qa1WdSCaE4l
NDtKsH43sn4oLU2Q4Ag9g1zmEuWgCgYIKoEcz1UBgi2hRANCAATnkyph+Ukd5mSX
Dnr0d0JNH5lzMCYlFIf/8e3LOb8R1qvYEI/ePU6TVX7UcEbCAnVPlDMlv/oesYsn
j8PiaBZv
-----END PRIVATE KEY-----`,
		_MspDir,
		_HttpsCert,
	)
}

//NewMockConfig
func NewMockConfig() (*Config, error) {

	config := &Config{
		nodeApi:  "https://quanzhounode.bsngate.com:17602",
		mspDir:   _MspDir,
		httpCert: _HttpsCert,
		appCert: certInfo{
			AppPublicCert: `-----BEGIN CERTIFICATE-----
MIIC+zCCAqGgAwIBAgIUARhAfFSyhzcx9q4LdiYKl2UHo1YwCgYIKoZIzj0EAwIw
TjELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaWppbmcxDDAKBgNVBAoTA0JTTjEP
MA0GA1UECxMGY2xpZW50MQ4wDAYDVQQDEwVic25jYTAgFw0xOTA5MjYxMDI0MDBa
GA8yMDk5MDkwNTAyMDQwMFowgZYxCzAJBgNVBAYTAkNOMREwDwYDVQQIEwhDaGFu
Z3NoYTEOMAwGA1UEChMFQ21QYXkxPTALBgNVBAsTBHVzZXIwEgYDVQQLEwtob25n
emFvbm9kZTAOBgNVBAsTB2JzbmJhc2UwCgYDVQQLEwNjb20xJTAjBgNVBAMMHG5v
ZGVAaG9uZ3phb25vZGUuYnNuYmFzZS5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMB
BwNCAAQ/X2w5+pJoZXNCO81T4xMR+TxmFoYk6eG1kYML8HBPrUT6QflxtDXYsE9h
SzVAovq5DHww3vD8Xft/mxwsAXyuo4IBEDCCAQwwDgYDVR0PAQH/BAQDAgeAMAwG
A1UdEwEB/wQCMAAwHQYDVR0OBBYEFDPVPRqPANJavkNOg/WhPkUkH6wqMB8GA1Ud
IwQYMBaAFJuwcYba1G07p1ySkpzyes8L79OPMCUGA1UdEQQeMByCGmNhLmhvbmd6
YW9ub2RlLmJzbmJhc2UuY29tMIGEBggqAwQFBgcIAQR4eyJhdHRycyI6eyJoZi5B
ZmZpbGlhdGlvbiI6Imhvbmd6YW9ub2RlLmJzbmJhc2UuY29tIiwiaGYuRW5yb2xs
bWVudElEIjoibm9kZUBob25nemFvbm9kZS5ic25iYXNlLmNvbSIsImhmLlR5cGUi
OiJ1c2VyIn19MAoGCCqGSM49BAMCA0gAMEUCIQD7FBAQJsgS0uhaL4mjJgILdFfY
RKXvNutyKz/MqJ54RQIgNS67sSUCbOZRx1rWDqYEcBF1zypEFik25fNgY3zk5gM=
-----END CERTIFICATE-----`,
			UserAppPrivateCert: `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgHX/6SmzxMrQCCQZg
aInjUENx5zcaPUd+af9EX8WlGrKhRANCAATCFSZlYUREYsgHdQePEenfnv6YuiKB
b6nD3mDaLxvv/xidH0sz14oHXS15E4AvtSra8sUBugtqrgMcg0gUmNAz
-----END PRIVATE KEY-----`,
		},
		user: userInfo{
			UserCode: "reddate",
		},
		app: appInfo{
			AppCode: "CL1851016378620191011150518",
		},
	}

	err := config.Init()

	if err != nil {
		return nil, err
	}

	return config, nil
}

//
func NewMockXuperchainSMConfig() (*Config, error) {

	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0006202007171549487497611",
		"app0006202007171545196904721",
		`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECwJ5ftuqndO9H3ks1hD8cB6IA9lx
/b0Z2hnFZ77rgRm9Q4lY1aqIhkM63Lh6X7uwPsoRC1xkS0PMp5x/jnRWcw==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgzPnol3cHQyF6tJ4k
cr8zqfEQrsfFM0MOmHzHWNvJuVygCgYIKoEcz1UBgi2hRANCAATAfT0PI75pduz/
LhD11NFPd2RYCrgeigmAbKBPs9Vj5FlrW3+PFVQAR7OJcUmB8MYR33VS1hCINv9u
EhxwJo5Q
-----END PRIVATE KEY-----`,
		_MspDir,
		_HttpsCert,
	)
}
