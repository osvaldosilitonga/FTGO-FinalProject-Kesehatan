package repository

import (
	"errors"
	"payment/models/entity"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(d *entity.Payments) error
	Update(orderId string, status string, amount int) (*entity.Payments, error)
}

type PaymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{
		DB: db,
	}
}

func (p *PaymentRepositoryImpl) Save(d *entity.Payments) error {
	err := p.DB.Save(d).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PaymentRepositoryImpl) Update(orderId string, status string, amount int) (*entity.Payments, error) {
	var data entity.Payments
	err := p.DB.Where("order_id = ?", orderId).First(&data).Error
	if err != nil {
		return nil, errors.New("Data not found")
	}

	if status != "" {
		data.Status = status
	}
	if amount > 1000 {
		data.Amount = amount
	}

	err = p.DB.Save(&data).Error
	if err != nil {
		return nil, errors.New("Failed to update data")
	}

	return &data, nil
}
