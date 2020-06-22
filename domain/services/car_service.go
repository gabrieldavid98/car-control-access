package services

import "github.com/gabrieldavid98/car-control-access/domain/models"

// CarService interface provides required methods
// for a payment service
type CarService interface {
	RegisterCar(car *models.Car) error
}
