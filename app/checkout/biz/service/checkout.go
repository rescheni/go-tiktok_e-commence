package service

import (
	"context"
	"e-commence/app/checkout/infra/rpc"
	"e-commence/rpc_gen/kitex_gen/cart"
	checkout "e-commence/rpc_gen/kitex_gen/checkout"
	"e-commence/rpc_gen/kitex_gen/payment"
	"e-commence/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info

func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {

	CartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: uint64(req.UserId),
	})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50001, err.Error())
	}

	if CartResult == nil || CartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(50002, err.Error())
	}
	Total := float32(0)

	for _, val := range CartResult.Items {
		productinfo, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: val.ProductId,
		})
		if err != nil {
			return nil, kerrors.NewGRPCBizStatusError(50002, err.Error())
		}

		if productinfo.Product == nil {
			continue
		}

		Total += float32(val.Quantity) * productinfo.Product.Price
	}

	var order_id string
	u, _ := uuid.NewRandom()
	order_id = u.String()

	payReq := &payment.ChargeReq{
		UserId:     uint64(req.UserId),
		OrderId:    order_id,
		Amount:     uint32(Total),
		CreditCard: req.CreditCard,
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartreq{
		UserId: uint64(req.UserId),
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50050, err.Error())
	}
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}
	klog.Info(paymentResult)

	return &checkout.CheckoutResp{
		OrderId:       order_id,
		TransactionId: paymentResult.TransactionId,
	}, nil
}
