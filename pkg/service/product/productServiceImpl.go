package service

import (
	"go-clean-architecture/pkg/entity"
	productRepo "go-clean-architecture/pkg/repository/product"
)

func NewProductService(productRepository *productRepo.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepository: *productRepository,
	}
}

type productServiceImpl struct {
	productRepository productRepo.ProductRepository
}

func (service *productServiceImpl) GetAll() []entity.Product {
	products := service.productRepository.GetAll()
	return products
}

func (service *productServiceImpl) Get(name string) (entity.Product, error) {
	product, err := service.productRepository.Get(name)
	return product, err
}
