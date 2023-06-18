package database

import (
	"finpro_golang/models"
	"finpro_golang/utils/initializer"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(initializer.DB_CONN))
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