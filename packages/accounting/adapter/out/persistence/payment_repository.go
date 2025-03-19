package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/accounting/domain/model"
	"github.com/llamadeus/ebike3/packages/accounting/domain/port/out"
	"github.com/llamadeus/ebike3/packages/accounting/infrastructure/utils"
)

type PaymentRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.PaymentRepository = (*PaymentRepository)(nil)

func NewPaymentRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *PaymentRepository {
	return &PaymentRepository{db: db, snowflake: snowflake}
}

func (r *PaymentRepository) Get(id uint64) (*model.Payment, error) {
	var payment model.Payment
	err := r.db.Get(&payment, "SELECT * FROM payments WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepository) GetAll() ([]*model.Payment, error) {
	var payments []*model.Payment
	err := r.db.Select(&payments, "SELECT * FROM payments")
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) GetByCustomerID(customerID uint64) ([]*model.Payment, error) {
	var payments []*model.Payment
	err := r.db.Select(&payments, "SELECT * FROM payments WHERE customer_id=$1", customerID)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *PaymentRepository) Create(customerID uint64, amount int32) (*model.Payment, error) {
	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO payments (id, customer_id, amount, status) VALUES (:id, :customer_id, :amount, :status)", map[string]any{
		"id":          id,
		"customer_id": customerID,
		"amount":      amount,
		"status":      model.PaymentStatusPending,
	})
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *PaymentRepository) Update(id uint64, status model.PaymentStatus) (*model.Payment, error) {
	_, err := r.db.Exec("UPDATE payments SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *PaymentRepository) Delete(id uint64) error {
	_, err := r.db.Exec("DELETE FROM payments WHERE id = $1", id)

	return err
}
