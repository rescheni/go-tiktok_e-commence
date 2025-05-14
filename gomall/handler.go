package main

import (
	"context"
	"e-commence/gomall/biz/service"
	user "e-commence/gomall/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserCreate implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserCreate(ctx context.Context, req *user.UserCreateReq) (resp *user.UserCreateResp, err error) {
	resp, err = service.NewUserCreateService(ctx).Run(req)

	return resp, err
}
