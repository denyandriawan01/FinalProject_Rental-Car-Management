package models

type MaintenanceHistory struct {
	MaintenanceID uint   `gorm:"primaryKey" json:"maintenance_id"`
	CarID         uint   `gorm:"foreignKey:CarID" json:"car_id"`
	LastOdometer  uint   `json:"last_odometer"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Expense       uint   `json:"expense"`
}
