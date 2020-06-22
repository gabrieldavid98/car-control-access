package services

import "bytes"

// PaymentService interface provides required methods
// for a payment service
type PaymentService interface {
	GetTotalAmount(plate string) (float64, error)
	Pay(plate string, amount float64) error
	PaymentReport() (*bytes.Buffer, error)
}
