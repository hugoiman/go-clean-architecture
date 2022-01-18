package repo

import "go-clean-architecture/pkg/entity"

type CustomerRepo interface {
	GetAll() []entity.Customer
	Get(username string) (entity.Customer, error)
}
