package models

import (
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
	database.AutoMigrate(&Rental{})
	database.AutoMigrate(&MaintenanceHistory{})
	database.AutoMigrate(&Taxes{})
	database.AutoMigrate(&Payment{})
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Car{})

	DB = database
}