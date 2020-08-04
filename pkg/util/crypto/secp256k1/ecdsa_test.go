package secp256k1

import (
	"fmt"
	"testing"
)

func TestNewSecp256k1Key(t *testing.T) {

	pk, err := NewSecp256k1Key()

	if err != nil {
		t.Fatal(err)
	}

	pkpem, err := PrivateKeyToPEM(pk)
	fmt.Println(string(pkpem))

	puk := pk.PublicKey

	pukpem, err := PublicKeyToPEM(&puk)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(pukpem))

}
