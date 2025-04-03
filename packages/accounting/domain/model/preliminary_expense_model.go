package model

import "time"

type PreliminaryExpense struct {
	ID         uint64    `db:"id"`
	InquiryID  uint64    `db:"inquiry_id"`
	CustomerID uint64    `db:"customer_id"`
	RentalID   uint64    `db:"rental_id"`
	Amount     int32     `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
	ExpiresAt  time.Time `db:"expires_at"`
}
