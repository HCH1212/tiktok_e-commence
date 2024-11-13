package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"sync"
)

var (
	err           error
	one           sync.Once
	AuthClient    authservice.Client
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
)

func InitClient() {
	one.Do(func() {
		initAuthClient()
		initUserClient()
		initProductClient()
	})
}

//func common() discovery.Resolver {
//	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return r
//}
