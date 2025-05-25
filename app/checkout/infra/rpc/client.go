package rpc

import (
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/payment/paymentservice"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"os"
	"sync"

	"github.com/cloudwego/kitex/pkg/discovery"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"k8s.io/klog/v2"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		InitCartClient()
		InitProductClient()
		InitPaymentClient()

	})
}

func getConsulServer() discovery.Resolver {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	return r
}

func InitCartClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func InitPaymentClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func InitProductClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
