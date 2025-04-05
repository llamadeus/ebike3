package events

const (
	VehiclesTopic                   = "vehicles"
	VehiclesVehicleCreatedEventType = "VehicleCreated"
	VehiclesVehicleUpdatedEventType = "VehicleUpdated"
	VehiclesVehicleDeletedEventType = "VehicleDeleted"
)

type VehicleCreatedEvent struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
	Battery   float64 `json:"battery"`
	Available bool    `json:"available"`
}

type VehicleUpdatedEvent struct {
	ID        string  `json:"id"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
	Battery   float64 `json:"battery"`
}

type VehicleDeletedEvent struct {
	ID string `json:"id"`
}
