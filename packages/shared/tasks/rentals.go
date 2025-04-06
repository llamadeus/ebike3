package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const (
	TypeRentalsChargeActiveRental = "rentals:charge-active-rental"
)

type RentalsChargeActiveRentalPayload struct {
	RentalID  string    `json:"rentalId"`
	Timestamp time.Time `json:"timestamp"`
}

func NewRentalsChargeActiveRentalTask(rentalID string, timestamp time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(RentalsChargeActiveRentalPayload{RentalID: rentalID, Timestamp: timestamp})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeRentalsChargeActiveRental, payload), nil
}
