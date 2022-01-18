package service

import (
	"go-clean-architecture/pkg/entity"
	customerRepo "go-clean-architecture/pkg/repository/customer"
)

func NewCustomerService(customerRepository *customerRepo.CustomerRepo) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: *customerRepository,
	}
}

type CustomerServiceImpl struct {
	CustomerRepository customerRepo.CustomerRepo
}

func (service *CustomerServiceImpl) GetAll() []entity.Customer {
	customers := service.CustomerRepository.GetAll()
	return customers
}

func (service *CustomerServiceImpl) Get(id string) (entity.Customer, error) {
	customer, err := service.CustomerRepository.Get(id)
	return customer, err
}
