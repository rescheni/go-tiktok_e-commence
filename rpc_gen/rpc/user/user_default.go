package user

import (
	"context"
	user "e-commence/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Regirster(ctx context.Context, req *user.RegirsterReq, callOptions ...callopt.Option) (resp *user.RegirsterResp, err error) {
	resp, err = defaultClient.Regirster(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Regirster call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
