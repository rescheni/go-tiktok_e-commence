package rpc

import (
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/order/orderservice"
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
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		iniCartClient()
		iniProductClient()
		iniPaymentClient()
		iniOrderClient()
	})
}

func getConsulServer() discovery.Resolver {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	return r
}

func iniCartClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniPaymentClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	PaymentClient, err = paymentservice.NewClient("payment", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniProductClient() {

	r := getConsulServer()
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniOrderClient() {
	r := getConsulServer()

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}

}
