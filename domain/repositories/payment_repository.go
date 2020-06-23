package repositories

import "github.com/gabrieldavid98/car-control-access/domain/models"

// PaymentRepository interface
type PaymentRepository interface {
	GetTotalAmount(plate string) (int, error)
	Pay(plate string, amount int) error
	PaymentReport() ([]models.Stay, error)
}
