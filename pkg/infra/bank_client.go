package infra

import (
	"context"
	"paymentplatform/pkg/entity"
)

type BankClient struct {
}

func NewBankClient() *BankClient {
	return &BankClient{}
}

func (bc *BankClient) ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails) (entity.PaymentResponse, error) {
	return entity.PaymentResponse{
		Success: true,
		Message: "Payment processed successfully",
	}, nil
}

func (bc *BankClient) ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) (entity.RefundResponse, error) {
	return entity.RefundResponse{
		Success: true,
		Message: "Refund processed successfully",
	}, nil
}
