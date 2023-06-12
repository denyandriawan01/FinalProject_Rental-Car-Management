package models

type MaintenanceHistory struct {
	MaintenanceID int64  `gorm:"primaryKey" json:"maintenance_id"`
<<<<<<< HEAD
	CarID         int64  `json:"car_id"`
=======
	CarID         int64  `gorm:"column:car_id" json:"car_id"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
	LastOdometer  int64  `json:"last_odometer"`
	Type          string `json:"type"`
	Description   string `json:"description"`
	Expense       int64  `json:"expense"`
<<<<<<< HEAD
	Car           Car    `gorm:"references:CarID"`
=======
	Car	      Car    `gorm:"foreignKey:CarID" json:"car"`
>>>>>>> 799eeb4513015e2f1a8022533d8158e006e5e25e
}
