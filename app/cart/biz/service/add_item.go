package service

import (
	"context"
	"e-commence/app/cart/biz/dal/mysql"
	"e-commence/app/cart/biz/model"
	"e-commence/app/cart/rpc"
	cart "e-commence/rpc_gen/kitex_gen/cart"
	"e-commence/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// TODO: 单元测试
// Run create note info
func (s *AddItemService) Run(req *cart.AddItemreq) (resp *cart.AddItemresp, err error) {

	// 逻辑传入商品id判断是否真实
	productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40001, "商品id错误/商品不存在")
	}

	err = model.AddItem(s.ctx, mysql.DB, model.Cart{UserId: req.UserId, ProductId: req.Item.ProductId, Qty: uint32(req.Item.Quantity)})

	if err != nil {
		kerrors.NewBizStatusError(50000, err.Error())
	}

	return &cart.AddItemresp{}, nil
}
