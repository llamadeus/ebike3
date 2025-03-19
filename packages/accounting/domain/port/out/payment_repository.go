package out

import "github.com/llamadeus/ebike3/packages/accounting/domain/model"

type PaymentRepository interface {
	Get(id uint64) (*model.Payment, error)

	GetAll() ([]*model.Payment, error)

	GetByCustomerID(customerID uint64) ([]*model.Payment, error)

	Create(customerID uint64, amount int) (*model.Payment, error)

	Update(id uint64, status model.PaymentStatus) (*model.Payment, error)

	Delete(id uint64) error
}
