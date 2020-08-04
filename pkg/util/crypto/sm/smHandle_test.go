package sm

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSm2Handle_Sign(t *testing.T) {
	//	puk := `-----BEGIN PUBLIC KEY-----
	//MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEDW9srwJ97PuwNTXKpCBLz+Kgp8Bo
	//KS/i2zlbzA3gnrZPKjh8jfh++exUmliaJ1qlzeNeXHyEbV31Rqk4+Go3Tw==
	//-----END PUBLIC KEY-----`

	prik := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgYy1PjNmYLJrs1YL4
7xeGSBtc7bHii4fcQBoX+zOkgBigCgYIKoEcz1UBgi2hRANCAAQNb2yvAn3s+7A1
NcqkIEvP4qCnwGgpL+LbOVvMDeCetk8qOHyN+H757FSaWJonWqXN415cfIRtXfVG
qTj4ajdP
-----END PRIVATE KEY-----`

	sm, err := NewSM2Handle("", prik)

	if err != nil {
		fmt.Println(err)
	}

	data := []byte("123456")

	si, _ := sm.Sign(data)

	b, err := sm.Verify(si, data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)
}

func TestEncrypt(t *testing.T) {

	puk := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAECBTmBCyjjyg0h4F1f/PiLVNJyDM1
YRgctLay3FE5wWLqG0OH0p5fP8I5UT+pb1gkirIchlDxuwdVdVlUuQMrTQ==
-----END PUBLIC KEY-----`

	prik := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgDvuumI/YLlcYhtO0
klhtbxYHH8Mk8clOEtDjG3obhUugCgYIKoEcz1UBgi2hRANCAAQIFOYELKOPKDSH
gXV/8+ItU0nIMzVhGBy0trLcUTnBYuobQ4fSnl8/wjlRP6lvWCSKshyGUPG7B1V1
WVS5AytN
-----END PRIVATE KEY-----`

	sm, err := NewSM2Handle(puk, prik)
	if err != nil {
		t.Fatal(err)
	}

	data := []byte("abc")

	cr, err := sm.Encrypt(data)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Encrypt：", hex.EncodeToString(cr))

	data, err = sm.Decrypt(cr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Decrypt：", string(data))

}
