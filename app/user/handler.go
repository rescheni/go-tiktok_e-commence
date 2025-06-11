package main

import (
	"context"
	"e-commence/app/user/biz/service"
	"e-commence/rpc_gen/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Regirster implements the UserServiceImpl interface.
func (s *UserServiceImpl) Regirster(ctx context.Context, req *user.RegirsterReq) (resp *user.RegirsterResp, err error) {
	resp, err = service.NewRegirsterService(ctx).Run(req)

	return resp, err
}
