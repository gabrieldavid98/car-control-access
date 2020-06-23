package mysql

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetDB gets the db
func GetDB(connString string) *gorm.DB {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
