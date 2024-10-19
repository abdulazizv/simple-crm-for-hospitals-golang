package postgres

import (
	"github.com/jmoiron/sqlx"
)

type storageRepo struct {
	db *sqlx.DB
}

func NewClinicRepo(db *sqlx.DB) *storageRepo {
	return &storageRepo{
		db: db,
	}
}
