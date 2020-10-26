// Copyright (c) 2019. Baidu Inc. All Rights Reserved.

// package common is related to generate crypto client
package crypto

import (
	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/crypto/client/service/xchain"
)

// GetXchainCryptoClient get xchain crypto client
func GetXchainCryptoClient() *xchain.XchainCryptoClient {
	return &xchain.XchainCryptoClient{}
}

// GetGmCryptoClient get gm crypto client
func GetGmCryptoClient() *gm.GmCryptoClient {
	return &gm.GmCryptoClient{}
}
