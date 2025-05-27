package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

// 服务名称 检测服务端口 注册中心地址
func IniMetric(serviceName, metricsPort, regirstryAddr string) (registry.Registry, *registry.Info) {

	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	r, _ := consul.NewConsulRegister(regirstryAddr)

	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	regirstryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags: map[string]string{
			"service": serviceName,
		},
	}
	_ = r.Register(regirstryInfo)

	server.RegisterShutdownHook(func() {
		r.Deregister(regirstryInfo)

	})
	// Prometheus 的 HTTP handler 负责 在你的服务中暴露 /metrics 接口。
	http.Handle("/metrices", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	go http.ListenAndServe(metricsPort, nil)

	return r, regirstryInfo

}
