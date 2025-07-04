// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	product "e-commence/rpc_gen/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListProduct(ctx context.Context, Req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProduct(ctx context.Context, Req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductCataLogServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductCataLogServiceClient struct {
	*kClient
}

func (p *kProductCataLogServiceClient) ListProduct(ctx context.Context, Req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListProduct(ctx, Req)
}

func (p *kProductCataLogServiceClient) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kProductCataLogServiceClient) SearchProduct(ctx context.Context, Req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchProduct(ctx, Req)
}
