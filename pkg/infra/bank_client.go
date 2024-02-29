package infra

import (
	"context"
	"paymentplatform/pkg/entity"
)

//go:generate mockgen -source=bank_client.go -destination=mocks/back_client_mock.go -package=mocks

type BankClient interface {
	ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails) (entity.PaymentResponse, error)
	ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) (entity.RefundResponse, error)
}

type bankClient struct {
}

func NewBankClient() BankClient {
	return &bankClient{}
}

func (bc *bankClient) ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails) (entity.PaymentResponse, error) {
	return entity.PaymentResponse{
		Success: true,
		Message: "Payment processed successfully",
	}, nil
}

func (bc *bankClient) ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) (entity.RefundResponse, error) {
	return entity.RefundResponse{
		Success: true,
		Message: "Refund processed successfully",
	}, nil
}
