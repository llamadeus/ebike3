package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
	"github.com/llamadeus/ebike3/packages/rentals/infrastructure/utils"
	"time"
)

type RentalRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.RentalRepository = (*RentalRepository)(nil)

func NewRentalRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *RentalRepository {
	return &RentalRepository{db: db, snowflake: snowflake}
}

func (r *RentalRepository) Get(id uint64) (*model.Rental, error) {
	var rental model.Rental
	err := r.db.Get(&rental, "SELECT * FROM rentals WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rental, nil
}

func (r *RentalRepository) GetActiveRentalByCustomerID(customerID uint64) (*model.Rental, error) {
	var rental model.Rental
	err := r.db.Get(&rental, "SELECT * FROM rentals WHERE customer_id=$1 AND `end` IS NULL LIMIT 1", customerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rental, nil
}

func (r *RentalRepository) GetPastRentalsByCustomerID(customerID uint64) ([]*model.Rental, error) {
	var rentals []*model.Rental
	err := r.db.Select(&rentals, "SELECT * FROM rentals WHERE customer_id=$1 AND `end` IS NOT NULL", customerID)
	if err != nil {
		return nil, err
	}

	return rentals, nil
}

func (r *RentalRepository) GetActiveRentalByVehicleID(vehicleID uint64) (*model.Rental, error) {
	var rental model.Rental
	err := r.db.Get(&rental, "SELECT * FROM rentals WHERE vehicle_id=$1 AND `end` IS NULL LIMIT 1", vehicleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rental, nil
}

func (r *RentalRepository) CreateRental(customerID uint64, vehicleID uint64) (*model.Rental, error) {
	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO rentals (id, customer_id, vehicle_id, start, `end`) VALUES (:id, :customer_id, :vehicle_id, :start, :end)", map[string]any{
		"id":          id,
		"customer_id": customerID,
		"vehicle_id":  vehicleID,
		"start":       time.Now(),
		"end":         nil,
	})
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

func (r *RentalRepository) StopRental(id uint64) (*model.Rental, error) {
	_, err := r.db.Exec("UPDATE rentals SET `end` = $1 WHERE id = $2", time.Now(), id)
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}
