package main

import (
	"fmt"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/client/fabric"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/config"
	"github.com/BSNDA/PCNGateway-Go-SDK/pkg/core/entity/req/fabric/user"
	"log"
)

func main() {

	fmt.Println("github.com/BSNDA/PCNGateway-Go-SDK")

	api := ""      //PCN gateway address
	userCode := "" //user code
	appCode := ""  //DApp code
	puk := ""      //public key
	prk := ""      //private key
	mspDir := ""   //cert storage directory
	cert := ""     //cert

	conf, err := config.NewConfig(api, userCode, appCode, puk, prk, mspDir, cert)
	if err != nil {
		log.Fatal(err)
	}

	client, err := fabric.InitFabricClient(conf)
	if err != nil {
		log.Fatal(err)
	}
	req := user.RegisterReqDataBody{
		Name:   "",
		Secret: "",
	}

	res, err := client.RegisterUser(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Header.Code != 0 {
		log.Fatal(res.Header.Msg)
	}

}
