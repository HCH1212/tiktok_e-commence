package rpc

import (
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart/cartservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order/orderservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment/paymentservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
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

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithHostPorts("localhost:8083"))
	if err != nil {
		log.Fatal(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithHostPorts("localhost:8084"))
	if err != nil {
		log.Fatal(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithHostPorts("localhost:8085"))
	if err != nil {
		log.Fatal(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", client.WithHostPorts("localhost:8086"))
	if err != nil {
		log.Fatal(err)
	}
}
