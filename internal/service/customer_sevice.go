package service

import (
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/repository"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		return nil, &ServiceError{
			Message: "customer with this phone number already exists",
			Code:    400,
		}
	}

	request.Id = uuid.New().String()
	if err := s.customerRepo.CreateCustomer(&request); err != nil {
		return nil, err
	}
	logrus.Infof("Customer saved: %s", request.CustomerName)
	return &request, nil
}

func (s *customerService) GetCustomer(request model.RequestPagination) (*[]model.Customer, error) {
	if request.Page == "" {
		request.Page = "1"
	}

	if request.Limit == "" {
		request.Limit = "10"
	}

	customers, err := s.customerRepo.FindCustomers(request)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// func (s *customerService) UpdateCustomer(custId string, request *model.Customer) (*model.Customer, error) {
// 	customer, err := s.customerRepo.FindCustomerByID(custId)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, &ServiceError{Message: "customer with this id is not found", Code: 404}
// 		}
// 		return nil, err
// 	}

// }
