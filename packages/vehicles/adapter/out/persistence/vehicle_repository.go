package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/out"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/utils"
)

type VehicleRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.VehicleRepository = (*VehicleRepository)(nil)

func NewVehicleRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *VehicleRepository {
	return &VehicleRepository{db: db, snowflake: snowflake}
}

func (r *VehicleRepository) GetVehicles() ([]*model.Vehicle, error) {
	var vehicles []*model.Vehicle
	err := r.db.Select(&vehicles, "SELECT * FROM vehicles")
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehicleRepository) GetVehicleByID(id uint64) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	err := r.db.Get(&vehicle, "SELECT * FROM vehicles WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &vehicle, nil
}

func (r *VehicleRepository) CreateVehicle(type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.Vehicle, error) {
	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO vehicles (id, type, position_x, position_y, battery) VALUES (:id, :type, :position_x, :position_y, :battery)", map[string]any{
		"id":         id,
		"type":       type_,
		"position_x": positionX,
		"position_y": positionY,
		"battery":    battery,
	})
	if err != nil {
		return nil, err
	}

	return r.GetVehicleByID(id)
}

func (r *VehicleRepository) UpdateVehicle(id uint64, positionX float64, positionY float64, battery float64) (*model.Vehicle, error) {
	_, err := r.db.Exec("UPDATE vehicles SET position_x = $1, position_y = $2, battery = $3 WHERE id = $4", positionX, positionY, battery, id)
	if err != nil {
		return nil, err
	}

	return r.GetVehicleByID(id)
}

func (r *VehicleRepository) DeleteVehicle(id uint64) error {
	_, err := r.db.Exec("DELETE FROM vehicles WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
