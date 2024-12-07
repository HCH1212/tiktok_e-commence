package mtl

import (
	"github.com/cloudwego/kitex/pkg/registry"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var Registry *prometheus.Registry

func InitMetric(serviceName, MetricsPort, RegistryAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(prometheus.NewGoCollector())
	Registry.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	// 注册到consul
	r, _ := consul.NewConsulRegister(RegistryAddr)
	addr, _ := net.ResolveTCPAddr("tcp", MetricsPort)
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}
	_ = r.Register(registryInfo)
	//server.RegisterShutdownHook(func() {
	//	_ = r.Deregister(registryInfo)
	//})

	// 启动独立的 Prometheus HTTP 服务
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
		_ = http.ListenAndServe(MetricsPort, mux)
	}()

	return r, registryInfo
}
