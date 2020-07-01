package config

import (
	"fmt"
	"testing"
)

func TestConfigInit(t *testing.T) {

	config, err := NewMockConfig()

	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(*config)
	}

}
