package models

type Car struct {
	CarID              int64  `gorm:"primaryKey" json:"car_id"`
	Brand              string `gorm:"type:varchar(50)" json:"brand"`
	Model              string `gorm:"type:varchar(50)" json:"model"`
	Tipe               string `gorm:"type:varchar(50)" json:"Tipe"`
	Capasity           string `gorm:"type:varchar(50)" json:"Capacity"`
	Year               int64  `gorm:"int64" json:"year"`
	RegistrationNumber string `gorm:"type:varchar(50)" json:"registration_number"`
	VIN                string `gorm:"type:varchar(50)" json:"vin"`
	EngineNumber       string `gorm:"type:varchar(50)" json:"engine_number"`
	Color              string `gorm:"type:varchar(50)" json:"color"`
	IsAvailable        bool   `gorm:"type:boolean" json:"is_available"`
}
