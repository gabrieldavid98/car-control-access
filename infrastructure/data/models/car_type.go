package models

// CarType table
type CarType struct {
	ID   byte `gorm:"primary_key;auto_increment"`
	Name string
	Fee  int
}
