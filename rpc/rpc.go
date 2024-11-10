package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spf13/viper"
	"log"
)

func initAuthClient() {
	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}
	AuthClient, err = authservice.NewClient("auth", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}

func initUserClient() {
	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
}
