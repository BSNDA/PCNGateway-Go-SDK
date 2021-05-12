package config

import (
	"fmt"
	"testing"
)

func TestNewConfigFormFile(t *testing.T) {

	//path :="../../../conf/config.json"

	conf, err := NewConfigFormFile("")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(conf.user.UserCode)
}
