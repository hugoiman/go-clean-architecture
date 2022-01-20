package repo

import "go-clean-architecture/pkg/entity"

type ProductRepository interface {
	GetAll() []entity.Product
	Get(name string) (entity.Product, error)
}
