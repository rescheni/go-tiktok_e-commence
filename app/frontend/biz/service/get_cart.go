package service

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/cart"
	"e-commence/rpc_gen/kitex_gen/product"
	"strconv"

	common "gomall/hertz_gen/frontend/common"
	"gomall/infra/rpc"
	frontendUtils "gomall/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

// 获取购物车列表
func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {

	cartrResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint64(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}

	var Cart_Sum float64

	var items []map[string]string

	Cart_Sum = 0
	for _, v := range cartrResp.Items {
		productInfo, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		Cart_Sum += float64(v.Quantity) * float64(productInfo.Product.Price)
		items = append(items, map[string]string{
			"Name":        productInfo.Product.Name,
			"Description": productInfo.Product.Description,
			// "Picture":     productInfo.Product.Picture,
			"Picture": "https://api.paugram.com/wallpaper/",
			"Price":   strconv.FormatFloat(float64(productInfo.Product.Price), 'f', 2, 64),
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
	}
	// resp["Items"] = items
	// resp["Title"] = "Cart"

	return utils.H{
		"Title": "Cart",
		"Items": items,
		"Total": strconv.FormatFloat(float64(Cart_Sum), 'f', 2, 64),
	}, nil
}
