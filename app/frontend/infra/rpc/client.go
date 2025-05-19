package rpc

import (
	"e-commence/rpc_gen/kitex_gen/user/userservice"
	"os"
	"sync"

	frontendUtil "gomall/utils"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(
		func() {
			iniUserClient()
		},
	)
}
func iniUserClient() {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		klog.Fatal("Error creating consul resolver")
	}
	frontendUtil.MustHandleError(err)

	UserClient, err = userservice.NewClient("user", client.WithResolver(r))

	if err != nil {
		klog.Fatal("Error creating user client")
	}
}
