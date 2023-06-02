package models

type Rental struct {
	RentalID    int64  `gorm:"primaryKey" json:"rental_id"`
	UserID      int64  `gorm:"foreignKey:UserID" json:"user_id"`
	CarID       int64  `gorm:"foreignKey:CarID" json:"car_id"`
	UsageRegion string `json:"usage_region"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
	TotalPrice  int64  `json:"total_price"`
	IsCompleted bool   `json:"is_completed"`
}
