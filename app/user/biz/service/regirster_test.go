package service

import (
	"context"
	"testing"
	user "e-commence/rpc_gen/kitex_gen/user"
)

func TestRegirster_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegirsterService(ctx)
	// init req and assert value

	req := &user.RegirsterReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
