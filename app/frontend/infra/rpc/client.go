package rpc

import (
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"e-commence/rpc_gen/kitex_gen/user/userservice"
	"os"
	"sync"

	frontendUtil "gomall/utils"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

// 注册中心
func Init() {
	once.Do(
		func() {
			iniUserClient()
			iniProductClient()
			iniCartClient()
		},
	)
}
func iniUserClient() {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))

	var opts []client.Option

	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	frontendUtil.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
func iniProductClient() {

	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniCartClient() {
	//TODO:整合服务注册工具函数
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
