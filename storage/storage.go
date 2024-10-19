package storage

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/backend/storage/postgres"
	"gitlab.com/backend/storage/repo"
)

type StorageI interface {
	Clinic() repo.ClinicI
}

type StoragePg struct {
	Db         *sqlx.DB
	clinicRepo repo.ClinicI
}

// NewStoragePg
func NewStoragePg(db *sqlx.DB) *StoragePg {
	return &StoragePg{
		Db:         db,
		clinicRepo: postgres.NewClinicRepo(db),
	}
}

func (s StoragePg) Clinic() repo.ClinicI {
	return s.clinicRepo
}
