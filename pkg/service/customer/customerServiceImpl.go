package service

import (
	"context"
	"go-clean-architecture/pkg/dto"
	"go-clean-architecture/pkg/entity"
	customerRepo "go-clean-architecture/pkg/repository/customer"
)

func NewCustomerService(customerRepository *customerRepo.CustomerRepository) CustomerService {
	return &customerServiceImpl{
		customerRepository: *customerRepository,
	}
}

type customerServiceImpl struct {
	customerRepository customerRepo.CustomerRepository
}

func (service *customerServiceImpl) GetAll() []entity.Customer {
	customers := service.customerRepository.GetAll()
	return customers
}

func (service *customerServiceImpl) Get(id string) (entity.Customer, error) {
	customer, err := service.customerRepository.Get(id)
	return customer, err
}

func (service *customerServiceImpl) GetMyInfo(ctx context.Context) *dto.UserClaims {
	myInfo := ctx.Value("userInfo").(*dto.UserClaims)
	return myInfo
}
