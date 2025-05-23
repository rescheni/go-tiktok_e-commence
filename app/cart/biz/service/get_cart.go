package service

import (
	"context"
	"e-commence/app/cart/biz/dal/mysql"
	"e-commence/app/cart/biz/model"
	cart "e-commence/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

//TODO: 单元测试

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	cartList, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)

	if err != nil {
		return nil, kerrors.NewBizStatusError(50033, err.Error())
	}

	resp = &cart.GetCartResp{}

	for _, v := range cartList {
		t := &cart.CartItem{
			ProductId: v.ProductId,
			Quantity:  uint64(v.Qty),
		}
		resp.Items = append(resp.Items, t)
	}
	return
}
