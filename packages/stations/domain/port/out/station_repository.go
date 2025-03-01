package out

import "github.com/llamadeus/ebike3/packages/stations/domain/model"

// StationRepository is an interface for a repository that handles station operations.
type StationRepository interface {
	// GetStations returns all the stations.
	GetStations() ([]*model.Station, error)

	// GetStationByID returns the station with the given id.
	GetStationByID(id uint64) (*model.Station, error)

	// CreateStation creates a new station with the given name and position.
	CreateStation(name string, positionX float64, positionY float64) (*model.Station, error)

	// UpdateStation updates the station with the given id.
	UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.Station, error)

	// DeleteStation deletes the station with the given id.
	DeleteStation(id uint64) error
}
