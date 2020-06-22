package repositories

import "github.com/gabrieldavid98/car-control-access/domain/models"

// PaymentRepository interface
type PaymentRepository interface {
	GetTotalAmount(plate string) (float64, error)
	Pay(plate string, amount float64) error
	PaymentReport() (*[]models.PaymentReport, error)
}
