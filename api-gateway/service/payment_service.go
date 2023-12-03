package service

type Payment interface {
	CreatePayment() error
}

type PaymentImpl struct {
}

func NewPaymentService() Payment {
	return &PaymentImpl{}
}

func (p *PaymentImpl) CreatePayment() error {
	return nil
}
