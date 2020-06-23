package models

import "time"

// Car table
type Car struct {
	Plate       string `gorm:"primary_key"`
	CarTypeID   byte
	CarType     *CarType `gorm:"foreignkey:id;association_foreignkey:car_type_id"`
	TotalTime   int
	TotalAmount int
	CreatedAt   time.Time
}
