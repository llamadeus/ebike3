package model

import "time"

type Expense struct {
	ID         uint64    `db:"id"`
	CustomerID uint64    `db:"customer_id"`
	RentalID   uint64    `db:"rental_id"`
	Amount     int       `db:"amount"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
