package repo

import "go-clean-architecture/pkg/entity"

type CustomerRepository interface {
	GetAll() []entity.Customer
	Get(username string) (entity.Customer, error)
}
