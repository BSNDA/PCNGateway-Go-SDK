package keystore

import (
	"github.com/BSNDA/bsn-sdk-crypto/keystore/key"
	"github.com/BSNDA/bsn-sdk-crypto/utils"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/wonderivan/logger"

	ksecdsa "github.com/BSNDA/bsn-sdk-crypto/keystore/ecdsa"
	kssm "github.com/BSNDA/bsn-sdk-crypto/keystore/sm"

	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// NewFileBasedKeyStore instantiated a file-based key store at a given position.
// The key store can be encrypted if a non-empty password is specified.
// It can be also be set as read only. In this case, any store operation
// will be forbidden
func NewFileBasedKeyStore(pwd []byte, path string, readOnly bool) (key.KeyStore, error) {
	ks := &fileBasedKeyStore{}
	return ks, ks.Init(pwd, path, readOnly)
}

// fileBasedKeyStore is a folder-based KeyStore.
// Each key is stored in a separated file whose name contains the key's SKI
// and flags to identity the key's type. All the keys are stored in
// a folder whose path is provided at initialization time.
// The KeyStore can be initialized with a password, this password
// is used to encrypt and decrypt the files storing the keys.
// A KeyStore can be read only to avoid the overwriting of keys.
type fileBasedKeyStore struct {
	path string

	readOnly bool
	isOpen   bool

	pwd []byte

	// Sync
	m sync.Mutex
}

// Init initializes this KeyStore with a password, a path to a folder
// where the keys are stored and a read only flag.
// Each key is stored in a separated file whose name contains the key's SKI
// and flags to identity the key's type.
// If the KeyStore is initialized with a password, this password
// is used to encrypt and decrypt the files storing the keys.
// The pwd can be nil for non-encrypted KeyStores. If an encrypted
// key-store is initialized without a password, then retrieving keys from the
// KeyStore will fail.
// A KeyStore can be read only to avoid the overwriting of keys.
func (ks *fileBasedKeyStore) Init(pwd []byte, path string, readOnly bool) error {
	// Validate inputs
	// pwd can be nil

	if len(path) == 0 {
		return errors.New("An invalid KeyStore path provided. Path cannot be an empty string.")
	}

	ks.m.Lock()
	defer ks.m.Unlock()

	if ks.isOpen {
		return errors.New("KeyStore already initilized.")
	}

	ks.path = path
	ks.pwd = utils.Clone(pwd)

	err := ks.createKeyStoreIfNotExists()
	if err != nil {
		return err
	}

	err = ks.openKeyStore()
	if err != nil {
		return err
	}

	ks.readOnly = readOnly

	return nil
}

// ReadOnly returns true if this KeyStore is read only, false otherwise.
// If ReadOnly is true then StoreKey will fail.
func (ks *fileBasedKeyStore) ReadOnly() bool {
	return ks.readOnly
}

// GetKey returns a key object whose SKI is the one passed.
func (ks *fileBasedKeyStore) GetKey(ski []byte) (key.Key, error) {
	// Validate arguments
	if len(ski) == 0 {
		return nil, errors.New("Invalid SKI. Cannot be of zero length.")
	}

	suffix := ks.getSuffix(hex.EncodeToString(ski))
	switch suffix {
	//case "key":
	//	// Load the key
	//	key, err := ks.loadKey(hex.EncodeToString(ski))
	//	if err != nil {
	//		return nil, fmt.Errorf("Failed loading key [%x] [%s]", ski, err)
	//	}
	//
	//	return &aesPrivateKey{key, false}, nil
	case "sk":
		// Load the private key
		key, err := ks.loadPrivateKey(hex.EncodeToString(ski))
		if err != nil {
			return nil, fmt.Errorf("Failed loading secret key [%x] [%s]", ski, err)
		}

		switch key.(type) {
		case *ecdsa.PrivateKey:
			return ksecdsa.NewEcdsaPrivateKey(key.(*ecdsa.PrivateKey)), nil
		//case *rsa.PrivateKey:
		//	return &rsaPrivateKey{key.(*rsa.PrivateKey)}, nil
		case *sm2.PrivateKey:
			return kssm.NewSMPrivateKey(key.(*sm2.PrivateKey)), nil
		default:
			return nil, errors.New("Secret key type not recognized")
		}
	case "pk":
		// Load the public key
		key, err := ks.loadPublicKey(hex.EncodeToString(ski))
		if err != nil {
			return nil, fmt.Errorf("Failed loading public key [%x] [%s]", ski, err)
		}

		switch key.(type) {
		case *ecdsa.PublicKey:
			return ksecdsa.NewEcdsaPublicKey(key.(*ecdsa.PublicKey)), nil
		case *sm2.PublicKey:
			return kssm.NewSMPublicKey(key.(*sm2.PublicKey)), nil
		default:
			return nil, errors.New("Public key type not recognized")
		}
	default:
		return ks.searchKeystoreForSKI(ski)
	}
}

// StoreKey stores the key k in this KeyStore.
// If this KeyStore is read only then the method will fail.
func (ks *fileBasedKeyStore) StoreKey(k key.Key) (err error) {
	if ks.readOnly {
		return errors.New("Read only KeyStore.")
	}

	if k == nil {
		return errors.New("Invalid key. It must be different from nil.")
	}
	switch k.(type) {
	case *ksecdsa.EcdsaPrivateKey:
		kk := k.(*ksecdsa.EcdsaPrivateKey)
		err = ks.storePrivateKey(hex.EncodeToString(k.SKI()), kk.GetPrivateKey())
		if err != nil {
			return fmt.Errorf("Failed storing ECDSA private key [%s]", err)
		}
	case *kssm.SMPrivateKey:
		kk := k.(*kssm.SMPrivateKey)
		err = ks.storePrivateKey(hex.EncodeToString(k.SKI()), kk.GetPrivateKey())
		if err != nil {
			return fmt.Errorf("Failed storing ECDSA private key [%s]", err)
		}
	case *ksecdsa.EcdsaPublicKey:
		kk := k.(*ksecdsa.EcdsaPublicKey)

		err = ks.storePublicKey(hex.EncodeToString(k.SKI()), kk.GetPublicKey())
		if err != nil {
			return fmt.Errorf("Failed storing ECDSA public key [%s]", err)
		}
	case *kssm.SMPublicKey:
		kk := k.(*kssm.SMPublicKey)

		err = ks.storePublicKey(hex.EncodeToString(k.SKI()), kk.GetPublicKey())
		if err != nil {
			return fmt.Errorf("Failed storing ECDSA public key [%s]", err)
		}
	//case *rsaPrivateKey:
	//	kk := k.(*rsaPrivateKey)
	//
	//	err = ks.storePrivateKey(hex.EncodeToString(k.SKI()), kk.privKey)
	//	if err != nil {
	//		return fmt.Errorf("Failed storing RSA private key [%s]", err)
	//	}
	//
	//case *rsaPublicKey:
	//	kk := k.(*rsaPublicKey)
	//
	//	err = ks.storePublicKey(hex.EncodeToString(k.SKI()), kk.pubKey)
	//	if err != nil {
	//		return fmt.Errorf("Failed storing RSA public key [%s]", err)
	//	}
	//
	//case *aesPrivateKey:
	//	kk := k.(*aesPrivateKey)
	//
	//	err = ks.storeKey(hex.EncodeToString(k.SKI()), kk.privKey)
	//	if err != nil {
	//		return fmt.Errorf("Failed storing AES key [%s]", err)
	//	}

	default:
		return fmt.Errorf("Key type not reconigned [%s]", k)
	}

	return
}

func (ks *fileBasedKeyStore) searchKeystoreForSKI(ski []byte) (k key.Key, err error) {

	files, _ := ioutil.ReadDir(ks.path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if f.Size() > (1 << 16) { //64k, somewhat arbitrary limit, considering even large RSA keys
			continue
		}

		raw, err := ioutil.ReadFile(filepath.Join(ks.path, f.Name()))
		if err != nil {
			continue
		}

		key, err := utils.PEMtoPrivateKey(raw, ks.pwd)
		if err != nil {
			continue
		}

		switch key.(type) {
		case *ecdsa.PrivateKey:
			k = ksecdsa.NewEcdsaPrivateKey(key.(*ecdsa.PrivateKey))
		case *sm2.PrivateKey:
			k = kssm.NewSMPrivateKey(key.(*sm2.PrivateKey))
		default:
			continue
		}

		if !bytes.Equal(k.SKI(), ski) {
			continue
		}

		return k, nil
	}

	return nil, fmt.Errorf("Key with SKI %s not found in %s", hex.EncodeToString(ski), ks.path)
}

func (ks *fileBasedKeyStore) getSuffix(alias string) string {
	files, _ := ioutil.ReadDir(ks.path)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), alias) {
			if strings.HasSuffix(f.Name(), "sk") {
				return "sk"
			}
			if strings.HasSuffix(f.Name(), "pk") {
				return "pk"
			}
			if strings.HasSuffix(f.Name(), "key") {
				return "key"
			}
			break
		}
	}
	return ""
}

func (ks *fileBasedKeyStore) storePrivateKey(alias string, privateKey interface{}) error {
	//fmt.Println("SKi:",alias)
	rawKey, err := utils.PrivateKeyToPEM(privateKey, ks.pwd)

	if err != nil {
		logger.Debug("Failed converting private key to PEM [%s]: [%s]", alias, err)
		return err
	}
	err = ioutil.WriteFile(ks.getPathForAlias(alias, "sk"), rawKey, 0600)
	if err != nil {
		logger.Debug("Failed storing private key [%s]: [%s]", alias, err)
		return err
	}

	return nil
}

func (ks *fileBasedKeyStore) storePublicKey(alias string, publicKey interface{}) error {
	rawKey, err := utils.PublicKeyToPEM(publicKey, ks.pwd)
	if err != nil {
		logger.Error("Failed converting public key to PEM [%s]: [%s]", alias, err)
		return err
	}
	err = ioutil.WriteFile(ks.getPathForAlias(alias, "pk"), rawKey, 0600)
	if err != nil {
		logger.Error("Failed storing private key [%s]: [%s]", alias, err)
		return err
	}

	return nil
}

func (ks *fileBasedKeyStore) storeKey(alias string, key []byte) error {
	pem, err := utils.AEStoEncryptedPEM(key, ks.pwd)
	if err != nil {
		logger.Error("Failed converting key to PEM [%s]: [%s]", alias, err)
		return err
	}

	err = ioutil.WriteFile(ks.getPathForAlias(alias, "key"), pem, 0600)
	if err != nil {
		logger.Error("Failed storing key [%s]: [%s]", alias, err)
		return err
	}

	return nil
}

func (ks *fileBasedKeyStore) loadPrivateKey(alias string) (interface{}, error) {
	path := ks.getPathForAlias(alias, "sk")

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("Failed loading private key [%s]: [%s].", alias, err.Error())

		return nil, err
	}

	privateKey, err := utils.PEMtoPrivateKey(raw, ks.pwd)
	if err != nil {

		smk, err := sm2.ReadPrivateKeyFromMem(raw, ks.pwd)
		if err != nil {
			logger.Error("Failed parsing private key [%s]: [%s].", alias, err.Error())

			return nil, err
		}

		return smk, nil

	}

	return privateKey, nil
}

func (ks *fileBasedKeyStore) loadPublicKey(alias string) (interface{}, error) {
	path := ks.getPathForAlias(alias, "pk")
	logger.Debug("Loading public key [%s] at [%s]...", alias, path)

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("Failed loading public key [%s]: [%s].", alias, err.Error())

		return nil, err
	}

	privateKey, err := utils.PEMtoPublicKey(raw, ks.pwd)
	if err != nil {
		logger.Error("Failed parsing private key [%s]: [%s].", alias, err.Error())

		return nil, err
	}

	return privateKey, nil
}

func (ks *fileBasedKeyStore) loadKey(alias string) ([]byte, error) {
	path := ks.getPathForAlias(alias, "key")
	logger.Debug("Loading key [%s] at [%s]...", alias, path)

	pem, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("Failed loading key [%s]: [%s].", alias, err.Error())

		return nil, err
	}

	key, err := utils.PEMtoAES(pem, ks.pwd)
	if err != nil {
		logger.Error("Failed parsing key [%s]: [%s]", alias, err)

		return nil, err
	}

	return key, nil
}

func (ks *fileBasedKeyStore) createKeyStoreIfNotExists() error {
	// Check keystore directory
	ksPath := ks.path
	missing, err := utils.DirMissingOrEmpty(ksPath)

	if missing {
		logger.Debug("KeyStore path [%s] missing [%t]: [%s]", ksPath, missing, utils.ErrToString(err))

		err := ks.createKeyStore()
		if err != nil {
			logger.Error("Failed creating KeyStore At [%s]: [%s]", ksPath, err.Error())
			return nil
		}
	}

	return nil
}

func (ks *fileBasedKeyStore) createKeyStore() error {
	// Create keystore directory root if it doesn't exist yet
	ksPath := ks.path
	logger.Debug("Creating KeyStore at [%s]...", ksPath)

	os.MkdirAll(ksPath, 0755)

	logger.Debug("KeyStore created at [%s].", ksPath)
	return nil
}

func (ks *fileBasedKeyStore) openKeyStore() error {
	if ks.isOpen {
		return nil
	}
	ks.isOpen = true

	return nil
}

func (ks *fileBasedKeyStore) getPathForAlias(alias, suffix string) string {
	return filepath.Join(ks.path, alias+"_"+suffix)
}