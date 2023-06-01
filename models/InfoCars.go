package models

type Car struct {
	CarID              uint   `gorm:"primaryKey" json:"car_id"`
	Brand              string `json:"brand"`
	Model              string `json:"model"`
	Year               string `json:"year"`
	RegistrationNumber string `json:"registration_number"`
	VIN                string `json:"vin"`
	EngineNumber       string `json:"engine_number"`
	Color              string `json:"color"`
	IsAvailable        bool   `json:"is_available"`
}
