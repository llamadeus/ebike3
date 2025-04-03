package events

import "time"

const (
	AccountingTopic                   = "accounting"
	AccountingPaymentCreatedEventType = "PaymentCreated"
	AccountingPaymentUpdatedEventType = "PaymentUpdated"
	AccountingPaymentDeletedEventType = "PaymentDeleted"
	AccountingExpenseCreatedEventType = "ExpenseCreated"

	AccountingPreliminaryExpenseCreatedEventType   = "PreliminaryExpenseCreated"
	AccountingPreliminaryExpenseFinalizedEventType = "PreliminaryExpenseFinalized"
)

type PaymentCreatedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Amount     int32  `json:"amount"`
	Status     string `json:"status"`
}

type PaymentUpdatedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Amount     int32  `json:"amount"`
	Status     string `json:"status"`
}

type PaymentDeletedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Amount     int32  `json:"amount"`
	Status     string `json:"status"`
}

type ExpenseCreatedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	RentalID   string `json:"rentalId"`
	Amount     int32  `json:"amount"`
}

type PreliminaryExpenseCreatedEvent struct {
	ID         string    `json:"id"`
	InquiryID  string    `json:"inquiryId"`
	CustomerID string    `json:"customerId"`
	Amount     int32     `json:"amount"`
	ExpiresAt  time.Time `json:"expiresAt"`
}

type PreliminaryExpenseFinalizedEvent struct {
	ID         string `json:"id"`
	InquiryID  string `json:"inquiryId"`
	CustomerID string `json:"customerId"`
	RentalID   string `json:"rentalId"`
	Amount     int32  `json:"amount"`
}
