package service

import (
	"context"

	auth "gomall/hertz_gen/frontend/auth"
	common "gomall/hertz_gen/frontend/common"

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
	// TODO user svc api
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 114514)
	err = session.Save()

	if err != nil {
		return nil, err
	}

	return
}
