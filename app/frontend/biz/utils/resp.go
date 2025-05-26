package utils

import (
	"context"
	"e-commence/rpc_gen/kitex_gen/cart"
	"gomall/infra/rpc"
	frontendUtil "gomall/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpRespose(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// todo edit custom code

	userId := frontendUtil.GetUserIdFromCtx(ctx)

	content["user_id"] = userId

	// 获取购物车角标
	if userId > 0 {
		cartResp, err := rpc.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: uint64(frontendUtil.GetUserIdFromCtx(ctx)),
		})
		if err != nil {
			return nil
		}
		if cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}
	return content
}
