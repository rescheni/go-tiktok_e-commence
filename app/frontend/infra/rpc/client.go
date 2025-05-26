package rpc

import (
	"e-commence/rpc_gen/kitex_gen/cart/cartservice"
	"e-commence/rpc_gen/kitex_gen/checkout/checkoutservice"
	"e-commence/rpc_gen/kitex_gen/order/orderservice"
	"e-commence/rpc_gen/kitex_gen/product/productcatalogservice"
	"e-commence/rpc_gen/kitex_gen/user/userservice"
	"os"
	"sync"

	frontendUtil "gomall/utils"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
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

func consulInit() (discovery.Resolver, error) {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	return r, err
}

func iniUserClient() {
	r, err := consulInit()
	frontendUtil.MustHandleError(err)

	var opts []client.Option

	opts = append(opts, client.WithResolver(r))

	UserClient, err = userservice.NewClient("user", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
func iniProductClient() {
	r, err := consulInit()
	frontendUtil.MustHandleError(err)

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniCartClient() {
	r, err := consulInit()
	frontendUtil.MustHandleError(err)

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniCheckoutClient() {
	r, err := consulInit()
	frontendUtil.MustHandleError(err)

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}

func iniOrderClient() {
	r, err := consulInit()
	frontendUtil.MustHandleError(err)

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)

	if err != nil {
		klog.Fatal("Error creating user client")
	}

}
