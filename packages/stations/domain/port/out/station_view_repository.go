package out

import "github.com/llamadeus/ebike3/packages/stations/domain/model"

// StationViewRepository is an interface for a Mongo repository for station view data.
type StationViewRepository interface {
	// GetStations returns all the stations.
	GetStations() ([]*model.StationView, error)

	// GetStationByID returns the station with the given id.
	GetStationByID(id uint64) (*model.StationView, error)

	// CreateStation creates a new station with the given name and position.
	CreateStation(id uint64, name string, positionX float64, positionY float64) (*model.StationView, error)

	// UpdateStation updates the station with the given id.
	UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.StationView, error)

	// DeleteStation deletes the station with the given id.
	DeleteStation(id uint64) error
}
