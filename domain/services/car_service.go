package services

// CarService interface provides required methods
// for a payment service
type CarService interface {
	RegisterCar(plate string, carType byte) error
}
