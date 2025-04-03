package out

import (
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"time"
)

type PreliminaryExpenseRepository interface {
	Get(id uint64) (*model.PreliminaryExpense, error)

	GetByCustomerID(customerID uint64) ([]*model.PreliminaryExpense, error)

	Create(inquiryID uint64, customerID uint64, rentalID uint64, amount int32, expiresAt time.Time) (*model.PreliminaryExpense, error)

	DeleteWithTx(tx *sqlx.Tx, id uint64) error

	DeleteExpired() error
}
