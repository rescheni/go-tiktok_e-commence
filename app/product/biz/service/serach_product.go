package service

import (
	"context"
	"e-commence/app/product/biz/model"
	"e-commence/gomall/biz/dal/mysql"
	product "e-commence/rpc_gen/kitex_gen/product"
)

type SerachProductService struct {
	ctx context.Context
} // NewSerachProductService new SerachProductService
func NewSerachProductService(ctx context.Context) *SerachProductService {
	return &SerachProductService{ctx: ctx}
}

// Run create note info
func (s *SerachProductService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.SearchProduct(req.Query)
	cnt := int32(len(p))
	resp = &product.SearchProductResp{}
	for _, v := range p {
		resp.Products = append(resp.Products, &product.Product{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Picture:     v.Picture,
			Price:       v.Price,
		})
	}
	resp.Total = cnt
	return
}
