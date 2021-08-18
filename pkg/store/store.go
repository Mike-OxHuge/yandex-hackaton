package store

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func New() *Store {
	_db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost password=PASSWORD user=posgtres sslmode=disable database=yandex",
	}))
	if err != nil {
		return nil
	}
	_db.AutoMigrate()
	return &Store{
		DB: _db,
	}
}
