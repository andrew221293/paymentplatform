package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"paymentplatform/pkg/entity"
	"paymentplatform/pkg/infra"
	"paymentplatform/pkg/repository"
)

//go:generate mockgen -source=payment_usecase.go -destination=mocks/payment_usecase_mock.go -package=mocks

type PaymentUsecase interface {
	ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails, merchantID, customerID int) error
	GetPaymentDetails(ctx context.Context, paymentID int) (*entity.Payment, error)
	ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) error
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepository
	bankClient  infra.BankClient
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepository, bankClient infra.BankClient) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		bankClient:  bankClient,
	}
}

func (uc *paymentUsecase) ProcessPayment(ctx context.Context, paymentDetails entity.PaymentDetails, merchantID, customerID int) error {
	response, err := uc.bankClient.ProcessPayment(ctx, paymentDetails)
	if err != nil {
		return err
	}
	if !response.Success {
		return fmt.Errorf("payment processing failed: %s", response.Message)
	}

	payment := &entity.Payment{
		MerchantID: generateRandomID(),
		CustomerID: generateRandomID(),
		Amount:     paymentDetails.Amount,
		Currency:   paymentDetails.Currency,
		Status:     "Processed",
		CreatedAt:  time.Now(),
	}

	return uc.paymentRepo.SavePayment(ctx, payment)
}

func (uc *paymentUsecase) GetPaymentDetails(ctx context.Context, paymentID int) (*entity.Payment, error) {
	return uc.paymentRepo.GetPaymentByID(ctx, paymentID)
}

func (uc *paymentUsecase) ProcessRefund(ctx context.Context, refundDetails entity.RefundDetails) error {
	response, err := uc.bankClient.ProcessRefund(ctx, refundDetails)
	if err != nil {
		return err
	}
	if !response.Success {
		return fmt.Errorf("refund processing failed: %s", response.Message)
	}

	err = uc.paymentRepo.UpdatePaymentStatus(ctx, refundDetails.PaymentID, "Refunded")
	if err != nil {
		return fmt.Errorf("failed to update payment status to refunded: %v", err)
	}

	refund := &entity.Refund{
		PaymentID: refundDetails.PaymentID,
		Amount:    refundDetails.Amount,
		Status:    "Processed",
		CreatedAt: time.Now(),
	}
	return uc.paymentRepo.SaveRefund(ctx, refund)
}

func generateRandomID() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(90000) + 10000
}
