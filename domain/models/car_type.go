package models

// CarTypeID type
type CarTypeID byte

const (
	// Vip car type id
	Vip CarTypeID = iota + 1

	// Resident car type id
	Resident

	// NoResident car type id
	NoResident
)

// CarType model
type CarType struct {
	ID   byte   `json:"id"`
	Name string `json:"name"`
	Fee  int    `json:"fee"`
}
