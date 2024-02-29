package main

import (
	"log"
	"net/http"

	httpAdapter "paymentplatform/pkg/adapter/http"
	"paymentplatform/pkg/config"
	"paymentplatform/pkg/infra"
	"paymentplatform/pkg/repository"
	"paymentplatform/pkg/usecase"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db := infra.NewDB(cfg)
	defer db.Close()

	bankClient := infra.NewBankClient()

	paymentRepo := repository.NewPaymentRepository(db)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, bankClient)
	paymentHandler := httpAdapter.NewPaymentHandler(paymentUsecase)

	router := httpAdapter.SetupRouter(paymentHandler)

	log.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
