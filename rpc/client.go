package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/discovery"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spf13/viper"
	"log"
)

var (
	err        error
	AuthClient authservice.Client
	UserClient userservice.Client
)

func InitClient() {
	initAuthClient()
	initUserClient()
}

func common() discovery.Resolver {
	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}
	return r
}
