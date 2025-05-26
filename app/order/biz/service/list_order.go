package service

import (
	"context"
	"e-commence/app/order/biz/dal/mysql"
	"e-commence/app/order/biz/model"
	"e-commence/rpc_gen/kitex_gen/cart"
	order "e-commence/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {

	lists, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)

	resp = &order.ListOrderResp{}

	for _, v := range lists {

		var items []*order.OrderItem
		for _, ov := range v.OrderItem {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: ov.ProductId,
					Quantity:  uint64(ov.Quantity),
				},
				Cost: float32(ov.Cost),
			})
		}

		resp.Orders = append(resp.Orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       uint32(v.UserId),
			CreateAt:     int32(v.CreatedAt.Unix()),
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				City:          v.Consignee.Ciry,
				State:         v.Consignee.State,
				Country:       v.Consignee.Country,
				ZipCode:       v.Consignee.ZipCode,
				StreetAddress: v.Consignee.StreetAddress,
			},
			Items: items,
		})
	}

	return
}
