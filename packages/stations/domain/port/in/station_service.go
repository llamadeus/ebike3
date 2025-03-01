package in

import "github.com/llamadeus/ebike3/packages/stations/domain/model"

type StationService interface {
	// CreateStation creates a new station with the given name and position.
	CreateStation(name string, positionX float64, positionY float64) (*model.Station, error)

	// UpdateStation updates the station with the given id.
	UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.Station, error)

	// DeleteStation deletes the station with the given id.
	DeleteStation(id uint64) (*model.Station, error)

	// GetStationViews returns all the stations.
	GetStationViews() ([]*model.StationView, error)

	// CreateStationView creates a new station with the given name and position.
	CreateStationView(id uint64, name string, positionX float64, positionY float64) error

	// UpdateStationView updates the station with the given id.
	UpdateStationView(id uint64, name string, positionX float64, positionY float64) error

	// DeleteStationView deletes the station with the given id.
	DeleteStationView(id uint64) error
}
