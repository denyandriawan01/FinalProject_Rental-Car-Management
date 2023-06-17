package models

type MaintenanceHistory struct {
	ID 				int  	`gorm:"column:maintenance_id;type:int;primaryKey;autoIncrement" 	json:"maintenance_id"`
	CarID         	int  	`gorm:"column:car_id;type:int"										json:"car_id"`
	LastOdometer  	int  	`gorm:"column:last_odometer;type:int"								json:"last_odometer"`
	Type          	string 	`gorm:"column:type;type:varchar(50)"								json:"type"`
	Description   	string 	`gorm:"column:description;type:varchar(255)"						json:"description"`
	Expense       	int  	`gorm:"column:expense;type:int"										json:"expense"`
	Car           	Car    	`gorm:"foreignKey:CarID" 											json:"car"`
}