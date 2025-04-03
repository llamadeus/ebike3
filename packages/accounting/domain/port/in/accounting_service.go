package in

import "github.com/llamadeus/ebike3/packages/accounting/domain/model"

type AccountingService interface {
	GetAllPayments() ([]*model.Payment, error)

	GetPaymentsForCustomer(customerID uint64) ([]*model.Payment, error)

	CreatePayment(customerID uint64, amount int32) (*model.Payment, error)

	UpdatePayment(id uint64, status model.PaymentStatus) (*model.Payment, error)

	DeleteCustomerPayment(id uint64, customerID uint64) (*model.Payment, error)

	GetExpensesForCustomer(customerID uint64) ([]*model.Expense, error)

	CreateExpense(customerID uint64, rentalID uint64, amount int32) (*model.Expense, error)

	CreatePreliminaryExpense(inquiryID uint64, customerID uint64, rentalID uint64, amount int32) (*model.PreliminaryExpense, error)

	FinalizePreliminaryExpense(id uint64) (*model.Expense, error)

	GetCreditBalanceForCustomer(customerID uint64) (int32, error)
}
