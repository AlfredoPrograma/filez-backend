// Package database exposes connection functions.
package database

import (
	"fmt"

	"github.com/alfredoprograma/filez-server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect establishes a connection with database using the provided parameters through config.
func Connect(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(); err != nil {
		panic(err)
	}

	return db
}
