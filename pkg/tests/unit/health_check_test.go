package unit

import (
	"net/http"
	"net/http/httptest"
	"testing"

	http2 "paymentplatform/pkg/adapter/http"

	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bunrouter"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router := bunrouter.New()
	router.GET("/health", http2.HealthCheckHandler)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Should return status code 200")

	assert.Equal(t, "OK", w.Body.String(), "Response body should be 'OK'")
}
