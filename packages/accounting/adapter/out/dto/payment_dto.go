package dto

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"time"
)

type PaymentDTO struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Amount     int32  `json:"amount"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func PaymentToDTO(payment *model.Payment) *PaymentDTO {
	return &PaymentDTO{
		ID:         IDToDTO(payment.ID),
		CustomerID: IDToDTO(payment.CustomerID),
		Amount:     payment.Amount,
		Status:     StatusToDTO(payment.Status),
		CreatedAt:  payment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  payment.UpdatedAt.Format(time.RFC3339),
	}
}

func StatusToDTO(status model.PaymentStatus) string {
	return string(status)
}

func StatusFromDTO(status string) (model.PaymentStatus, error) {
	switch status {
	case string(model.PaymentStatusPending):
		return model.PaymentStatusPending, nil
	case string(model.PaymentStatusConfirmed):
		return model.PaymentStatusConfirmed, nil
	case string(model.PaymentStatusRejected):
		return model.PaymentStatusRejected, nil
	}

	return model.PaymentStatusPending, fmt.Errorf("invalid payment status: %s", status)
}
