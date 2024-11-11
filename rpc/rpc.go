package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"log"
)

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}
