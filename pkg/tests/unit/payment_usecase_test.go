package unit

import (
	"context"
	"fmt"
	"testing"

	"paymentplatform/pkg/entity"
	mock_infra "paymentplatform/pkg/infra/mocks"
	mock_repository "paymentplatform/pkg/repository/mocks"
	"paymentplatform/pkg/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestProcessPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankClient := mock_infra.NewMockBankClient(ctrl)
	mockPaymentRepo := mock_repository.NewMockPaymentRepository(ctrl)

	ctx := context.Background()

	tests := []struct {
		name           string
		setupMocks     func()
		paymentDetails entity.PaymentDetails
		wantErr        bool
	}{
		{
			name: "Success: Payment processed successfully",
			setupMocks: func() {
				mockBankClient.EXPECT().
					ProcessPayment(ctx, gomock.Any()).
					Return(entity.PaymentResponse{Success: true, Message: "Payment processed successfully"}, nil)

				mockPaymentRepo.EXPECT().
					SavePayment(ctx, gomock.Any()).
					Return(nil)
			},
			paymentDetails: entity.PaymentDetails{
				Amount:     100.00,
				Currency:   "USD",
				CardNumber: "4111111111111111",
				CVV:        "123",
				ExpiryDate: "10/24",
			},
			wantErr: false,
		},
		{
			name: "Failure: Bank processing fails",
			setupMocks: func() {
				mockBankClient.EXPECT().
					ProcessPayment(ctx, gomock.Any()).
					Return(entity.PaymentResponse{Success: false, Message: "Payment failed"}, fmt.Errorf("bank processing error"))
			},
			paymentDetails: entity.PaymentDetails{
				Amount:     50.00,
				Currency:   "EUR",
				CardNumber: "4222222222222",
				CVV:        "321",
				ExpiryDate: "11/25",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			uc := usecase.NewPaymentUsecase(mockPaymentRepo, mockBankClient)
			err := uc.ProcessPayment(ctx, tt.paymentDetails, 12345, 67890)

			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetPaymentDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepo := mock_repository.NewMockPaymentRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name       string
		setupMocks func()
		paymentID  int
		want       *entity.Payment
		wantErr    bool
	}{
		{
			name: "Success: Payment details retrieved successfully",
			setupMocks: func() {
				mockPaymentRepo.EXPECT().
					GetPaymentByID(ctx, gomock.Eq(1)).
					Return(&entity.Payment{
						ID:         1,
						MerchantID: 12345,
						CustomerID: 67890,
						Amount:     100.00,
						Currency:   "USD",
						Status:     "Processed",
					}, nil)
			},
			paymentID: 1,
			want: &entity.Payment{
				ID:         1,
				MerchantID: 12345,
				CustomerID: 67890,
				Amount:     100.00,
				Currency:   "USD",
				Status:     "Processed",
			},
			wantErr: false,
		},
		{
			name: "Failure: Payment details not found",
			setupMocks: func() {
				mockPaymentRepo.EXPECT().
					GetPaymentByID(ctx, gomock.Eq(2)).
					Return(nil, fmt.Errorf("payment not found"))
			},
			paymentID: 2,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			uc := usecase.NewPaymentUsecase(mockPaymentRepo, nil) // BankClient not used in this test
			got, err := uc.GetPaymentDetails(ctx, tt.paymentID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaymentDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !assert.Equal(t, tt.want, got) {
				t.Errorf("GetPaymentDetails() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessRefund(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBankClient := mock_infra.NewMockBankClient(ctrl)
	mockPaymentRepo := mock_repository.NewMockPaymentRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name          string
		setupMocks    func()
		refundDetails entity.RefundDetails
		wantErr       bool
	}{
		{
			name: "Success: Refund processed successfully",
			setupMocks: func() {
				mockBankClient.EXPECT().
					ProcessRefund(ctx, gomock.Any()).
					Return(entity.RefundResponse{Success: true, Message: "Refund processed successfully"}, nil)

				mockPaymentRepo.EXPECT().
					UpdatePaymentStatus(ctx, gomock.Any(), "Refunded").
					Return(nil)

				mockPaymentRepo.EXPECT().
					SaveRefund(ctx, gomock.Any()).
					Return(nil)
			},
			refundDetails: entity.RefundDetails{
				PaymentID: 1,
				Amount:    50.00,
			},
			wantErr: false,
		},
		{
			name: "Failure: Bank refund processing fails",
			setupMocks: func() {
				mockBankClient.EXPECT().
					ProcessRefund(ctx, gomock.Any()).
					Return(entity.RefundResponse{Success: false, Message: "Refund processing failed"}, fmt.Errorf("bank refund processing error"))
			},
			refundDetails: entity.RefundDetails{
				PaymentID: 2,
				Amount:    25.00,
			},
			wantErr: true,
		},
		{
			name: "Failure: Update payment status fails",
			setupMocks: func() {
				mockBankClient.EXPECT().
					ProcessRefund(ctx, gomock.Any()).
					Return(entity.RefundResponse{Success: true, Message: "Refund processed successfully"}, nil)

				mockPaymentRepo.EXPECT().
					UpdatePaymentStatus(ctx, gomock.Any(), "Refunded").
					Return(fmt.Errorf("failed to update payment status"))
			},
			refundDetails: entity.RefundDetails{
				PaymentID: 3,
				Amount:    75.00,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			uc := usecase.NewPaymentUsecase(mockPaymentRepo, mockBankClient)
			err := uc.ProcessRefund(ctx, tt.refundDetails)

			if (err != nil) != tt.wantErr {
				t.Errorf("ProcessRefund() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
