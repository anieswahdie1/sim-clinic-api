package repository

import (
	"gorm.io/gorm"
	"sim-clinic-api/internal/model"
)

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) CreateCustomer(customer *model.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepository) FindCustomerByID(id string) (*model.Customer, error) {
	var customer []model.Customer
	err := r.db.Where("id = ?", id).Find(&customer).Error
	if err != nil {
		return nil, err
	}
	if len(customer) < 1 {
		return nil, nil
	}
	return &customer[0], nil
}

func (r *customerRepository) FindCustomerByPhoneNumber(phoneNumber string) (*model.Customer, error) {
	var customer []model.Customer
	err := r.db.Where("phone_number = ?", phoneNumber).Find(&customer).Error
	if err != nil {
		return nil, err
	}
	if len(customer) < 1 {
		return nil, nil
	}
	return &customer[0], nil
}
