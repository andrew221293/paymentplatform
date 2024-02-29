package http

import (
	"fmt"
	"net/http"

	"github.com/uptrace/bunrouter"
)

func healthCheckHandler(w http.ResponseWriter, _ bunrouter.Request) error {
	fmt.Println("Hello, World!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
	return nil
}
