package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	types "e-commence/app/frontend/biz/types"
	common "e-commence/app/frontend/hertz_gen/frontend/common"
	"e-commence/app/frontend/infra/rpc"
	frontendUtils "e-commence/app/frontend/utils"
	"e-commence/rpc_gen/kitex_gen/order"
	"e-commence/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/common/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {

	user_id := frontendUtils.GetUserIdFromCtx(h.Context)
	rpcresp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{
		UserId: uint64(user_id),
	})

	if err != nil {
		return nil, err
	}
	// lists := make([]types.Order, 0)
	var lists []*types.Order
	for _, v := range rpcresp.Orders {
		var (
			total float32
			items []types.OrderItems
		)
		total = 0
		for _, vo := range v.Items {

			productInfo, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: vo.Item.ProductId})
			total += productInfo.Product.Price
			if err != nil {
				return nil, errors.New("GetProductInfo Error")
			}
			if productInfo == nil || productInfo.Product == nil {
				return nil, errors.New("ProductInfo Get Error")
			}
			productInfo.Product.Picture = "https://api.paugram.com/wallpaper/"
			items = append(items, types.OrderItems{
				ProductName: productInfo.Product.Name,
				Picture:     productInfo.Product.Picture,
				Qty:         uint32(vo.Item.Quantity),
				Cost:        vo.Cost * float32(vo.Item.Quantity),
			})
		}
		fmt.Println(v.CreateAt)
		createAt := time.Unix(int64(v.CreateAt), 0).Format(time.DateTime)

		lists = append(lists, &types.Order{
			Items:       items,
			OrderId:     v.OrderId,
			CreatedDate: createAt,
			Cost:        total,
			Qty:         uint32(len(items)),
		})

	}

	return utils.H{
		"Title":  "Order",
		"orders": lists,
	}, nil
}
