package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart/cartservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order/orderservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment/paymentservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spf13/viper"
	"log"
)

func common() discovery.Resolver {
	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}
	return r
}

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

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(common()))
	if err != nil {
		log.Fatal(err)
	}
}
