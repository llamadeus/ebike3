package globals

import "github.com/jmoiron/sqlx"

var (
	_db *sqlx.DB
)

func SetDB(db *sqlx.DB) {
	_db = db
}

func GetDB() *sqlx.DB {
	return _db
}
