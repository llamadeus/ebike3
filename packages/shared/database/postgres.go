package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Options struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func Open(options Options) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		options.Host,
		options.Port,
		options.User,
		options.Password,
		options.Database,
	)

	var db *sqlx.DB
	var err error
	start := time.Now()
	timeout := 10 * time.Second

	for time.Since(start) < timeout {
		db, err = sqlx.Connect("postgres", psqlInfo)
		if err == nil {
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sqlx.DB, schema string) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
