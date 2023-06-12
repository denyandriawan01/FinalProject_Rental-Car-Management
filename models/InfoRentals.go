package models

type Rental struct {
	RentalID    int64  `gorm:"primaryKey" json:"rental_id"`
<<<<<<< HEAD
	UserID      int64  `json:"user_id"`
	CarID       int64  `json:"car_id"`
=======
	UserID      int64  `gorm:"column:user_id" json:"user_id"`
	CarID       int64  `gorm:"column:car_id" json:"car_id"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
	UsageRegion string `json:"usage_region"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
	TotalPrice  int64  `json:"total_price"`
	IsCompleted bool   `json:"is_completed"`
<<<<<<< HEAD
	User        User   `gorm:"references:UserID"`
	Car         Car    `gorm:"references:CarID"`
=======
	User	    User   `gorm:"foreignKey:UserID" json:"user"`
	Car   	    Car    `gorm:"foreignKey:CarID" json:"car"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
}
