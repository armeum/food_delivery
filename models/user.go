package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `gorm:"typevarchar(50);unique_index" json:"phone_number"`
	Email       string `gorm:"typevarchar(100);unique_index" json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
	Cart        []Cart `gorm:"ForeignKey:UserID" json:"cart"`
}
