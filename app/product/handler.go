package main

import (
	"context"
	"e-commence/app/product/biz/service"
	product "e-commence/rpc_gen/kitex_gen/product"
)

// ProductCataLogServiceImpl implements the last service interface defined in the IDL.
type ProductCataLogServiceImpl struct{}

// ListProduct implements the ProductCataLogServiceImpl interface.
func (s *ProductCataLogServiceImpl) ListProduct(ctx context.Context, req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	resp, err = service.NewListProductService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCataLogServiceImpl interface.
func (s *ProductCataLogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SerachProduct implements the ProductCataLogServiceImpl interface.
func (s *ProductCataLogServiceImpl) SerachProduct(ctx context.Context, req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	resp, err = service.NewSerachProductService(ctx).Run(req)

	return resp, err
}
