package services

import (
	"bytes"

	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	"github.com/gabrieldavid98/car-control-access/domain/services"
)

type paymentService struct {
	repo repositories.PaymentRepository
}

// NewPaymentService creates a new brand payment service
func NewPaymentService(repo repositories.PaymentRepository) services.PaymentService {
	return &paymentService{repo}
}

func (p *paymentService) GetTotalAmount(plate string) (int, error) {
	totalAmount, err := p.repo.GetTotalAmount(plate)
	if err != nil {
		return 0, err
	}

	return totalAmount, nil
}

func (p *paymentService) Pay(plate string, amount int) error {
	return p.repo.Pay(plate, amount)
}

func (p *paymentService) PaymentReport() (*bytes.Buffer, error) {
	panic("Not implemented")
}
