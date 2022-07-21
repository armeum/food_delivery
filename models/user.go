package models

import "time"

type User struct {
	Id          uint      `json:"id"`
	FirstName   string    `json:"name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

