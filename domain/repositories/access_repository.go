package repositories

import "github.com/gabrieldavid98/car-control-access/domain/models"

// AccessRepository interface
type AccessRepository interface {
	RegisterAccess(plate string) error
	RegisterExit(plate string) (*models.Stay, error)
	UpdateTotalAmount(plate string, newAmount int) error
}
