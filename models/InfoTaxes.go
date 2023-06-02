package models

type Taxes struct {
	TaxID      int64  `gorm:"primaryKey" json:"tax_id"`
	CarID      int64  `gorm:"foreignKey:CarID" json:"car_id"`
	TaxType    string `json:"tax_type"`
	ValidUntil string `json:"valid_until"`
	Expense    int64  `json:"expense"`
}
