package events

const (
	AccountingTopic                   = "accounting"
	AccountingPaymentCreatedEventType = "PaymentCreated"
	AccountingPaymentUpdatedEventType = "PaymentUpdated"
	AccountingPaymentDeletedEventType = "PaymentDeleted"
	AccountingExpenseCreatedEventType = "ExpenseCreated"
)

type PaymentCreatedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
}

type PaymentUpdatedEvent struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type PaymentDeletedEvent struct {
	ID string `json:"id"`
}

type ExpenseCreatedEvent struct {
	ID         string `json:"id"`
	CustomerID string `json:"customerId"`
	RentalID   string `json:"rentalId"`
	Amount     int    `json:"amount"`
}
