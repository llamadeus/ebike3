package in

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/llamadeus/ebike3/packages/rentals/adapter/out/dto"
	"github.com/llamadeus/ebike3/packages/rentals/domain/port/in"
	"github.com/llamadeus/ebike3/packages/rentals/domain/tasks"
	"log/slog"
)

func MakeRentalTasksProcessor(mux *asynq.ServeMux, rentalService in.RentalService) {
	mux.HandleFunc(tasks.TypeRentalsChargeActiveRental, func(ctx context.Context, task *asynq.Task) error {
		var payload tasks.RentalsChargeActiveRentalPayload
		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
		}

		slog.Info(
			"charge active rental",
			"id", payload.RentalID,
		)

		rentalID, err := dto.IDFromDTO(payload.RentalID)
		if err != nil {
			return err
		}

		return rentalService.ChargeActiveRental(ctx, rentalID, payload.Timestamp)
	})
}
