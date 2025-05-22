package service

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/product"

	category "gomall/hertz_gen/frontend/category"
	"gomall/infra/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {

	p, err := rpc.ProductClient.ListProduct(h.Context, &product.ListProductReq{
		CategoryName: req.Category,
	})

	if err != nil {
		return nil, err
	}
	for _, v := range p.Products {
		v.Picture = "https://api.paugram.com/wallpaper/"
	}

	return utils.H{
		"Title": "Category",
		"items": p.Products,
	}, nil
}
