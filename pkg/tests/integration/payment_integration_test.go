// tests/integration/payment_usecase_test.go
package integration_test

import (
	"context"
	"log"
	"os"
	"testing"

	"paymentplatform/pkg/config"
	"paymentplatform/pkg/infra"
	"paymentplatform/pkg/repository"
	"paymentplatform/pkg/usecase"

	"github.com/stretchr/testify/require"
)

func TestPaymentUsecase_GetPaymentDetails_Integration(t *testing.T) {
	os.Setenv("POSTGRES_USER", "developer")
	os.Setenv("POSTGRES_PASSWORD", "test123")
	os.Setenv("POSTGRES_DB", "payment_platform")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db := infra.NewDB(cfg)
	defer db.Close()

	bankClient := infra.NewBankClient()

	paymentRepo := repository.NewPaymentRepository(db)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, bankClient)

	paymentID := 1
	payment, err := paymentUsecase.GetPaymentDetails(context.Background(), paymentID)

	require.NoError(t, err)
	require.NotNil(t, payment, "The payment must exist.")
}
