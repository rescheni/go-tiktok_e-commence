package mtl

import (
	"fmt"
	"net"
	"net/http"
	"os"

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
	hostname, _ := os.Hostname()
	port := metricsPort
	if len(port) > 0 && port[0] == ':' {
		port = port[1:]
	}
	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", hostname, port))
	regirstryInfo := &registry.Info{
		// ServiceName: "prometheus",
		ServiceName: serviceName,
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
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	go http.ListenAndServe(metricsPort, nil)

	return r, regirstryInfo

}
