package userstore

import (
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/msp"
)

type UserStore interface {
	Load(user *msp.UserData) error
	LoadAll(appCode string) []*msp.UserData
	Store(user *msp.UserData) error
}
