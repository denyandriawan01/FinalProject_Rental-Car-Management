package models

type Taxes struct {
	ID      	int 	`gorm:"column:tax_id;type:int;primaryKey;autoIncrement" 	json:"tax_id"`
	CarID 		int  	`gorm:"column:car_id;type:int"								json:"car_id"`
	TaxType    	string 	`gorm:"column:tax_type;type:varchar(255)"					json:"tax_type"`
	ValidUntil 	string 	`gorm:"column:valid_until;type:varchar(255)"				json:"valid_until"`
	Expense    	int 	`gorm:"column:expense;type:int" 							json:"expense"`
	Car        	Car    	`gorm:"foreignKey:CarID" 									json:"car"`
}