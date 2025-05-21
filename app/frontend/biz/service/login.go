package service

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/user"
	"fmt"

	auth "gomall/hertz_gen/frontend/auth"
	"gomall/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	resp, err := rpc.UserClient.Login(
		h.Context, &user.LoginReq{
			Username: req.Username,
			Password: req.Password,
		},
		// callopt.WithConnectTimeout(1000),
		// callopt.WithRPCTimeout(1000),
	)
	if err != nil {
		return " ", err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	session.Set("user_name", resp.Username)
	err = session.Save()

	if err != nil {
		return "Session err", err
	}
	// if req.Username == "" || req.Password == "" {
	// 	return "Username or password is empty", nil
	// }
	redirect = req.Next
	fmt.Println("req.Next = ", req.Next)
	if req.Next == "" {
		redirect = "/"
	}

	return
}
