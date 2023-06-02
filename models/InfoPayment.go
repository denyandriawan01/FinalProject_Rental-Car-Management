package models

type Payment struct {
	PaymentID   int64  `gorm:"primaryKey" json:"payment_id"`
	RentalID    int64  `gorm:"foreignKey:RentalID" json:"rental_id"`
	Amount      int64  `json:"amount"`
	PaymentDate string `json:"payment_date"`
}
