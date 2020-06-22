package models

import "time"

// Stay model
type Stay struct {
	ID         int       `json:"id"`
	AccessTime time.Time `json:"enterTime"`
	ExitTime   time.Time `json:"exitTime"`
	Car        Car       `json:"car"`
}
