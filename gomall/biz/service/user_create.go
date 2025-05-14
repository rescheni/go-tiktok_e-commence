package service

import (
	"context"
	user "e-commence/gomall/kitex_gen/user"
)

type UserCreateService struct {
	ctx context.Context
} // NewUserCreateService new UserCreateService
func NewUserCreateService(ctx context.Context) *UserCreateService {
	return &UserCreateService{ctx: ctx}
}

// Run create note info
func (s *UserCreateService) Run(req *user.UserCreateReq) (resp *user.UserCreateResp, err error) {
	// Finish your business logic.
	return
}
