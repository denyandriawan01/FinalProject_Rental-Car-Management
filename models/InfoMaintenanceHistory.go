package models

type MaintenanceHistory struct {
	MaintenanceID int64  `gorm:"primaryKey" json:"maintenance_id"`
	CarID         int64  `gorm:"column:car_id" json:"car_id"`
	LastOdometer  int64  `json:"last_odometer"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Expense       int64  `json:"expense"`
	Car	      Car    `gorm:"foreignKey:CarID" json:"car"`
}
