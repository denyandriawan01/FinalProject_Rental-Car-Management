package models

type Car struct {
	ID              	int 	`gorm:"column:car_id;type:int;primaryKey;autoIncrement" 	json:"car_id"`
	Brand              	string 	`gorm:"column:brand;type:varchar(50)" 						json:"brand"`
	Model              	string 	`gorm:"column:model;type:varchar(50)" 						json:"model"`
	Year               	int 	`gorm:"column:year;type:int" 								json:"year"`
	RegistrationNumber 	string 	`gorm:"column:registration_number;type:varchar(50)" 		json:"registration_number"`
	VIN                	string 	`gorm:"column:vin;type:varchar(50)" 						json:"vin"`
	EngineNumber       	string 	`gorm:"column:engine_number;type:varchar(50)" 				json:"engine_number"`
	Color              	string	`gorm:"column:color;type:varchar(50)" 						json:"color"`
	IsAvailable        	bool  	`gorm:"column:is_available;type:boolean" 					json:"is_available"`
}