package account

import "testing"

func TestGetAddressFromPublicKey(t *testing.T) {
	publicKey := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAELkdua4oPfiLIq6u3mqcUrAwqPBQn
bdKnwmu3hgguwRHRnq21+KmuY96pY5df3eiNNHO73hQL4Bjz/AxU/pne6w==
-----END PUBLIC KEY-----`
	puk, err := GetAddressFromPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("address:%s", puk)
}

func TestGetEcdsaPublicKeyJsonFormatFromPublicKey(t *testing.T) {
	publicKey := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAELkdua4oPfiLIq6u3mqcUrAwqPBQn
bdKnwmu3hgguwRHRnq21+KmuY96pY5df3eiNNHO73hQL4Bjz/AxU/pne6w==
-----END PUBLIC KEY-----`
	puk, err := GetEcdsaPublicKeyJsonFormatFromPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("jsonFormat:%s", puk)
}
