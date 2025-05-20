package service

import (
	"context"
	"testing"
	product "e-commence/rpc_gen/kitex_gen/product"
)

func TestListProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductService(ctx)
	// init req and assert value

	req := &product.ListProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
