package models

import "time"

// Car model
type Car struct {
	Plate       string  `json:"plate"`
	CarType     CarType `json:"type"`
	TotalTime   int     `json:"totalTime"`
	TotalAmount int
	CreatedAt   time.Time `json:"createdAt"`
}
