package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/out"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/utils"
	"time"
)

type PreliminaryExpenseRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.PreliminaryExpenseRepository = (*PreliminaryExpenseRepository)(nil)

func NewPreliminaryExpenseRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *PreliminaryExpenseRepository {
	return &PreliminaryExpenseRepository{db: db, snowflake: snowflake}
}

func (r *PreliminaryExpenseRepository) Get(id uint64) (*model.PreliminaryExpense, error) {
	var preliminaryExpense model.PreliminaryExpense
	err := r.db.Get(&preliminaryExpense, "SELECT * FROM preliminary_expenses WHERE id=$1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	return &preliminaryExpense, nil
}

func (r *PreliminaryExpenseRepository) GetByInquiryID(inquiryID uint64) (*model.PreliminaryExpense, error) {
	var preliminaryExpense model.PreliminaryExpense
	err := r.db.Get(&preliminaryExpense, "SELECT * FROM preliminary_expenses WHERE inquiry_id=$1 LIMIT 1", inquiryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &preliminaryExpense, nil
}

func (r *PreliminaryExpenseRepository) GetByCustomerID(customerID uint64) ([]*model.PreliminaryExpense, error) {
	var preliminaryExpenses []*model.PreliminaryExpense
	err := r.db.Select(&preliminaryExpenses, "SELECT * FROM preliminary_expenses WHERE customer_id=$1", customerID)
	if err != nil {
		return nil, err
	}

	return preliminaryExpenses, nil
}

func (r *PreliminaryExpenseRepository) Create(inquiryID uint64, customerID uint64, rentalID uint64, amount int32, expiresAt time.Time) (*model.PreliminaryExpense, error) {
	existing, err := r.GetByInquiryID(inquiryID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return existing, nil
	}

	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO preliminary_expenses (id, inquiry_id, customer_id, rental_id, amount, expires_at) VALUES (:id, :inquiry_id, :customer_id, :rental_id, :amount, :expires_at)", map[string]any{
		"id":          id,
		"inquiry_id":  inquiryID,
		"customer_id": customerID,
		"rental_id":   rentalID,
		"amount":      amount,
		"expires_at":  expiresAt,
	})
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *PreliminaryExpenseRepository) DeleteWithTx(tx *sqlx.Tx, id uint64) error {
	_, err := tx.Exec("DELETE FROM preliminary_expenses WHERE id = $1", id)

	return err
}

func (r *PreliminaryExpenseRepository) DeleteExpired() error {
	_, err := r.db.Exec("DELETE FROM preliminary_expenses WHERE expires_at < NOW()")

	return err
}
