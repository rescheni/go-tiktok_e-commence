package service

import (
	"context"
	"e-commence/app/product/biz/dal/mysql"
	"e-commence/app/product/biz/model"
	product "e-commence/rpc_gen/kitex_gen/product"
)

type SearchProductService struct {
	ctx context.Context
} // NewSearchProductService new SearchProductService
func NewSearchProductService(ctx context.Context) *SearchProductService {
	return &SearchProductService{ctx: ctx}
}

// Run create note info
func (s *SearchProductService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {

	productQuery := model.NewProductQuery(s.ctx, mysql.DB)
	p, err := productQuery.SearchProduct(req.Query)
	// fmt.Printf("%#v\n", req)
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
