package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zerodev/golang_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Create new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&&charset=utf8mb4&collation=utf8mb4_unicode_ci", dbUser, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Label{}, &entity.Task{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
