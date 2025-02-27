package model

import (
	"github.com/guregu/null/v5"
	"time"
)

type Session struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt null.Time `db:"deleted_at"`
}
