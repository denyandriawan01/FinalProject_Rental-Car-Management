package models

type Payment struct {
	PaymentID   uint   `gorm:"primaryKey" json:"payment_id"`
	RentalID    uint   `gorm:"foreignKey:RentalID" json:"rental_id"`
	Amount      uint   `json:"amount"`
	PaymentDate string `json:"payment_date"`
}
