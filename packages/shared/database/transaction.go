package database

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

type Transactor interface {
	Beginx() (*sqlx.Tx, error)
}

func RunInTx(transactor Transactor, fn func(tx *sqlx.Tx) error) error {
	tx, err := transactor.Beginx()
	if err != nil {
		return err
	}

	err = fn(tx)
	if err == nil {
		return tx.Commit()
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
