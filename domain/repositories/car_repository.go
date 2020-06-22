package repositories

import "github.com/gabrieldavid98/car-control-access/domain/models"

// CarRepository interface
type CarRepository interface {
	RegisterCar(car *models.Car)
}
