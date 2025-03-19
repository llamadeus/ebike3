package dto

import (
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"time"
)

type ExpenseDTO struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	RentalID   string `json:"rentalId"`
	Amount     int    `json:"amount"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func ExpenseToDTO(expense *model.Expense) *ExpenseDTO {
	return &ExpenseDTO{
		ID:         IDToDTO(expense.ID),
		CustomerID: IDToDTO(expense.CustomerID),
		RentalID:   IDToDTO(expense.RentalID),
		Amount:     expense.Amount,
		CreatedAt:  expense.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  expense.UpdatedAt.Format(time.RFC3339),
	}
}
