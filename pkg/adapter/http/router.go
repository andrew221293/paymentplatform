package http

import (
	"encoding/base64"
	"net/http"
	"os"
	"strings"

	"github.com/uptrace/bunrouter"
)

func SetupRouter(paymentHandler PaymentHandler) *bunrouter.Router {
	router := bunrouter.New()

	router.GET("/health", HealthCheckHandler)

	router.GET("/payments/:payment_id", basicAuthMiddleware(paymentHandler.GetPaymentDetails))
	router.POST("/payments", basicAuthMiddleware(paymentHandler.ProcessPayment))
	router.PATCH("/payments/:payment_id/refund", basicAuthMiddleware(paymentHandler.ProcessRefund))

	return router
}

func basicAuthMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		// Obtener el nombre de usuario y la contraseña desde las variables de entorno
		username := os.Getenv("BASIC_AUTH_USERNAME")
		password := os.Getenv("BASIC_AUTH_PASSWORD")

		auth := req.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		creds, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		if string(creds) != username+":"+password {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		return next(w, req)
	}
}
