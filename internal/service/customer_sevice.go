package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/repository"
)

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (s *customerService) CreateCustomer(request model.Customer) (*model.Customer, error) {
	// check if customer already exists by phone number
	existing, _ := s.customerRepo.FindCustomerByPhoneNumber(request.PhoneNumber)
	if existing != nil {
		return nil, &ServiceError{Message: "customer with this phone number already exists", Code: 400}
	}

	request.Id = uuid.New().String()
	if err := s.customerRepo.CreateCustomer(&request); err != nil {
		return nil, err
	}
	logrus.Infof("Customer saved: %s (%s)", request.CustomerName)
	return &request, nil
}
