package service

import (
	"context"
	"e-commence/app/product/biz/dal/mysql"
	"e-commence/app/product/biz/model"
	product "e-commence/rpc_gen/kitex_gen/product"
)

type ListProductService struct {
	ctx context.Context
} // NewListProductService new ListProductService
func NewListProductService(ctx context.Context) *ListProductService {
	return &ListProductService{ctx: ctx}
}

// Run create note info
func (s *ListProductService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {

	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := categoryQuery.GetProductByCategoryQuery(req.CategoryName)

	if err != nil {
		return nil, err
	}

	resp = &product.ListProductResp{}

	cnt := int32(0)

	for _, v1 := range c {
		for _, v := range v1.Product {
			resp.Products = append(resp.Products, &product.Product{
				Id:          v.ID,
				Name:        v.Name,
				Description: v.Description,
				Picture:     v.Picture,
				Price:       v.Price,
			})
			cnt++
		}
	}

	resp.Total = cnt
	return
}
