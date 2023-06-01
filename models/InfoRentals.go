package models

type Rental struct {
	RentalID    uint   `gorm:"primaryKey" json:"rental_id"`
	UserID      uint   `gorm:"foreignKey:UserID" json:"user_id"`
	CarID       uint   `gorm:"foreignKey:CarID" json:"car_id"`
	UsageRegion string `json:"usage_region"`
	RentalDate  string `json:"rental_date"`
	ReturnDate  string `json:"return_date"`
	TotalPrice  uint   `json:"total_price"`
	IsCompleted bool   `json:"is_completed"`
}
