package repositories

import (
	"time"

	"github.com/gabrieldavid98/car-control-access/domain/models"
	"github.com/gabrieldavid98/car-control-access/domain/repositories"
	dbmodels "github.com/gabrieldavid98/car-control-access/infrastructure/data/models"
	"github.com/jinzhu/gorm"
)

type accessRepository struct {
	db *gorm.DB
}

// NewAccessRepository creates a new brand access repository
func NewAccessRepository(db *gorm.DB) repositories.AccessRepository {
	return &accessRepository{db}
}

func (a *accessRepository) RegisterAccess(plate string) error {
	stay := &dbmodels.Stay{
		CarPlate:   plate,
		AccessTime: time.Now(),
	}

	return a.db.Create(stay).Error
}

func (a *accessRepository) RegisterExit(plate string) (*models.Stay, error) {
	var stay dbmodels.Stay
	var car dbmodels.Car

	err := a.db.Where("car_plate = ?").First(&stay).Error
	if err != nil {
		return nil, err
	}

	stay.ExitTime = time.Now()

	if err := a.db.Save(&stay).Error; err != nil {
		return nil, err
	}

	err = a.db.Preload("CarType").Where("plate = ?", plate).First(&car).Error
	if err != nil {
		return nil, err
	}

	result := &models.Stay{
		ID:         stay.ID,
		AccessTime: stay.AccessTime,
		ExitTime:   stay.ExitTime,
		Car: models.Car{
			CarType: models.CarType{
				ID:  car.CarType.ID,
				Fee: car.CarType.Fee,
			},
		},
	}

	return result, nil
}

func (a *accessRepository) UpdateTotalAmount(plate string, newAmount int) error {
	var car dbmodels.Car
	err := a.db.Where("plate = ?", plate).First(&car).Error
	if err != nil {
		return err
	}

	car.TotalAmount = newAmount

	return a.db.Save(&car).Error
}
