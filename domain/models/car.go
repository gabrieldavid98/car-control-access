package models

import "time"

// Car model
type Car struct {
	Plate     string    `json:"plate" validate:"eq=6 & format=alnum"`
	Type      CarType   `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}
