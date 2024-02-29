package entity

type (
	PaymentDetails struct {
		Amount     float64
		Currency   string
		CardNumber string
		CVV        string
		ExpiryDate string
	}

	PaymentResponse struct {
		Success bool
		Message string
	}

	RefundDetails struct {
		PaymentID int
		Amount    float64
	}

	RefundResponse struct {
		Success bool
		Message string
	}
)
