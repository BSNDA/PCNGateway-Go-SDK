package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

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

func GetHash(data []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(data)
	if err != nil {
		return nil, err
	}
	digest := h.Sum(nil)
	return digest, nil
}
