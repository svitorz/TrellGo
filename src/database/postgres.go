package database

import (
	"TrellGo/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	config.Load()
	return gorm.Open(postgres.Open(config.StringConnection), &gorm.Config{})
}
