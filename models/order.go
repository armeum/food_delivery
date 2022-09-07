package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	UserId             uint       `gorm:"column:user_id;foreignKey:id" json:"user_id"`
	DriverId           uint       `gorm:"column:driver_id;foreignKey:id" json:"driver_id"`
	Status             string     `gorm:"column:status" json:"status"`
	OrderCompletedTime *time.Time `gorm:"column:order_completed_time" json:"order_completed_time"`
	OrderInProgress    *time.Time `gorm:"column:order_in_progress" json:"order_in_progress"`
	OrderCanceledTime  *time.Time `gorm:"column:order_canceled_time" json:"order_canceled_time"`
}
