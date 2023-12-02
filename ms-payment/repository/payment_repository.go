package repository

import (
	"errors"
	"payment/models/entity"
	"payment/models/web"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(d *entity.Payments) error
	UpdateFromXendit(d *web.XenditCallbackBody) (*entity.Payments, error)
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

func (p *PaymentRepositoryImpl) UpdateFromXendit(d *web.XenditCallbackBody) (*entity.Payments, error) {
	var data entity.Payments

	tx := p.DB.Begin()
	defer tx.Commit()

	err := tx.Where("order_id = ?", d.ExternalID).First(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Data not found")
	}

	data.Status = d.Status
	data.PaymentMethod = d.PaymentMethod
	data.MerchantName = d.MerchantName
	data.Currency = d.Currency
	data.UpdatedAt = d.Updated

	err = p.DB.Save(&data).Error
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Failed to update data")
	}

	return &data, nil
}
