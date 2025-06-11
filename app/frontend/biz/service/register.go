package service

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/user"

	auth "e-commence/app/frontend/hertz_gen/frontend/auth"
	common "e-commence/app/frontend/hertz_gen/frontend/common"
	"e-commence/app/frontend/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	userResp, err := rpc.UserClient.Regirster(h.Context, &user.RegirsterReq{
		Username:      req.Username,
		Password:      req.Password,
		Checkpassword: req.CheckPassword,
		Email:         req.Email,
	})
	if err != nil {
		return nil, err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()

	if err != nil {
		return nil, err
	}

	return
}
