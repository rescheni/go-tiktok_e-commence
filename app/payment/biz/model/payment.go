package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        int64     `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        uint32    `json:"amount"`
	PayAt         time.Time `json:"pay_at"`
}

func (p PaymentLog) TableName() string {
	return "payment_log"
}

func CreatePaymentLog(db *gorm.DB) {

}
