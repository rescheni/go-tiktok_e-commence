package rpc

import (
	"e-commence/common/clientsuite"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"e-commence/rpc_gen/kitex_gen/user/userservice"
	"os"
	"sync"

	"e-commence/app/cart/conf"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	once          sync.Once
	err           error
)

var (
	serverName   = conf.GetConf().Kitex.Service
	registryAddr = os.Getenv(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
)

func Init() {
	once.Do(
		func() {
			iniUserClient()
			iniProductClient()
		},
	)
}

func iniUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serverName,
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
			RegistryAddr:       registryAddr,
			CurrentServiceName: serverName,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
