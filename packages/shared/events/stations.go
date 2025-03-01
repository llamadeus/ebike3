package events

const (
	StationsTopic                   = "stations"
	StationsStationCreatedEventType = "StationCreated"
	StationsStationUpdatedEventType = "StationUpdated"
	StationsStationDeletedEventType = "StationDeleted"
)

type StationCreatedEvent struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
}

type StationUpdatedEvent struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
}

type StationDeletedEvent struct {
	ID string `json:"id"`
}
