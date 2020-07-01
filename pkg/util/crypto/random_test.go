package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestGetRandomNonce(t *testing.T) {

	for i := 0; i < 1000; i++ {
		nonce, _ := GetRandomNonce()

		fmt.Println("nonce:", base64.StdEncoding.EncodeToString(nonce))

	}

}
