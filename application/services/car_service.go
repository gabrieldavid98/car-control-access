package services

import (
	"time"

	validator "github.com/asaskevich/govalidator"
	"github.com/gabrieldavid98/car-control-access/domain/errors"
	"github.com/gabrieldavid98/car-control-access/domain/models"
	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	"github.com/gabrieldavid98/car-control-access/domain/services"
)

type carService struct {
	repo repositories.CarRepository
}

// NewCarService creates a new brand car service
func NewCarService(repo repositories.CarRepository) services.CarService {
	return &carService{repo}
}

func (c carService) RegisterCar(plate string, carTypeID byte) error {

	if len(plate) < 6 {
		return errors.ErrPlateLength
	}

	if !validator.IsAlphanumeric(plate) {
		return errors.ErrNotAlphaNumericPlate
	}

	carType, err := c.repo.GetCarType(carTypeID)
	if err != nil {
		return err
	}

	car := models.Car{
		Plate:       plate,
		CarType:     carType,
		TotalTime:   0,
		TotalAmount: carType.Fee,
		CreatedAt:   time.Now(),
	}

	return c.repo.RegisterCar(car)
}
