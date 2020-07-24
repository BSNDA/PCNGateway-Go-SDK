package secp256r1

import (
	"fmt"
	"testing"
)

func TestEcdsaHandle_Sign(t *testing.T) {
	//	puk := `-----BEGIN PUBLIC KEY-----
	//MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHxaNHI1l5vRXtRf7ynmGTPucSBgG
	//vKYL29fnPrPG0asPvcVr0qEL/ALj9kxSXuIk8WSRNbS8R1bERHx9Jjh9aQ==
	//-----END PUBLIC KEY-----`

	priker := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcRniHqapLZ4dwFpJ
Zo3ExKJfnRrYlOzHtLgWYEtiOy2hRANCAAQfFo0cjWXm9Fe1F/vKeYZM+5xIGAa8
pgvb1+c+s8bRqw+9xWvSoQv8AuP2TFJe4iTxZJE1tLxHVsREfH0mOH1p
-----END PRIVATE KEY-----`

	ec, err := NewEcdsaR1Handle("", priker)

	if err != nil {
		fmt.Println(err)
	}

	data := []byte("123456")

	si, _ := ec.Sign(data)

	b, err := ec.Verify(si, data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

}
