package service

import (
	"context"
	"fmt"

	checkout "e-commence/app/frontend/hertz_gen/frontend/checkout"
	"e-commence/app/frontend/infra/rpc"
	frondendUtils "e-commence/app/frontend/utils"
	rpcchenkout "e-commence/rpc_gen/kitex_gen/checkout"
	"e-commence/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	user_id := frondendUtils.GetUserIdFromCtx(h.Context)

	fmt.Println(req)
	checkoutRpcReq := &rpcchenkout.CheckoutReq{
		UserId:    uint32(user_id),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpcchenkout.Address{
			StreetAddress: req.Street,
			ZipCode:       req.Zipcode,
			County:        req.Country,
			City:          req.City,
			State:         req.Province,
		},
		CreditCard: &payment.CreditCardInfo{
			CraditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.Expiration_Year,
			CreditCardExpirationMount: req.ExpirationMonth,
		},
	}
	fmt.Println()
	fmt.Println(checkoutRpcReq.CreditCard)
	_, err = rpc.CheckoutClient.Checkout(h.Context, checkoutRpcReq)

	if err != nil {
		return nil, err
	}
	return utils.H{
		"Title":    "Waiting",
		"redirect": "/Checkout/Result",
	}, nil
}
