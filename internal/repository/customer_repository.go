package repository

import (
	"errors"
	"sim-clinic-api/internal/model"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var tagCustomerRepository = "internal.repository.customer_repository."

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

func (r *customerRepository) UpdateCustomer(customer *model.Customer) error {
	return r.db.Save(customer).Error
}

func (r *customerRepository) FindCustomers(req model.RequestPagination) (*[]model.Customer, error) {
	var (
		tag                 = tagCustomerRepository + "FindCustomers."
		mCustomers          []model.Customer
		offset, page, limit int
	)

	page = cast.ToInt(req.Page)
	limit = cast.ToInt(req.Limit)

	offset = (page - 1) * limit

	queryBuilder := r.db.Model(&model.Customer{}).Limit(limit).Offset(offset)
	if err := queryBuilder.Find(&mCustomers).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "01",
			"error": err.Error(),
		})
		return nil, err
	}

	if len(mCustomers) == 0 {
		logrus.WithFields(logrus.Fields{
			"tag":   tag + "02",
			"error": "data customer not found",
		})
		return nil, errors.New("RESOURCE_NOT_FOUND")
	}

	return &mCustomers, nil
}
