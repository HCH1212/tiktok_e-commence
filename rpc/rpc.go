package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"log"
)

//func initAuthClient() {
//	AuthClient, err = authservice.NewClient("auth", client.WithResolver(common()))
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func initUserClient() {
//	UserClient, err = userservice.NewClient("user", client.WithResolver(common()))
//	if err != nil {
//		log.Fatal(err)
//	}
//}

// 暂时不使用consul
func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", client.WithHostPorts("localhost:8081"))
	if err != nil {
		log.Fatal(err)
	}
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithHostPorts("localhost:8082"))
	if err != nil {
		log.Fatal(err)
	}
}
