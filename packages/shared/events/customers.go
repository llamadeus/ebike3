package events

const (
	CustomersTopic                            = "vehicles"
	CustomersCustomerPositionUpdatedEventType = "CustomerPositionUpdated"
)

type CustomerPositionUpdatedEvent struct {
	ID        string  `json:"id"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
}
