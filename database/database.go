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

	cars := []models.Car{
		{
			Brand:              "Toyota",
			Model:              "Avanza",
			Tipe:               "MPV",
			Capasity:           "7 seats",
			Year:               2019,
			RegistrationNumber: "B 1234 ABC",
			VIN:                "JTDKN3DU4A3000001",
			EngineNumber:       "K3VE123456",
			Color:              "Silver",
			IsAvailable:        true,
		},
		{
			Brand:              "Honda",
			Model:              "City",
			Tipe:               "Sedan",
			Capasity:           "5 seats",
			Year:               2020,
			RegistrationNumber: "B 5678 DEF",
			VIN:                "JHMZF1D68CS100002",
			EngineNumber:       "L15Z123456",
			Color:              "White",
			IsAvailable:        true,
		},
		{
			Brand:              "Daihatsu",
			Model:              "Terios",
			Tipe:               "SUV",
			Capasity:           "5 seats",
			Year:               2018,
			RegistrationNumber: "B 9012 GHI",
			VIN:                "JDA1111113A000003",
			EngineNumber:       "3SZ123456",
			Color:              "Black",
			IsAvailable:        true,
		},
		{
			Brand:              "Suzuki",
			Model:              "Ertiga",
			Tipe:               "MPV",
			Capasity:           "7 seats",
			Year:               2021,
			RegistrationNumber: "B 3456 JKL",
			VIN:                "KL5TN6AHHKC000004",
			EngineNumber:       "K15B123456",
			Color:              "Red",
			IsAvailable:        true,
		},
		{
			Brand:              "Mitsubishi",
			Model:              "Xpander",
			Tipe:               "MPV",
			Capasity:           "7 seats",
			Year:               2022,
			RegistrationNumber: "B 7890 MNO",
			VIN:                "MMT111111BB000005",
			EngineNumber:       "4A91ABC123",
			Color:              "Grey",
			IsAvailable:        true,
		},
		{
			Brand:              "Toyota",
			Model:              "Innova",
			Tipe:               "MPV",
			Capasity:           "8 seats",
			Year:               2017,
			RegistrationNumber: "B 2345 PQR",
			VIN:                "KUN40A00000000006",
			EngineNumber:       "2KD567890",
			Color:              "Silver",
			IsAvailable:        true,
		},
		{
			Brand:              "Toyota",
			Model:              "Hiace",
			Tipe:               "Van",
			Capasity:           "12 seats",
			Year:               2020,
			RegistrationNumber: "B 6789 STU",
			VIN:                "JTF111111G3000007",
			EngineNumber:       "1GD123456",
			Color:              "White",
			IsAvailable:        true,
		},
		{
			Brand:              "Honda",
			Model:              "CR-V",
			Tipe:               "SUV",
			Capasity:           "5 seats",
			Year:               2019,
			RegistrationNumber: "B 1234 VWX",
			VIN:                "5J6RW1H82KA100008",
			EngineNumber:       "R20A123456",
			Color:              "Blue",
			IsAvailable:        true,
		},
		{
			Brand:              "Nissan",
			Model:              "Grand Livina",
			Tipe:               "MPV",
			Capasity:           "7 seats",
			Year:               2016,
			RegistrationNumber: "B 5678 YZ",
			VIN:                "MMNAFAD123FN000009",
			EngineNumber:       "HR15DE12345",
			Color:              "Black",
			IsAvailable:        true,
		},
		{
			Brand:              "Isuzu",
			Model:              "MU-X",
			Tipe:               "SUV",
			Capasity:           "7 seats",
			Year:               2023,
			RegistrationNumber: "B 9012 ABC",
			VIN:                "MPA1234000U000010",
			EngineNumber:       "4JJ123456",
			Color:              "Grey",
			IsAvailable:        true,
		},
	}

	for _, car := range cars {
		database.Create(&car)
	}

	DB = database
}
