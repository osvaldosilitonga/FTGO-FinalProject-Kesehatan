package repository

import (
	"errors"
	"payment/models/entity"
	"payment/models/web"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Save(d *entity.Payments) error
	FindByInvoiceID(id string) (*entity.Payments, error)
	FindByOrderID(id string) (*entity.Payments, error)
	FindByUserID(id, page int, status string) (*[]entity.Payments, error)
	Update(orderId string, d *entity.Payments) error

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

func (p *PaymentRepositoryImpl) FindByInvoiceID(id string) (*entity.Payments, error) {
	var data entity.Payments

	err := p.DB.Where("invoice_id = ?", id).First(&data).Error
	if err != nil {
		return nil, errors.New("Data not found")
	}

	return &data, nil
}

func (p *PaymentRepositoryImpl) FindByOrderID(id string) (*entity.Payments, error) {
	var data entity.Payments

	err := p.DB.Where("order_id = ?", id).First(&data).Error
	if err != nil {
		return nil, errors.New("Data not found")
	}

	return &data, nil
}

func (p *PaymentRepositoryImpl) FindByUserID(id, page int, status string) (*[]entity.Payments, error) {
	var data []entity.Payments

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * 10

	if status == "ALL" {
		err := p.DB.Order("updated_at desc").Where("user_id = ?", id).Offset(offset).Limit(10).Find(&data).Error
		if err != nil {
			return nil, errors.New("Data not found")
		}
	} else {
		err := p.DB.Order("updated_at desc").Where("user_id = ? AND status = ?", id, status).Offset(offset).Limit(10).Find(&data).Error
		if err != nil {
			return nil, errors.New("Data not found")
		}
	}

	return &data, nil
}

func (p *PaymentRepositoryImpl) Update(orderId string, d *entity.Payments) error {
	err := p.DB.Where("order_id = ?", orderId).Updates(d).Error
	if err != nil {
		return err
	}

	return nil
}
