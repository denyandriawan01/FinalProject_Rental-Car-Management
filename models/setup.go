package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/finpro_golang"))
	// Membuka koneksi database menggunakan driver MySQL dengan parameter koneksi: "root:@tcp(localhost:3306)/final_project_golang"
	if err != nil {
		panic(err) // Jika terjadi kesalahan saat membuka koneksi, lempar error dan hentikan program
	}

	database.AutoMigrate(&MaintenanceHistory{})
	database.AutoMigrate(&Taxes{})
	database.AutoMigrate(&User{})
	database.AutoMigrate(&Car{})
	database.AutoMigrate(&Rental{})
	database.AutoMigrate(&Payment{})

	DB = database
}
