package database

import (
	"FinalProject_Rental-Car-Management/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DB")))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&models.Rental{})
	database.AutoMigrate(&models.MaintenanceHistory{})
	database.AutoMigrate(&models.Taxes{})
	database.AutoMigrate(&models.Payment{})
	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Car{})

	DB = database
}
