package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Payment struct {
	bun.BaseModel `bun:"payments"`

	ID         int       `bun:"payment_id"`
	MerchantID int       `bun:"merchant_id"`
	CustomerID int       `bun:"customer_id"`
	Amount     float64   `bun:"amount"`
	Currency   string    `bun:"currency"`
	Status     string    `bun:"status"`
	CreatedAt  time.Time `bun:"created_at"`
}

type Refund struct {
	bun.BaseModel `bun:"refunds"`

	ID        int       `bun:"refund_id"`
	PaymentID int       `bun:"payment_id"`
	Amount    float64   `bun:"amount"`
	Status    string    `bun:"status"`
	CreatedAt time.Time `bun:"created_at"`
}

type AuditTrail struct {
	bun.BaseModel `bun:"audit_trail"`

	ID             int       `bun:"event_id"`
	EventType      string    `bun:"event_type"`
	EventTimestamp time.Time `bun:"event_timestamp"`
	UserID         int       `bun:"user_id"`
	Description    string    `bun:"description"`
}
