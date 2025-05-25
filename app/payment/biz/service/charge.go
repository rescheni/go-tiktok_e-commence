package service

import (
	"context"
	"e-commence/app/payment/biz/dal/mysql"
	"e-commence/app/payment/biz/model"
	payment "e-commence/rpc_gen/kitex_gen/payment"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {

	card := creditcard.Card{
		Number: req.CreditCard.CraditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMount)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	err = card.Validate(true)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(40099, err.Error())
	}

	resp = &payment.ChargeResp{}
	transactionId, err := uuid.NewRandom()
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(40049, "Get UUID Error")
	}

	err = model.CreatePaymentLog(mysql.DB, s.ctx, &model.PaymentLog{
		UserId:        int64(req.UserId),
		OrderId:       req.OrderId,
		TransactionId: transactionId.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})

	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(40050, "Create Payment Log Error: "+err.Error())
	}

	return &payment.ChargeResp{
		TransactionId: transactionId.String(),
	}, nil
}
