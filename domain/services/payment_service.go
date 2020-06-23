package services

import "bytes"

// PaymentService interface provides required methods
// for a payment service
type PaymentService interface {
	GetTotalAmount(plate string) (int, error)
	Pay(plate string, amount int) error
	PaymentReport() (*bytes.Buffer, error)
}
