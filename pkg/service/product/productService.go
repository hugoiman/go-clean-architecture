package service

import "go-clean-architecture/pkg/entity"

type ProductService interface {
	GetAll() []entity.Product
	Get(name string) (entity.Product, error)
}
