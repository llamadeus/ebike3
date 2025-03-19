package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/out"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/utils"
)

type ExpenseRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.ExpenseRepository = (*ExpenseRepository)(nil)

func NewExpenseRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *ExpenseRepository {
	return &ExpenseRepository{db: db, snowflake: snowflake}
}

func (r *ExpenseRepository) Get(id uint64) (*model.Expense, error) {
	var expense model.Expense
	err := r.db.Get(&expense, "SELECT * FROM expenses WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &expense, nil
}

func (r *ExpenseRepository) GetAll() ([]*model.Expense, error) {
	var expenses []*model.Expense
	err := r.db.Select(&expenses, "SELECT * FROM expenses")
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepository) GetByCustomerID(customerID uint64) ([]*model.Expense, error) {
	var expenses []*model.Expense
	err := r.db.Select(&expenses, "SELECT * FROM expenses WHERE customer_id=$1", customerID)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpenseRepository) Create(customerID uint64, rentalID uint64, amount int) (*model.Expense, error) {
	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO expenses (id, customer_id, rental_id, amount) VALUES (:id, :customer_id, :rental_id, :amount)", map[string]any{
		"id":          id,
		"customer_id": customerID,
		"rental_id":   rentalID,
		"amount":      amount,
	})
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}
