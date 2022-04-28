package keystore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
)

type UserCertStore interface {
	Load(user *msp.UserData) error
	LoadAll(appCode string) []*msp.UserData
	Store(user *msp.UserData) error
}
