package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Stay table
type Stay struct {
	gorm.Model
	AccessTime time.Time
	ExitTime   time.Time
	Car        *Car   `gorm:"foreignkey:plate;association_foreignkey:car_plate"`
	CarPlate   string `gorm:"index:car_plate_index"`
}
