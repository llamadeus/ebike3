package model

import "time"

type Payment struct {
	ID         uint64        `db:"id"`
	CustomerID uint64        `db:"customer_id"`
	Amount     int           `db:"amount"`
	Status     PaymentStatus `db:"status"`
	CreatedAt  time.Time     `db:"created_at"`
	UpdatedAt  time.Time     `db:"updated_at"`
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusConfirmed PaymentStatus = "CONFIRMED"
	PaymentStatusRejected  PaymentStatus = "REJECTED"
)
