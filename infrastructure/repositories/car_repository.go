package repositories

import (
	"github.com/gabrieldavid98/car-control-access/domain/models"
	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	dbmodels "github.com/gabrieldavid98/car-control-access/infrastructure/data/models"
	"github.com/jinzhu/gorm"
)

type carRepository struct {
	db *gorm.DB
}

// NewCarRepository creates a new brand car repository
func NewCarRepository(db *gorm.DB) repositories.CarRepository {
	return &carRepository{db}
}

func (c *carRepository) RegisterCar(car models.Car) error {
	newCar := &dbmodels.Car{
		Plate:       car.Plate,
		CarTypeID:   car.CarType.ID,
		TotalTime:   car.TotalTime,
		TotalAmount: car.TotalAmount,
	}

	db := c.db.Create(newCar)
	return db.Error
}

func (c *carRepository) GetCarType(carTypeID byte) (models.CarType, error) {
	carType := &dbmodels.CarType{}
	db := c.db.Find(carType, carTypeID)

	carTypeResult := models.CarType{
		ID:   carType.ID,
		Name: carType.Name,
		Fee:  carType.Fee,
	}

	return carTypeResult, db.Error
}
