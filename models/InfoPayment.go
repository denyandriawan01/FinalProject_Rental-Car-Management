package models

type Payment struct {
	PaymentID   int64  `gorm:"primaryKey" json:"payment_id"`
<<<<<<< HEAD
	RentalID    int64  `json:"rental_id"`
	Amount      int64  `json:"amount"`
	PaymentDate string `json:"payment_date"`
	Rental      Rental `gorm:"references:RentalID"`
=======
	RentalID    int64  `gorm:"column:rental_id" json:"rental_id"`
	Amount      int64  `json:"amount"`
	PaymentDate string `json:"payment_date"`
	Rental	    Rental `gorm:"foreignKey:RentalID" json:"rental"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
}
