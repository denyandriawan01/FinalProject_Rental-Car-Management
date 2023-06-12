package models

type Taxes struct {
	TaxID      int64  `gorm:"primaryKey" json:"tax_id"`
	CarID      int64  `gorm:"column:car_id" json:"car_id"`
	TaxType    string `json:"tax_type"`
	ValidUntil string `json:"valid_until"`
	Expense    int64  `json:"expense"`
<<<<<<< HEAD
	Car        Car    `gorm:"references:CarID"`
=======
	Car 	   Car    `gorm:"foreignKey:CarID" json:"car"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
}
