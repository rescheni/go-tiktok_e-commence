package rpc

import (
	"e-commence/common/clientsuite"
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/checkout/checkoutservice"
	"e-commence/rpc_gen/kitex_gen/order/orderservice"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"e-commence/rpc_gen/kitex_gen/user/userservice"
	frontendUtils "gomall/utils"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
)

// 注册中心
func Init() {
	once.Do(
		func() {
			iniUserClient()
			iniProductClient()
			iniCartClient()
			iniCheckoutClient()
			iniOrderClient()
		},
	)
}

var (
	serverName   = frontendUtils.ServerName_Frontend
	registryAddr = frontendUtils.RegistryAddr_Frontend
	err          error
)

func iniUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}

	UserClient, err = userservice.NewClient("user", opts...)

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

func iniCheckoutClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serverName,
			RegistryAddr:       registryAddr,
		}),
	}
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)

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
