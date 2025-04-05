package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeRentalsChargeActiveRental = "rentals:charge-active-rental"
)

type RentalsChargeActiveRentalPayload struct {
	RentalID string `json:"rentalId"`
}

func NewRentalsChargeActiveRentalTask(rentalID string) (*asynq.Task, error) {
	payload, err := json.Marshal(RentalsChargeActiveRentalPayload{RentalID: rentalID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeRentalsChargeActiveRental, payload), nil
}
