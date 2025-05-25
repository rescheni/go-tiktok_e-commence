package checkout

import (
	"context"

	"gomall/biz/service"
	"gomall/biz/utils"
	checkout "gomall/hertz_gen/frontend/checkout"
	common "gomall/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Checkout .
// @router /Checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	c.HTML(consts.StatusOK, "checkout", utils.WarpRespose(ctx, c, resp))
}

// CheckoutWaiting .
// @router /Checkout/Waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "waiting", utils.WarpRespose(ctx, c, resp))
}

// CheckoutResult .
// @router /Checkout/Result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	c.HTML(consts.StatusOK, "result", utils.WarpRespose(ctx, c, resp))
}
