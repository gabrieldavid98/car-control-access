package models

// CarType model
type CarType struct {
	ID   byte   `json:"id"`
	Name string `json:"name"`
	Fee  int    `json:"fee"`
}
