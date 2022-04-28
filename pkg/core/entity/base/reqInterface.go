package base

type ReqInterface interface {
	SetMac(mac string)
	GetEncryptionValue() string
}
