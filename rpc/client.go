package rpc

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth/authservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart/cartservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order/orderservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment/paymentservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product/productcatalogservice"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	consulconfig "github.com/kitex-contrib/config-consul/consul"
	"log"
	"tiktok_e-commence/common/clientsuite"
)

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "auth"}))
	if err != nil {
		log.Fatal(err)
	}
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "user"}))
	if err != nil {
		log.Fatal(err)
	}
}

func initProductClient() {
	//添加熔断
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("client/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)
	consulClient, _ := consulconfig.NewClient(consulconfig.Options{}) // consul配置中心

	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "product"}),
		client.WithCircuitBreaker(cbs),
		client.WithFallback( // 添加fallback，降级返回策略
			fallback.NewFallbackPolicy(
				fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
					if err != nil {
						return nil, err // 目前有错误也没有做返回策略
					}
					return resp, nil
				}),
			),
		),
		client.WithSuite(consulclient.NewSuite("product", "product_consul_config", consulClient)), // consul配置中心
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		log.Fatal(err)
	}
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "cart"}))
	if err != nil {
		log.Fatal(err)
	}
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "order"}))
	if err != nil {
		log.Fatal(err)
	}
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", client.WithSuite(clientsuite.CommonClientSuite{CurrentServiceName: "payment"}))
	if err != nil {
		log.Fatal(err)
	}
}
