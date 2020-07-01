package crypto

import (
	"fmt"
	"testing"
)

func TestComputeTxnID(t *testing.T) {

	n := "abc"
	c := "123"

	id, _ := ComputeTxnID([]byte(n), []byte(c))

	fmt.Println(id)
}
