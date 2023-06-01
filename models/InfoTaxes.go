package models

type Taxes struct {
	TaxID      uint   `gorm:"primaryKey" json:"tax_id"`
	CarID      uint   `gorm:"foreignKey:CarID" json:"car_id"`
	TaxType    string `json:"tax_type"`
	ValidUntil string `json:"valid_until"`
	Expense    uint   `json:"expense"`
}
