package dto

import (
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"time"
)

type PreliminaryExpenseDTO struct {
	ID         string `json:"id"`
	InquiryID  string `json:"inquiryId"`
	CustomerID string `json:"customerId"`
	Amount     int32  `json:"amount"`
	CreatedAt  string `json:"createdAt"`
	ExpiresAt  string `json:"expiresAt"`
}

func PreliminaryExpenseToDTO(expense *model.PreliminaryExpense) *PreliminaryExpenseDTO {
	return &PreliminaryExpenseDTO{
		ID:         IDToDTO(expense.ID),
		InquiryID:  IDToDTO(expense.InquiryID),
		CustomerID: IDToDTO(expense.CustomerID),
		Amount:     expense.Amount,
		CreatedAt:  expense.CreatedAt.Format(time.RFC3339),
		ExpiresAt:  expense.ExpiresAt.Format(time.RFC3339),
	}
}
