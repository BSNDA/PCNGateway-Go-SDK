package secp256r1

import (
	"fmt"
	"testing"
)

func TestEcdsaHandle_Sign(t *testing.T) {
	puk := `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEy4wBr/o5jSJHopiBfe9rhPhn//y
+Qf35AH4wwa92AjxLuhk28GlzOK7YiB5BitgttSlk+wLgTlEPF9m18cAvw==
-----END PUBLIC KEY-----`

	priker := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgxjJR06JfqaHDcpIN
JQb+MH0Bs2nWIhRUFJ3P4fIA8kCgCgYIKoZIzj0DAQehRANCAATko6mtCruC7pLI
MOZ4ktl9J2Lg5uQKx4fLIqT2oSZiFsZRhoMnaKmUfAPcYy3zaVmTtRkddHnTi0EC
V/xD6Mpe
-----END PRIVATE KEY-----`

	ec, err := NewEcdsaR1Handle(puk, priker)

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
