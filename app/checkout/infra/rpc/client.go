package rpc

import (
	"e-commence/app/checkout/conf"
	"e-commence/common/clientsuite"
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/order/orderservice"
	"e-commence/rpc_gen/kitex_gen/payment/paymentservice"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
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
var (
	serverName   = conf.GetConf().Kitex.Service
	registryAddr = os.Getenv(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
)

func InitClient() {
	once.Do(func() {
		iniCartClient()
		iniProductClient()
		iniPaymentClient()
		iniOrderClient()
	})
}

func iniCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniPaymentClient() {

	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Fatal("Error creating user client")
	}

}
