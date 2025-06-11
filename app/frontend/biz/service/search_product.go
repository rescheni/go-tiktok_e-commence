package service

import (
	"context"

	product "e-commence/app/frontend/hertz_gen/frontend/product"
	"e-commence/app/frontend/infra/rpc"
	rpcproduct "e-commence/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductService(Context context.Context, RequestContext *app.RequestContext) *SearchProductService {
	return &SearchProductService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductService) Run(req *product.SearchProductReq) (resp map[string]any, err error) {

	s, err := rpc.ProductClient.SearchProduct(h.Context, &rpcproduct.SearchProductReq{
		Query: req.Q,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range s.Products {
		v.Picture = "https://api.paugram.com/wallpaper/"
	}
	return utils.H{
		"Title": "Search",
		"items": s.Products,
	}, nil
}
