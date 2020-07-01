package ecdsa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestEcdsaHandle_Sign(t *testing.T) {
	puk := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHxaNHI1l5vRXtRf7ynmGTPucSBgG
vKYL29fnPrPG0asPvcVr0qEL/ALj9kxSXuIk8WSRNbS8R1bERHx9Jjh9aQ==
-----END PUBLIC KEY-----`

	priker := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgcRniHqapLZ4dwFpJ
Zo3ExKJfnRrYlOzHtLgWYEtiOy2hRANCAAQfFo0cjWXm9Fe1F/vKeYZM+5xIGAa8
pgvb1+c+s8bRqw+9xWvSoQv8AuP2TFJe4iTxZJE1tLxHVsREfH0mOH1p
-----END PRIVATE KEY-----`

	ec, err := NewEcdsaR1Handle(puk, priker)

	if err != nil {
		fmt.Println(err)
	}

	str := `USER0001202004151958010871292app0001202004161020152918451test8888123456`
	//     USER0001202004161009309407413app0001202004161017141233920QU5oOUxvZW5JQS4uOXFkTTB8YzNLT3FBcc_app0001202004161017141233920_00set{"baseKey":"test20200421","baseValue":"this is string "}testtesttesttesttest

	data := []byte(str)
	digest, _ := ec.Hash(data)
	sign := "MEQCID1GWAsqfWt8/rDMKIIkg+xigE3GEfBb6AveYT2dKMSjAiAls4YaxNDBZ6K0vXgUvMWdcidnmdhbKCqw7NfM9JgFww=="

	si, _ := base64.StdEncoding.DecodeString(sign)

	//si, _ := ec.Sign(digest)

	//sign := base64.StdEncoding.EncodeToString(si)

	fmt.Println("sign:", sign)

	b, err := ec.Verify(si, digest)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

}
