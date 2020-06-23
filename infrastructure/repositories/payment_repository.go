package repositories

import (
	"github.com/gabrieldavid98/car-control-access/domain/models"
	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	"github.com/jinzhu/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository creates a new brand payment repository
func NewPaymentRepository(db *gorm.DB) repositories.PaymentRepository {
	return &paymentRepository{db}
}

func (p *paymentRepository) GetTotalAmount(plate string) (int, error) {
	car, err := p.getCar(plate)
	return car.TotalAmount, err
}

func (p *paymentRepository) Pay(plate string, amount int) error {
	car, err := p.getCar(plate)
	if err != nil {
		return err
	}

	car.TotalAmount -= amount
	p.db.Save(car)

	return nil
}

func (p *paymentRepository) PaymentReport() ([]models.Stay, error) {
	panic("Write tests")
}

func (p *paymentRepository) getCar(plate string) (*models.Car, error) {
	var car models.Car
	db := p.db.Where("plate = ?", plate).First(&car)

	return &car, db.Error
}
