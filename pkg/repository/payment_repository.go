package repository

import (
	"context"
	"fmt"

	"paymentplatform/pkg/entity"

	"github.com/uptrace/bun"
)

//go:generate mockgen -source=payment_repository.go -destination=mocks/payment_repository_mock.go -package=mocks

type PaymentRepository interface {
	SavePayment(ctx context.Context, payment *entity.Payment) error
	GetPaymentByID(ctx context.Context, paymentID int) (*entity.Payment, error)
	UpdatePaymentStatus(ctx context.Context, paymentID int, status string) error
	SaveRefund(ctx context.Context, refund *entity.Refund) error
}

type paymentRepository struct {
	db *bun.DB
}

func NewPaymentRepository(db *bun.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (repo *paymentRepository) SavePayment(ctx context.Context, payment *entity.Payment) error {
	_, err := repo.db.NewInsert().Model(payment).Exec(ctx)
	if err != nil {
		return fmt.Errorf("error saving payment: %w", err)
	}
	return nil
}

func (repo *paymentRepository) GetPaymentByID(ctx context.Context, paymentID int) (*entity.Payment, error) {
	payment := new(entity.Payment)
	err := repo.db.NewSelect().Model(payment).Where("payment_id = ?", paymentID).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving payment by ID %d: %w", paymentID, err)
	}
	return payment, nil
}

func (repo *paymentRepository) UpdatePaymentStatus(ctx context.Context, paymentID int, status string) error {
	_, err := repo.db.NewUpdate().Model(&entity.Payment{}).
		Set("status = ?", status).
		Where("payment_id = ?", paymentID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("error updating payment status for ID %d: %w", paymentID, err)
	}
	return nil
}

func (repo *paymentRepository) SaveRefund(ctx context.Context, refund *entity.Refund) error {
	_, err := repo.db.NewInsert().Model(refund).Exec(ctx)
	if err != nil {
		return fmt.Errorf("error saving refund: %w", err)
	}
	return nil
}
