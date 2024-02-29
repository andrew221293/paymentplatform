package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"paymentplatform/pkg/entity"
	"paymentplatform/pkg/usecase"

	"github.com/uptrace/bunrouter"
)

type PaymentHandler interface {
	ProcessPayment(w http.ResponseWriter, req bunrouter.Request) error
	GetPaymentDetails(w http.ResponseWriter, req bunrouter.Request) error
	ProcessRefund(w http.ResponseWriter, req bunrouter.Request) error
}

type paymentHandler struct {
	paymentUsecase usecase.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase usecase.PaymentUsecase) PaymentHandler {
	return &paymentHandler{
		paymentUsecase: paymentUsecase,
	}
}

func (ph *paymentHandler) ProcessPayment(w http.ResponseWriter, req bunrouter.Request) error {
	var paymentDetails entity.PaymentDetails
	if err := json.NewDecoder(req.Body).Decode(&paymentDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	merchantIDStr := req.Header.Get("Merchant-ID")
	customerIDStr := req.Header.Get("Customer-ID")

	merchantID, err := strconv.Atoi(merchantIDStr)
	if err != nil {
		http.Error(w, "Invalid Merchant ID", http.StatusBadRequest)
		return err
	}

	customerID, err := strconv.Atoi(customerIDStr)
	if err != nil {
		http.Error(w, "Invalid Customer ID", http.StatusBadRequest)
		return err
	}

	err = ph.paymentUsecase.ProcessPayment(req.Context(), paymentDetails, merchantID, customerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode("Payment processed successfully")
}

func (ph *paymentHandler) GetPaymentDetails(w http.ResponseWriter, req bunrouter.Request) error {
	paymentID, err := strconv.Atoi(req.Param("payment_id"))
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return err
	}

	payment, err := ph.paymentUsecase.GetPaymentDetails(req.Context(), paymentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(payment)
}

func (ph *paymentHandler) ProcessRefund(w http.ResponseWriter, req bunrouter.Request) error {
	paymentID, err := strconv.Atoi(req.Param("payment_id"))
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return err
	}

	var refundDetails entity.RefundDetails
	if err := json.NewDecoder(req.Body).Decode(&refundDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	refundDetails.PaymentID = paymentID

	err = ph.paymentUsecase.ProcessRefund(req.Context(), refundDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode("Refund processed successfully")
}
