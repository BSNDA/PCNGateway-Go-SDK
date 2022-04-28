package keystore

type KeyStore interface {
	StoreKey(rawPem []byte, alias string) error

	LoadKey(alias string) ([]byte, error)
}
