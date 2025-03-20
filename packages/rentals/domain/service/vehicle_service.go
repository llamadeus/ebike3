package service

import (
	"github.com/llamadeus/ebike3/packages/rentals/domain/model"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/out"
)

type VehicleService struct {
	viewRepository out.VehicleViewRepository
}

var _ in.VehicleService = (*VehicleService)(nil)

func NewVehicleService(viewRepository out.VehicleViewRepository) *VehicleService {
	return &VehicleService{
		viewRepository: viewRepository,
	}
}

func (s *VehicleService) GetVehicleByID(id uint64) (*model.VehicleView, error) {
	return s.viewRepository.Get(id)
}

func (s *VehicleService) CreateVehicleView(id uint64, type_ model.VehicleType) error {
	_, err := s.viewRepository.Create(id, type_)

	return err
}

func (s *VehicleService) DeleteVehicleView(id uint64) error {
	return s.viewRepository.Delete(id)
}
