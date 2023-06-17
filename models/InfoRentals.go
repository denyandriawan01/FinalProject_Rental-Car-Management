package models

type Rental struct {
	ID    		int  	`gorm:"column:rental_id;type:int;primaryKey;autoIncrement" 		json:"rental_id"`
	UserID      int  	`gorm:"column:user_id;type:int"									json:"user_id"`
	CarID       int  	`gorm:"column:car_id;type:int"									json:"car_id"`
	UsageRegion string 	`gorm:"column:usage_region;type:varchar(50)"					json:"usage_region"`
	RentalDate  string 	`gorm:"column:rental_date;type:varchar(50)"						json:"rental_date"`
	ReturnDate  string 	`gorm:"column:return_date;type:varchar(50)"						json:"return_date"`
	TotalPrice  int 	`gorm:"column:total_price;type:int"								json:"total_price"`
	IsCompleted bool   	`gorm:"column:is_completed;type:boolean"						json:"is_completed"`
	User        User   	`gorm:"foreignKey:UserID"										json:"user"`
	Car         Car    	`gorm:"foreignKey:CarID"										json:"car"`
}