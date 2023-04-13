package config

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=developmentuser1 password=password1 dbname=pos-online port=5432 sslmode=disable",
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// Set the maximum number of idle connections to the database.
	sqlDB.SetMaxIdleConns(5)

	// Set the maximum time a connection can be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
