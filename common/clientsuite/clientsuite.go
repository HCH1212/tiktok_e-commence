package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/spf13/viper"
	"log"
)

// 可包装options和middleware

type CommonClientSuite struct {
	CurrentServiceName string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: s.CurrentServiceName}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	r, err := consul.NewConsulResolver(viper.GetString("consul.addr"))
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
