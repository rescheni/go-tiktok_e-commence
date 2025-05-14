package basic_test

import (
	"context"
	"e-commence/gomall/kitex_gen/user"
	"e-commence/gomall/kitex_gen/user/userservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"log"
	"os"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"

	consul "github.com/kitex-contrib/registry-consul"
)

type userService struct{}

func (s *userService) UserCreate(ctx context.Context, req *user.UserCreateReq) (resp *user.UserCreateResp, err error) {
	resp = &user.UserCreateResp{}
	resp.Token = req.UserName + req.Password + req.Email + "-Token"
	resp.Err = "nil"
	return resp, nil
	//panic("implement me")
}

//func (h *userService) UserCreate(ctx context.Context, req api.Request) (resp api.Response, err error) {
//	resp = api.Response{Message: req.Message}
//	return resp, nil
//}

// 服务注册
func TestRegisterConsul(t *testing.T) {
	register, err := consul.NewConsulRegister(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		return
	}
	svc := userservice.NewServer(
		new(userService),
		server.WithRegistry(register),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "userservice",
			Weight:      1,
		}),
	)
	err = svc.Run()
	if err != nil {
		return
	}
}

// 服务发现
func TestDiscover(t *testing.T) {
	r, err := consul.NewConsulResolver(os.Getenv("GOMALL_CONSUL_URL") + ":" + os.Getenv("GOMALL_CONSUL_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("--------------")

	c := userservice.MustNewClient("userservice", client.WithResolver(r), client.WithRPCTimeout(3*time.Second))
	cxt := context.Background()
	//for {
	resp, err := c.UserCreate(cxt, &user.UserCreateReq{
		UserName: "username",
		Password: "password",
		Email:    "email@test.com",
	})

	if err != nil {
		t.Fatal(err)
		return
	}

	log.Println(resp.Token, resp.Err)
	//}
}
