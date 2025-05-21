package config

import (
	"os"

	"github.com/whitestudios/TaskFlowAPI/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqlite() (*gorm.DB, error) {
	logger := NewLogger("sqlite")
	dbDir := "./db"
	dbPath := "./db/task-api.db"

	// check if db exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create db directory
		err = os.MkdirAll(dbDir, os.ModePerm)

		if err != nil {
			logger.Err("File to create db directory")
			return nil, err
		}

		// Create db file
		file, err := os.Create(dbPath)

		if err != nil {
			logger.Err("File to create db file")
			return nil, err
		}

		file.Close()
	}

	// Connect with DB
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errf("Sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Task{}, &models.User{})

	if err != nil {
		logger.Errf("Error creating task and user model: %v", err)
		return nil, err
	}

	return db, nil
}
