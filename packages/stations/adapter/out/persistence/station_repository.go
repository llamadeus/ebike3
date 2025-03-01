package persistence

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/llamadeus/ebike3/packages/stations/domain/model"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/out"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/utils"
)

type StationRepository struct {
	db        *sqlx.DB
	snowflake *utils.SnowflakeGenerator
}

var _ out.StationRepository = (*StationRepository)(nil)

func NewStationRepository(db *sqlx.DB, snowflake *utils.SnowflakeGenerator) *StationRepository {
	return &StationRepository{db: db, snowflake: snowflake}
}

func (r *StationRepository) GetStations() ([]*model.Station, error) {
	var stations []*model.Station
	err := r.db.Select(&stations, "SELECT * FROM stations")
	if err != nil {
		return nil, err
	}

	return stations, nil
}

func (r *StationRepository) GetStationByID(id uint64) (*model.Station, error) {
	var station model.Station
	err := r.db.Get(&station, "SELECT * FROM stations WHERE id=$1 LIMIT 1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &station, nil
}

func (r *StationRepository) CreateStation(name string, positionX float64, positionY float64) (*model.Station, error) {
	id, err := r.snowflake.Generate()
	if err != nil {
		return nil, err
	}

	_, err = r.db.NamedExec("INSERT INTO stations (id, name, position_x, position_y) VALUES (:id, :name, :position_x, :position_y)", map[string]any{
		"id":         id,
		"name":       name,
		"position_x": positionX,
		"position_y": positionY,
	})
	if err != nil {
		return nil, err
	}

	return r.GetStationByID(id)
}

func (r *StationRepository) UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.Station, error) {
	_, err := r.db.Exec("UPDATE stations SET name = $1, position_x = $2, position_y = $3 WHERE id = $4", name, positionX, positionY, id)
	if err != nil {
		return nil, err
	}

	return r.GetStationByID(id)
}

func (r *StationRepository) DeleteStation(id uint64) error {
	_, err := r.db.Exec("DELETE FROM stations WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
