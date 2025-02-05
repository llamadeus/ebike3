package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type User struct {
	ID        uint64    `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      UserRole  `db:"role"`
	LastLogin null.Time `db:"last_login"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UserRole string

const (
	UserRoleAdmin UserRole = "ADMIN"
	UserRoleUser  UserRole = "USER"
)
