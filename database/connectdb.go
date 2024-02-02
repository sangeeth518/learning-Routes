package database

import (
	"log"

	"os"

	"github.com/sangeeth518/learning-routes/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectTODb() {
	var err error

	dsn := os.Getenv("dsn")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to DB")
	}
	log.Println("Database connected succesfully")

	DB.AutoMigrate(&models.User{})
}
