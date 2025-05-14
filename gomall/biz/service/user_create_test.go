package service

import (
	"context"
	"testing"
	user "e-commence/gomall/kitex_gen/user"
)

func TestUserCreate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUserCreateService(ctx)
	// init req and assert value

	req := &user.UserCreateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
