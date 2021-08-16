package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/common/errors"
	"math/big"
	"math/rand"
)

const (
	// NonceSize is the default NonceSize
	NonceSize = 24
)

// GetRandomBytes returns len random looking bytes
func GetRandomBytes(len int) ([]byte, error) {
	key := make([]byte, len)

	_, err := rand.Read(key)
	if err != nil {
		return nil, errors.New("error getting random bytes")
	}

	return key, nil
}

// GetRandomNonce returns a random byte array of length NonceSize
func GetRandomNonce() ([]byte, error) {
	return GetRandomBytes(NonceSize)
}

// GetRandomBigInt returns a random big int
func GetRandomBigInt() (*big.Int, error) {
	b, err := GetRandomBytes(32)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(b), nil

}

func ComputeTxnID(nonce, creator []byte) (string, error) {
	h := sha256.New()
	b := append(nonce, creator...)
	_, err := h.Write(b)
	if err != nil {
		return "", err
	}
	digest := h.Sum(nil)
	id := hex.EncodeToString(digest)
	return id, nil
}

// GetHash return sha256 hash
func GetHash(data []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	digest := h.Sum(nil)
	return digest, nil
}
