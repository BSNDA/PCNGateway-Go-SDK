package sm2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
)

var (
	default_uid = []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
)

func Sign(data string) string {
	pk, err := sm2.ReadPrivateKeyFromPem("./cert/sm2PriKeyPkcs8.pem", nil)

	if err != nil {
		log.Fatal(err)
	}

	r, s, err := sm2.Sm2Sign(pk, []byte(data), default_uid)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("R:", r)
	fmt.Println("S:", s)

	var buffer bytes.Buffer
	buffer.Write(r.Bytes())
	buffer.Write(s.Bytes())

	signature := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return signature
}

func Verify(data, sign string) bool {

	pubKeyPath := "./cert/sm2PubKey.pem"

	pubK, err := sm2.ReadPublicKeyFromPem(pubKeyPath, nil)

	if err != nil {
		log.Fatal(err)
	}

	d64, err := base64.StdEncoding.DecodeString(sign)

	if err != nil {
		log.Fatal(err)
	}

	l := len(d64)
	br := d64[:l/2]
	bs := d64[l/2:]

	var ri, si big.Int
	r := ri.SetBytes(br)
	s := si.SetBytes(bs)

	fmt.Println("R:", r)
	fmt.Println("S:", s)

	v := sm2.Sm2Verify(pubK, []byte(data), default_uid, r, s)

	return v
}
