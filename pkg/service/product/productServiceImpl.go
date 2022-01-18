package service

import (
	"go-clean-architecture/pkg/entity"
	productRepo "go-clean-architecture/pkg/repository/product"
)

func NewProductService(productRepository *productRepo.ProductRepo) ProductService {
	return &ProductServiceImpl{
		ProductRepository: *productRepository,
	}
}

type ProductServiceImpl struct {
	ProductRepository productRepo.ProductRepo
}

func (service *ProductServiceImpl) GetAll() []entity.Product {
	products := service.ProductRepository.GetAll()
	return products
}

func (service *ProductServiceImpl) Get(name string) (entity.Product, error) {
	product, err := service.ProductRepository.Get(name)
	return product, err
}
