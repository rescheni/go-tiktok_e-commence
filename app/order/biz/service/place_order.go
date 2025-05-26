package service

import (
	"context"
	mysql "e-commence/app/order/biz/dal/mysql"
	"e-commence/app/order/biz/model"
	order "e-commence/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	if len(req.Items) == 0 {
		//没有商品
		return nil, kerrors.NewGRPCBizStatusError(50000, "The ReqItem Is Empty")
	}
	// 事务开启

	err = mysql.DB.Transaction(func(tx *gorm.DB) error {
		// 事务处理
		uuid, _ := uuid.NewUUID()

		order_id := uuid.String()

		if req.Address == nil {
			return kerrors.NewGRPCBizStatusError(50002, "Req Address Is Empty")
		}

		o := &model.Order{
			OrderId:      order_id,
			UserId:       uint64(req.UserId),
			UserCurrency: req.UserCurrency,
			Consignee: model.Consignee{
				Email:         req.Email,
				StreetAddress: req.Address.StreetAddress,
				Ciry:          req.Address.City,
				State:         req.Address.State,
				Country:       req.Address.Country,
				ZipCode:       req.Address.ZipCode,
			},
		}
		err := tx.Create(o).Error
		if err != nil {
			return kerrors.NewGRPCBizStatusError(50003, "Transaction Order Insert Err")
		}

		items := make([]model.OrderItem, 0)

		// var items []model.OrderItem

		for _, v := range req.Items {
			items = append(items, model.OrderItem{
				ProductId:    v.Item.ProductId,
				OrderIdRefer: order_id,
				Quantity:     uint32(v.Item.Quantity),
				Cost:         float64(v.Cost),
			})
		}

		err = tx.Create(items).Error
		if err != nil {
			return kerrors.NewGRPCBizStatusError(50003, "Transaction OrderItem Insert Err")
		}
		resp = &order.PlaceOrderResp{
			Order: &order.OrderResult{
				OrderId: order_id,
			},
		}
		return nil
	})

	return
}
