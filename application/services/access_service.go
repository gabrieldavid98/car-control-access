package services

import (
	"github.com/gabrieldavid98/car-control-access/domain/models"
	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	"github.com/gabrieldavid98/car-control-access/domain/services"
)

type accessService struct {
	repo repositories.AccessRepository
}

// NewAccessService creates a new brand access service
func NewAccessService(repo repositories.AccessRepository) services.AccessService {
	return &accessService{repo}
}

func (a *accessService) RegisterAccess(plate string) error {
	return a.repo.RegisterAccess(plate)
}

func (a *accessService) RegisterExit(plate string) error {
	stay, err := a.repo.RegisterExit(plate)
	if err != nil {
		return err
	}

	if models.CarTypeID(stay.Car.CarType.ID) != models.NoResident {
		return nil
	}

	mins := stay.ExitTime.Sub(stay.AccessTime).Minutes()
	newAmount := stay.Car.CarType.Fee * int(mins)

	err = a.repo.UpdateTotalAmount(plate, newAmount)
	if err != nil {
		return err
	}

	return nil
}
