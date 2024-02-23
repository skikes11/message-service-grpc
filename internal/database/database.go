package database

import (
	"fmt"
	"log"
	"message-service/initializers"
	"message-service/internal/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB(config *initializers.Config) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             250 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false, // Don't include params in the SQL log
			Colorful:                  true,  // Enable color
		},
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	if err := DB.AutoMigrate(
		&models.Message{},
	); err != nil {
		panic(err)
	}

	fmt.Println("🚀 Connected Successfully to the Database")

	return DB
}
