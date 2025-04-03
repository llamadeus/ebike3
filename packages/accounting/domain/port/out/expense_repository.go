package out

import (
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
)

type ExpenseRepository interface {
	GetAll() ([]*model.Expense, error)

	GetByCustomerID(customerID uint64) ([]*model.Expense, error)

	Create(customerID uint64, rentalID uint64, amount int32) (*model.Expense, error)

	CreateWithTx(tx *sqlx.Tx, customerID uint64, rentalID uint64, amount int32) (*model.Expense, error)
}
