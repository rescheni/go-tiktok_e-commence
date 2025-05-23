package service

import (
	"context"
	"e-commence/app/cart/biz/dal/mysql"
	"e-commence/app/cart/biz/model"
	cart "e-commence/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

//TODO: 单元测试

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartreq) (resp *cart.EmptyCartresp, err error) {
	err = model.EmptyCart(s.ctx, mysql.DB, req.UserId)

	if err != nil {
		return nil, kerrors.NewBizStatusError(400001, err.Error())
	}

	return &cart.EmptyCartresp{}, nil
}
