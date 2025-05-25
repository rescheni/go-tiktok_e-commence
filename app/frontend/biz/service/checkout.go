package service

import (
	"context"
	rpccart "e-commence/rpc_gen/kitex_gen/cart"
	"e-commence/rpc_gen/kitex_gen/product"
	"strconv"

	common "gomall/hertz_gen/frontend/common"
	"gomall/infra/rpc"

	"github.com/cloudwego/hertz/pkg/common/utils"

	frontendUtils "gomall/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {

	var items []map[string]string

	user_id := frontendUtils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{
		UserId: uint64(user_id),
	})
	var total float32

	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp == nil {
			continue
		}
		total += float32(v.Quantity) * productResp.Product.Price

		items = append(items, map[string]string{
			"Name":    productResp.Product.Name,
			"Price":   strconv.FormatFloat(float64(productResp.Product.Price), 'f', 2, 64),
			"Picture": productResp.Product.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
	}
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title": "checkout",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
