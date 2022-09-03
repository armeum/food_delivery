package models

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FirstName string `json:"first_name" binding:"required"`
	PhoneNumber string `gorm:"type:varchar(9);unique_index" json:"phone_number" binding:"required"`
	Gender      GenderType `json:"gender" binding:"ENUM=GenderType"`
	// Email       string `gorm:"typevarchar(100);unique_index" json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Password string   `json:"password" binding:"required"`
}


type GenderType string

const (
	Male GenderType = "мужской"
	Female GenderType = "женский"
)