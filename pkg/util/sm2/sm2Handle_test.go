package sm2

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSign2(t *testing.T) {
	puk := `-----BEGIN CERTIFICATE-----
MIICuzCCAmGgAwIBAgIIcB9gqgFxReQwCgYIKoZIzj0EAwIwbzELMAkGA1UEBhMC
Y24xEDAOBgNVBAgMB2JlaWppbmcxEDAOBgNVBAcMB2JlaWppbmcxDDAKBgNVBAoM
A2JzbjEMMAoGA1UECwwDYnNuMSAwHgYDVQQDDBdjYS5yb290LnNtMi5ic25iYXNl
LmNvbTAgFw0yMDA0MTAwODUzMTFaGA8yMTIwMDMxNzA4NTMxMVowajELMAkGA1UE
BhMCY24xEDAOBgNVBAgMB2JlaWppbmcxEDAOBgNVBAcMB2JlaWppbmcxDDAKBgNV
BAoMA2JzbjEMMAoGA1UECwwDYnNuMRswGQYDVQQDDBJjYS5zbTIuYnNuYmFzZS5j
b20wWTATBgcqhkjOPQIBBggqgRzPVQGCLQNCAATrsJHrxZjIDCbeWpriqiQac/kV
QHTK4uS2kfOV16xn26sVc5mtKpYOGovG4XcQLAGu5CHC9nIrmWBlaZ7L2u2Lo4Hp
MIHmMBIGA1UdEwEB/wQIMAYBAf8CAQEwDgYDVR0PAQH/BAQDAgEGMIGgBgNVHSME
gZgwgZWAFBXKKylVPGQ45yj7y0x4vJcWb+NRoXOkcTBvMQswCQYDVQQGEwJjbjEQ
MA4GA1UECAwHYmVpamluZzEQMA4GA1UEBwwHYmVpamluZzEMMAoGA1UECgwDYnNu
MQwwCgYDVQQLDANic24xIDAeBgNVBAMMF2NhLnJvb3Quc20yLmJzbmJhc2UuY29t
ggg/c4E8vgNHvzAdBgNVHQ4EFgQUd6JSbzAx+oXI/AEcElx3r/b7EIcwCgYIKoZI
zj0EAwIDSAAwRQIhAIIWLvrLTSfTY/NrlrejPy+NklqK+R+5BDvUSt42XGZiAiAt
+uGk4driNal+/gtfgRApoLqdfJfZf0LJrZgohsY2nQ==
-----END CERTIFICATE-----`

	prik := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBG0wawIBAQQgkkH5AWNBkkuebFaX
DG34KWvxPjeWLRfjHTzflftNf8mhRANCAASrHJOXcdm6M7oO3klzMalBL7YOVVm0
yfxyHfL37wEwJUNCaJCrZK8gWmADtfPDu331TuTnwPmO+EZ6HAeQ9fxm
-----END PRIVATE KEY-----`

	sm, err := NewSM2Handle(puk, prik)

	if err != nil {
		fmt.Println(err)
	}

	str := "abc"

	data := []byte(str)
	//sign :="MEUCIAPzJzJ/2/tcTQS864NCHPM1SWxiWVVgQCpeLMP5fneiAiEArRbdm5pRfxqOn51mlDB4ThO5zUjiAxJyztnP6BRJu+M="

	s, _ := sm.Sign(data)

	sign := base64.StdEncoding.EncodeToString(s)

	fmt.Println(sign)

	b, err := sm.Verify(s, data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

}

func TestSign3(t *testing.T) {
	puk := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoEcz1UBgi0DQgAEDW9srwJ97PuwNTXKpCBLz+Kgp8Bo
KS/i2zlbzA3gnrZPKjh8jfh++exUmliaJ1qlzeNeXHyEbV31Rqk4+Go3Tw==
-----END PUBLIC KEY-----`

	prik := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgYy1PjNmYLJrs1YL4
7xeGSBtc7bHii4fcQBoX+zOkgBigCgYIKoEcz1UBgi2hRANCAAQNb2yvAn3s+7A1
NcqkIEvP4qCnwGgpL+LbOVvMDeCetk8qOHyN+H757FSaWJonWqXN415cfIRtXfVG
qTj4ajdP
-----END PRIVATE KEY-----`

	sm, err := NewSM2Handle(puk, prik)

	if err != nil {
		fmt.Println(err)
	}

	data := []byte("123456")
	sign := "MEQCIDUGejY4XbMg7AC/S8BtdETFZInPwr//QypbDsc33HCAAiAK2rX+c3VOTPrFdZTr/iE/Vz3n9CSDC3w3tLbbL8anAA=="

	si, _ := base64.StdEncoding.DecodeString(sign)

	//si,_ := sm.Sign(data)

	//sign :=base64.StdEncoding.EncodeToString(si)

	//fmt.Println("sign:",sign)

	b, err := sm.Verify(si, data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b)

}

func TestSign4(t *testing.T) {

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
		fmt.Println(err)
	}
	str := "123456"
	data := []byte(str)
	digest, _ := sm.Hash(data)
	fmt.Println("Hash:", digest)
	fmt.Println(base64.StdEncoding.EncodeToString(digest))

	sign := "MEUCIQDQKtx0Heiyj1DddEmkueb8+dAzdG6llDaG+Qg6KiGhzQIgRZJoCfZDefI+UhpaTMO35wE4BooQ6dhy99peMoGvoWc="

	s, _ := base64.StdEncoding.DecodeString(sign)

	//s, _ := sm.Sign(digest)
	//sign := base64.StdEncoding.EncodeToString(s)

	fmt.Println(sign)

	b, err := sm.Verify(s, digest)

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
