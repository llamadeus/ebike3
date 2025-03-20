package service

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/stations/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/stations/domain/events"
	"github.com/llamadeus/ebike3/packages/stations/domain/model"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/in"
	"github.com/llamadeus/ebike3/packages/stations/domain/port/out"
	"github.com/llamadeus/ebike3/packages/stations/infrastructure/micro"
)

// StationService implements the StationService interface.
type StationService struct {
	kafka          micro.Kafka
	repository     out.StationRepository
	viewRepository out.StationViewRepository
}

// Ensure that StationService implements the StationService interface.
var _ in.StationService = (*StationService)(nil)

// NewStationService creates a new instance of the StationService.
func NewStationService(kafka micro.Kafka, repository out.StationRepository, viewRepository out.StationViewRepository) *StationService {
	return &StationService{
		kafka:          kafka,
		repository:     repository,
		viewRepository: viewRepository,
	}
}

// CreateStation creates a new station with the given name and position.
func (s *StationService) CreateStation(name string, positionX float64, positionY float64) (*model.Station, error) {
	station, err := s.repository.CreateStation(name, positionX, positionY)
	if err != nil {
		return nil, err
	}

	event := micro.NewEvent(events.StationsStationCreatedEventType, events.StationCreatedEvent{
		ID:        dto.IDToDTO(station.ID),
		Name:      station.Name,
		PositionX: station.PositionX,
		PositionY: station.PositionY,
	})
	err = s.kafka.Producer().Send(events.StationsTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return station, nil
}

// UpdateStation updates the station with the given id.
func (s *StationService) UpdateStation(id uint64, name string, positionX float64, positionY float64) (*model.Station, error) {
	_, err := s.repository.GetStationByID(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("station with id %d not found", id))
	}

	updated, err := s.repository.UpdateStation(id, name, positionX, positionY)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to update station: %v", err))
	}

	event := micro.NewEvent(events.StationsStationUpdatedEventType, events.StationUpdatedEvent{
		ID:        dto.IDToDTO(updated.ID),
		Name:      updated.Name,
		PositionX: updated.PositionX,
		PositionY: updated.PositionY,
	})
	err = s.kafka.Producer().Send(events.StationsTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return updated, nil
}

// DeleteStation deletes the station with the given id.
func (s *StationService) DeleteStation(id uint64) (*model.Station, error) {
	station, err := s.repository.GetStationByID(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("station with id %d not found", id))
	}

	err = s.repository.DeleteStation(id)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to delete station: %v", err))
	}

	event := micro.NewEvent(events.StationsStationDeletedEventType, events.StationDeletedEvent{
		ID: dto.IDToDTO(station.ID),
	})
	err = s.kafka.Producer().Send(events.StationsTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return station, nil
}

// GetStationViews returns all the stations.
func (s *StationService) GetStationViews() ([]*model.StationView, error) {
	stations, err := s.viewRepository.GetStations()
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get stations: %v", err))
	}

	return stations, nil
}

// CreateStationView creates a new station with the given name and position.
func (s *StationService) CreateStationView(id uint64, name string, positionX float64, positionY float64) error {
	_, err := s.viewRepository.CreateStation(id, name, positionX, positionY)

	return err
}

// UpdateStationView updates the station with the given id.
func (s *StationService) UpdateStationView(id uint64, name string, positionX float64, positionY float64) error {
	_, err := s.viewRepository.UpdateStation(id, name, positionX, positionY)

	return err
}

// DeleteStationView deletes the station with the given id.
func (s *StationService) DeleteStationView(id uint64) error {
	return s.viewRepository.DeleteStation(id)
}
