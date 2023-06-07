package models

type Payment struct {
	PaymentID   int64  `gorm:"primaryKey" json:"payment_id"`
	RentalID    int64  `gorm:"column:rental_id" json:"rental_id"`
	Amount      int64  `json:"amount"`
	PaymentDate string `json:"payment_date"`
	Rental	    Rental `gorm:"foreignKey:RentalID" json:"rental"`
}
