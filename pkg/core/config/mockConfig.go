/**
 * @Author: Gao Chenxi
 * @Description:
 * @File:  mockConfig
 * @Version: 1.0.0
 * @Date: 2020/6/9 14:48
 */

package config

func NewMockFabricR1Config() (*Config, error) {
	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202004151958010871292",
		"app0001202004161020152918451",
		`-----BEGIN CERTIFICATE-----
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
		`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcRniHqapLZ4dwFpJ
Zo3ExKJfnRrYlOzHtLgWYEtiOy2hRANCAAQfFo0cjWXm9Fe1F/vKeYZM+5xIGAa8
pgvb1+c+s8bRqw+9xWvSoQv8AuP2TFJe4iTxZJE1tLxHVsREfH0mOH1p
-----END PRIVATE KEY-----`,
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewMockFabricSMConfig() (*Config, error) {
	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202007161739119605411",
		"app0001202007221352525405550",
		`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECwJ5ftuqndO9H3ks1hD8cB6IA9lx
/b0Z2hnFZ77rgRm9Q4lY1aqIhkM63Lh6X7uwPsoRC1xkS0PMp5x/jnRWcw==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgHOfxK/qJwIpj81VO
VXXPS5rJgVbB1kMhKChnk2/5ihigCgYIKoEcz1UBgi2hRANCAARkJeY6ywWCM5ch
kHjqi3XMBrabyIHqikq4Y6CgcNc8AEZGjS9rBKexOc8Yvm3OEzSnShBodp9G/HYn
r7IWHv/E
-----END PRIVATE KEY-----`,
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewNingBoFiscoConfig() (*Config, error) {

	return NewConfig(
		"https://ningbonode.bsngate.com:17602",
		"USER0001202005270937461865827",
		"app0001202005270940051261132",
		`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT
Vyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgRbYy4qWCVTNkWHsk
rf7BzHAb45HIL5rtATHDIJlSkz6gCgYIKoEcz1UBgi2hRANCAATA7r9nim0dOMFs
tz1EiBRZOg7Riv3Lr36vjn3l3iR5Ln5NGfbaYvbrtcId4bMEpAFWJErQErFj0CU6
M8E8TCxW
-----END PRIVATE KEY-----`,
		"test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewMockT1FiscoConfig() (*Config, error) {

	return NewConfig(
		"http://192.168.7.231:17505",
		"USER0001202006231440381776531",
		"app0001202006231731066689760",
		`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT
Vyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQghcqUjvK0biJFJi/w
P2KWmw+IwVh43lPPPGquzw7jCtigCgYIKoEcz1UBgi2hRANCAATux2zLQgS2+3k3
IxXPjjj/9/JqdT3x76gPn3rS53J6Nu3ZWiU/wRqaZhICr+RfzYl2xz/Eald4AWYJ
SorX/MpK
-----END PRIVATE KEY-----`,
		"test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewMockTestFiscoK1Config() (*Config, error) {

	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202006042321579692440",
		"app0001202006042323057101002",
		`-----BEGIN PUBLIC KEY-----
MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEh4WlY4pCv814i3WY5aRhtR3PoiIXOM1I
5xBGylyQTedo6DzJUdLfYZSZLs4py70D8FJtNICMVQCfezA7whHzUw==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGNAgEAMBAGByqGSM49AgEGBSuBBAAKBHYwdAIBAQQgs9DOx+bq2PlWVFRESHAM
VBKjDU9co5TIUzY203/utIugBwYFK4EEAAqhRANCAAR2T4i+jP7Tw1kFcHwGttKT
OMD7p1OHVE/evqTNlHRkYgDxEKBFE5Yoc/SsgStHhn9P9Isdz1xXYoiIzvPm9cFQ
-----END PRIVATE KEY-----`,
		"test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewMockTestFiscoSMConfig() (*Config, error) {

	return NewConfig(
		"http://192.168.1.43:17502",
		"USER0001202005281426464614357",
		"app0001202006221045063821068",
		`-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECwJ5ftuqndO9H3ks1hD8cB6IA9lx
/b0Z2hnFZ77rgRm9Q4lY1aqIhkM63Lh6X7uwPsoRC1xkS0PMp5x/jnRWcw==
-----END PUBLIC KEY-----`,
		`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg3ail5qa1WdSCaE4l
NDtKsH43sn4oLU2Q4Ag9g1zmEuWgCgYIKoEcz1UBgi2hRANCAATnkyph+Ukd5mSX
Dnr0d0JNH5lzMCYlFIf/8e3LOb8R1qvYEI/ePU6TVX7UcEbCAnVPlDMlv/oesYsn
j8PiaBZv
-----END PRIVATE KEY-----`,
		"test/msp",
		"F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}

func NewMockFiscoConfig1() (*Config, error) {
	config := &Config{
		nodeApi:  "https://ningbonode.bsngate.com:17602",
		mspDir:   " test/msp",
		httpCert: "F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
		appCert: certInfo{
			AppPublicCert: `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEIlh1C0iWAdcKnM/yAaZZT/42NVzT
Vyr31H9MDhHbPkp+/B3gsp5iZOr6OTAGO9jpN10/YMIrxt2IMg5auIEvMA==
-----END PUBLIC KEY-----`,
			UserAppPrivateCert: `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgRbYy4qWCVTNkWHsk
rf7BzHAb45HIL5rtATHDIJlSkz6gCgYIKoEcz1UBgi2hRANCAATA7r9nim0dOMFs
tz1EiBRZOg7Riv3Lr36vjn3l3iR5Ln5NGfbaYvbrtcId4bMEpAFWJErQErFj0CU6
M8E8TCxW
-----END PRIVATE KEY-----`,
		},
		user: userInfo{
			UserCode: "USER0001202005270937461865827",
		},
		app: appInfo{
			AppCode: "app0001202005270940051261132",
		},
	}

	err := config.Init()

	if err != nil {
		return nil, err
	}

	return config, nil
}

//NewMockConfig
func NewMockConfig() (*Config, error) {

	config := &Config{
		nodeApi:  "https://quanzhounode.bsngate.com:17602",
		mspDir:   "F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp",
		httpCert: "F:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
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

func NewMockT1Config() (*Config, error) {

	config := &Config{
		nodeApi:  "http://192.168.7.231:17505",
		mspDir:   "D:/gopath/src/github.com/BSNDA/PCNGateway-Go-SDK/test/msp",
		httpCert: "D:/gopath/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
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
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgVIRqaxmq4NSaQpbq
M4net13CpZmCvwGNxVgRyynWGSCgCgYIKoZIzj0DAQehRANCAASqVxhJEFyeJGr2
hNmom5z2ZE4++8KCv68RB9VjwQ8lCIZAQLfXj4peihVQVq2FBItnRTe/sSwdhuY/
eeZyjQmR
-----END PRIVATE KEY-----`,
		},
		user: userInfo{
			UserCode: "USER0001202006231440381776531",
		},
		app: appInfo{
			AppCode: "app0001202006231459040977686",
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
		"test/msp",
		"D:/Work/RedBaaS/04SourceCode/Gateway_sdk/src/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt",
	)
}
