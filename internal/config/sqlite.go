package config

import (
	"os"

	"github.com/gaboliveiradev/api-jobs/internal/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")
	dbPath := "./internal/db/main.db"

	// check if the database file exists, if not create it
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("SQLite database file does not exist, creating it...")

		// create the database file and directory if it does not exist
		err = os.MkdirAll("./internal/db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect to SQLite database: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to auto-migrate SQLite database: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
