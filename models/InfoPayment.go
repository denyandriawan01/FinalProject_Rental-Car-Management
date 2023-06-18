package models

type Payment struct {
	ID   		int 	`gorm:"column:payment_id;type:int;primaryKey;autoIncrement" 	json:"payment_id"`
	RentalID    int 	`gorm:"column:rental_id;type:int"								json:"rental_id"`
	Amount      int 	`gorm:"column:amount;type:int"									json:"amount"`
	PaymentDate string 	`gorm:"column:payment_date;type:varchar(50)"					json:"payment_date"`
	Rental      Rental 	`gorm:"foreignKey:RentalID"										json:"rental"`
}