package database

import (
	"fmt"
	"log"

	"github.com/kweku-xvi/todo/api/v1/models"
	"github.com/kweku-xvi/todo/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.ENV.DBHost, config.ENV.DBUser, config.ENV.DBPassword, config.ENV.DBName, config.ENV.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database")
	}

	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("error running migrations")
	}

	fmt.Println("Connected to database and migrations applied")

}
