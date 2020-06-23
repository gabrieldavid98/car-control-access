package data

import (
	"log"

	"github.com/gabrieldavid98/car-control-access/infrastructure/data/models"
	"github.com/jinzhu/gorm"
)

// Migrate applies migrations and creates default values
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.CarType{}, &models.Car{}, &models.Stay{})

	db.Create(&models.CarType{
		ID:   1,
		Name: "VIP",
		Fee:  0,
	})

	db.Create(&models.CarType{
		ID:   2,
		Name: "Residente",
		Fee:  50000,
	})

	db.Create(&models.CarType{
		ID:   3,
		Name: "No Residente",
		Fee:  50,
	})

	db.Model(&models.Car{}).
		AddForeignKey("car_type_id", "car_types(ID)", "RESTRICT", "RESTRICT")

	db.Model(&models.Stay{}).
		AddForeignKey("car_plate", "cars(plate)", "RESTRICT", "RESTRICT")

	log.Println("Migrations done!")
}
