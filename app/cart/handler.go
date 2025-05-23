package main

import (
	"context"
	"e-commence/app/cart/biz/service"
	cart "e-commence/rpc_gen/kitex_gen/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemreq) (resp *cart.AddItemresp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartreq) (resp *cart.EmptyCartresp, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}
