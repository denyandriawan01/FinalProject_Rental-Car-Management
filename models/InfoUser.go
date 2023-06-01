package models

type User struct {
	UserID      uint   `gorm:"primaryKey" json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}
