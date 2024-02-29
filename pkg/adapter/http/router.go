package http

import (
	"github.com/uptrace/bunrouter"
)

func SetupRouter(paymentHandler PaymentHandler) *bunrouter.Router {
	router := bunrouter.New()

	router.GET("/health", healthCheckHandler)

	router.GET("/payments/:payment_id", paymentHandler.GetPaymentDetails)
	router.POST("/payments", paymentHandler.ProcessPayment)
	router.PATCH("/payments/:payment_id/refund", paymentHandler.ProcessRefund)

	return router
}
