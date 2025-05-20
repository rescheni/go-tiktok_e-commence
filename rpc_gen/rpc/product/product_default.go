package product

import (
	"context"
	product "e-commence/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProduct(ctx context.Context, req *product.ListProductReq, callOptions ...callopt.Option) (resp *product.ListProductResp, err error) {
	resp, err = defaultClient.ListProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SerachProduct(ctx context.Context, req *product.SearchProductReq, callOptions ...callopt.Option) (resp *product.SearchProductResp, err error) {
	resp, err = defaultClient.SerachProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SerachProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
