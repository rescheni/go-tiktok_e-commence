package service

import (
	"context"
	"fmt"

	auth "gomall/hertz_gen/frontend/auth"

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

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 114514)
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
