package service

import (
	"context"

	rpcproduct "e-commence/rpc_gen/kitex_gen/product"

	product "gomall/hertz_gen/frontend/product"
	"gomall/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {

	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	p.Product.Picture = "https://api.paugram.com/wallpaper/"

	return utils.H{
		"item": p.Product,
	}, nil

}
