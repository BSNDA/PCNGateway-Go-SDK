package config

const (
	_MspDir    = "D:/test/bsn-sdk-go/msp"
	_HttpsCert = "D:/test/bsn-sdk-go/cert/bsn_gateway_https.crt"
)

func NewMockFabricConfig() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}

func NewMockFiscoConfig() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}

func NewMockXuperchainConfig() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}
