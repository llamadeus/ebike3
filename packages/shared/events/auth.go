package events

import "time"

const (
	AuthTopic                   = "auth"
	AuthUserRegisteredEventType = "UserRegistered"
	AuthUserLoggedInEventType   = "UserLoggedIn"
)

type UserRegisteredEvent struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserLoggedInEvent struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Timestamp time.Time `json:"timestamp"`
}
