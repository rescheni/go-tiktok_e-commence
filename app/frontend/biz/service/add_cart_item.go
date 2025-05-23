package service

import (
	"context"

	rpccart "e-commence/rpc_gen/kitex_gen/cart"
	cart "gomall/hertz_gen/frontend/cart"
	common "gomall/hertz_gen/frontend/common"
	"gomall/infra/rpc"
	frontendUtils "gomall/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {

	_, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemreq{
		UserId: uint64(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: req.ProductId,
			Quantity:  uint64(req.ProductNum),
		},
	})

	return
}
