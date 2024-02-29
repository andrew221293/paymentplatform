package unit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	http2 "paymentplatform/pkg/adapter/http"
	"paymentplatform/pkg/entity"
	mock_usecase "paymentplatform/pkg/usecase/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
)

func TestPaymentHandler_ProcessPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUsecase := mock_usecase.NewMockPaymentUsecase(ctrl)
	handler := http2.NewPaymentHandler(mockPaymentUsecase)

	tests := []struct {
		name           string
		setupMocks     func()
		requestBody    string
		wantStatusCode int
		wantResponse   string
	}{
		{
			name: "Success: Payment processed successfully",
			setupMocks: func() {
				mockPaymentUsecase.EXPECT().
					ProcessPayment(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil).Times(1)
			},
			requestBody:    `{"amount": 100, "currency": "USD", "cardNumber": "4111111111111111", "cvv": "123", "expiryDate": "10/24"}`,
			wantStatusCode: http.StatusOK,
			wantResponse:   "Payment processed successfully",
		},
		{
			name:           "Failure: Invalid request body",
			setupMocks:     func() {},
			requestBody:    `{"amount": "invalid", "currency": "USD"}`,
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMocks != nil {
				tt.setupMocks()
			}

			req := httptest.NewRequest(http.MethodPost, "/payments", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Merchant-ID", "123")
			req.Header.Set("Customer-ID", "456")
			w := httptest.NewRecorder()

			router := bunrouter.New()
			router.POST("/payments", handler.ProcessPayment)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			if tt.wantResponse != "" {
				assert.Contains(t, w.Body.String(), tt.wantResponse)
			}
		})
	}
}

func TestPaymentHandler_GetPaymentDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUsecase := mock_usecase.NewMockPaymentUsecase(ctrl)
	handler := http2.NewPaymentHandler(mockPaymentUsecase)

	tests := []struct {
		name           string
		setupMocks     func()
		paymentID      string
		wantStatusCode int
		wantResponse   *entity.Payment
		wantErrMessage string
	}{
		{
			name: "Success: Payment details retrieved",
			setupMocks: func() {
				mockPaymentUsecase.EXPECT().
					GetPaymentDetails(gomock.Any(), gomock.Eq(1)).
					Return(&entity.Payment{
						ID:         1,
						MerchantID: 123,
						CustomerID: 456,
						Amount:     100.00,
						Currency:   "USD",
						Status:     "Processed",
						CreatedAt:  time.Now(),
					}, nil).Times(1)
			},
			paymentID:      "1",
			wantStatusCode: http.StatusOK,
			wantResponse: &entity.Payment{
				ID:         1,
				MerchantID: 123,
				CustomerID: 456,
				Amount:     100.00,
				Currency:   "USD",
				Status:     "Processed",
			},
		},
		{
			name:           "Failure: Invalid payment ID format",
			setupMocks:     func() {},
			paymentID:      "abc",
			wantStatusCode: http.StatusBadRequest,
			wantErrMessage: "Invalid payment ID",
		},
		{
			name: "Failure: Payment details not found",
			setupMocks: func() {
				mockPaymentUsecase.EXPECT().
					GetPaymentDetails(gomock.Any(), gomock.Eq(2)).
					Return(nil, fmt.Errorf("payment not found")).Times(1)
			},
			paymentID:      "2",
			wantStatusCode: http.StatusInternalServerError,
			wantErrMessage: "payment not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMocks != nil {
				tt.setupMocks()
			}

			req := httptest.NewRequest(http.MethodGet, "/payments/"+tt.paymentID, nil)
			w := httptest.NewRecorder()

			router := bunrouter.New()
			router.GET("/payments/:payment_id", handler.GetPaymentDetails)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			if tt.wantResponse != nil {
				var got entity.Payment
				err := json.NewDecoder(w.Body).Decode(&got)
				assert.NoError(t, err)
				assert.Equal(t, tt.wantResponse.ID, got.ID)
				assert.Equal(t, tt.wantResponse.Amount, got.Amount)
			}
			if tt.wantErrMessage != "" {
				assert.Contains(t, w.Body.String(), tt.wantErrMessage)
			}
		})
	}
}

func TestPaymentHandler_ProcessRefund(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentUsecase := mock_usecase.NewMockPaymentUsecase(ctrl)
	handler := http2.NewPaymentHandler(mockPaymentUsecase)

	tests := []struct {
		name           string
		setupMocks     func()
		paymentID      string
		requestBody    string
		wantStatusCode int
		wantResponse   string
	}{
		{
			name: "Success: Refund processed successfully",
			setupMocks: func() {
				mockPaymentUsecase.EXPECT().
					ProcessRefund(gomock.Any(), gomock.Any()).
					Return(nil).Times(1)
			},
			paymentID:      "1",
			requestBody:    `{"amount": 50.00}`,
			wantStatusCode: http.StatusOK,
			wantResponse:   "Refund processed successfully",
		},
		{
			name:           "Failure: Invalid payment ID format",
			setupMocks:     func() {},
			paymentID:      "invalid",
			requestBody:    `{"amount": 50.00}`,
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   "Invalid payment ID",
		},
		{
			name:           "Failure: Invalid request body",
			setupMocks:     func() {},
			paymentID:      "1",
			requestBody:    `{"amount": "invalid"}`,
			wantStatusCode: http.StatusBadRequest,
			wantResponse:   "",
		},
		{
			name: "Failure: Refund processing fails",
			setupMocks: func() {
				mockPaymentUsecase.EXPECT().
					ProcessRefund(gomock.Any(), gomock.Any()).
					Return(fmt.Errorf("refund processing error")).Times(1)
			},
			paymentID:      "2",
			requestBody:    `{"amount": 25.00}`,
			wantStatusCode: http.StatusInternalServerError,
			wantResponse:   "refund processing error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMocks != nil {
				tt.setupMocks()
			}

			req := httptest.NewRequest(http.MethodPost, "/refunds/"+tt.paymentID, bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router := bunrouter.New()
			router.POST("/refunds/:payment_id", handler.ProcessRefund)

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatusCode, w.Code)
			if tt.wantResponse != "" {
				assert.Contains(t, w.Body.String(), tt.wantResponse)
			}
		})
	}
}
