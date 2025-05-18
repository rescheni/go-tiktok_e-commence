package service

import (
	"context"
	user "e-commence/rpc_gen/kitex_gen/user"
)

type RegirsterService struct {
	ctx context.Context
} // NewRegirsterService new RegirsterService
func NewRegirsterService(ctx context.Context) *RegirsterService {
	return &RegirsterService{ctx: ctx}
}

// Run create note info
func (s *RegirsterService) Run(req *user.RegirsterReq) (resp *user.RegirsterResp, err error) {
	// Finish your business logic.

	return
}
