package db

import (
	"github.com/Mogza/AstralGate/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Init() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	utils.LogFatal(err, "Failed to connect to database")

	err = db.AutoMigrate()
	utils.LogFatal(err, "Failed to migrate database")

	return db
}
