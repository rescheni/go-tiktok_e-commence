package service

import (
	"context"
	"testing"
	product "e-commence/rpc_gen/kitex_gen/product"
)

func TestSearchProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSearchProductService(ctx)
	// init req and assert value

	req := &product.SearchProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
