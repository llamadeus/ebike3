package events

const (
	AuthTopic                   = "auth"
	AuthUserRegisteredEventType = "UserRegistered"
)

type UserRegisteredEvent struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
