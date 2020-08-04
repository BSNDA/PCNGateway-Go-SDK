package secp256k1

import (
	"fmt"
	"testing"
)

func TestEcdsaHandle_Sign(t *testing.T) {
	//	puk := `-----BEGIN PUBLIC KEY-----
	//MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAENVH4W8jn2rCJUoH9eW7kopdaRoQ2/Rab
	//sk925gFXUyRaSDAQk9Ih5XR/fu3puF9ol+7WrVT3L+JKtwQ5ES6ieA==
	//-----END PUBLIC KEY-----`

	priker := `-----BEGIN PRIVATE KEY-----
MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgdEteON4eEPho9GVeDmi9
Plw7HzI7EGUB2FmbLtbgbPChRANCAAQ1UfhbyOfasIlSgf15buSil1pGhDb9Fpuy
T3bmAVdTJFpIMBCT0iHldH9+7em4X2iX7tatVPcv4kq3BDkRLqJ4
-----END PRIVATE KEY-----`

	ec, err := NewEcdsaK1Handle("", priker)

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
