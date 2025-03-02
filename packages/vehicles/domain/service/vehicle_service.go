package service

import (
	"fmt"
	"github.com/llamadeus/ebike3/packages/vehicles/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/events"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/model"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/in"
	"github.com/llamadeus/ebike3/packages/vehicles/domain/port/out"
	"github.com/llamadeus/ebike3/packages/vehicles/infrastructure/micro"
)

// VehicleService implements the VehicleService interface.
type VehicleService struct {
	kafka          micro.Kafka
	repository     out.VehicleRepository
	viewRepository out.VehicleViewRepository
}

// Ensure that VehicleService implements the VehicleService interface.
var _ in.VehicleService = (*VehicleService)(nil)

// NewVehicleService creates a new instance of the VehicleService.
func NewVehicleService(kafka micro.Kafka, repository out.VehicleRepository, viewRepository out.VehicleViewRepository) *VehicleService {
	return &VehicleService{
		kafka:          kafka,
		repository:     repository,
		viewRepository: viewRepository,
	}
}

// CreateVehicle creates a new vehicle with the given type, position, and battery.
func (s *VehicleService) CreateVehicle(type_ model.VehicleType, positionX float64, positionY float64, battery float64) (*model.Vehicle, error) {
	vehicle, err := s.repository.CreateVehicle(type_, positionX, positionY, battery)
	if err != nil {
		return nil, err
	}

	event := micro.NewEvent(events.VehiclesVehicleCreatedEventType, events.VehicleCreatedEvent{
		ID:        dto.IDToDTO(vehicle.ID),
		Type:      dto.TypeToDTO(vehicle.Type),
		PositionX: vehicle.PositionX,
		PositionY: vehicle.PositionY,
		Battery:   vehicle.Battery,
		Available: true,
	})
	err = s.kafka.Producer().Send(events.VehiclesTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return vehicle, nil
}

// UpdateVehicle updates the vehicle with the given id.
func (s *VehicleService) UpdateVehicle(id uint64, positionX float64, positionY float64, battery float64) (*model.Vehicle, error) {
	_, err := s.repository.GetVehicleByID(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("vehicle with id %d not found", id))
	}

	updated, err := s.repository.UpdateVehicle(id, positionX, positionY, battery)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to update vehicle: %v", err))
	}

	event := micro.NewEvent(events.VehiclesVehicleUpdatedEventType, events.VehicleUpdatedEvent{
		ID:        dto.IDToDTO(updated.ID),
		PositionX: updated.PositionX,
		PositionY: updated.PositionY,
		Battery:   updated.Battery,
	})
	err = s.kafka.Producer().Send(events.VehiclesTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return updated, nil
}

// DeleteVehicle deletes the vehicle with the given id.
func (s *VehicleService) DeleteVehicle(id uint64) (*model.Vehicle, error) {
	vehicle, err := s.repository.GetVehicleByID(id)
	if err != nil {
		return nil, micro.NewNotFoundError(fmt.Sprintf("vehicle with id %d not found", id))
	}

	err = s.repository.DeleteVehicle(id)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to delete vehicle: %v", err))
	}

	event := micro.NewEvent(events.VehiclesVehicleDeletedEventType, events.VehicleDeletedEvent{
		ID: dto.IDToDTO(vehicle.ID),
	})
	err = s.kafka.Producer().Send(events.VehiclesTopic, event.Payload.ID, event)
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to send kafka event: %v", err))
	}

	return vehicle, nil
}

// GetVehicleViews returns all the vehicles.
func (s *VehicleService) GetVehicleViews() ([]*model.VehicleView, error) {
	vehicles, err := s.viewRepository.GetVehicles()
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get vehicles: %v", err))
	}

	return vehicles, nil
}

// GetAvailableVehicleViews returns all the available vehicles.
func (s *VehicleService) GetAvailableVehicleViews() ([]*model.VehicleView, error) {
	vehicles, err := s.viewRepository.GetAvailableVehicles()
	if err != nil {
		return nil, micro.NewInternalServerError(fmt.Sprintf("failed to get vehicles: %v", err))
	}

	return vehicles, nil
}

// CreateVehicleView creates a new vehicle with the given type, position, and battery.
func (s *VehicleService) CreateVehicleView(id uint64, type_ model.VehicleType, positionX float64, positionY float64, battery float64) error {
	_, err := s.viewRepository.CreateVehicle(id, type_, positionX, positionY, battery)
	if err != nil {
		return micro.NewInternalServerError(fmt.Sprintf("failed to create vehicle: %v", err))
	}

	return nil
}

// UpdateVehicleView updates the vehicle with the given id.
func (s *VehicleService) UpdateVehicleView(id uint64, positionX float64, positionY float64, battery float64) error {
	_, err := s.viewRepository.UpdateVehicle(id, positionX, positionY, battery)
	if err != nil {
		return micro.NewInternalServerError(fmt.Sprintf("failed to update vehicle: %v", err))
	}

	return nil
}

// UpdateVehicleViewAvailability updates the availability of the vehicle with the given id.
func (s *VehicleService) UpdateVehicleViewAvailability(id uint64, available bool) error {
	_, err := s.viewRepository.UpdateVehicleAvailability(id, available)
	if err != nil {
		return micro.NewInternalServerError(fmt.Sprintf("failed to update vehicle availability: %v", err))
	}

	return nil
}

// DeleteVehicleView deletes the vehicle with the given id.
func (s *VehicleService) DeleteVehicleView(id uint64) error {
	err := s.viewRepository.DeleteVehicle(id)
	if err != nil {
		return micro.NewInternalServerError(fmt.Sprintf("failed to delete vehicle: %v", err))
	}

	return nil
}
