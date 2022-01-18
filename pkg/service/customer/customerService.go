package service

import "go-clean-architecture/pkg/entity"

type CustomerService interface {
	GetAll() []entity.Customer
	Get(id string) (entity.Customer, error)
}
