package service

import (
	"go-clean-architecture/pkg/entity"
	productRepo "go-clean-architecture/pkg/repository/product"
	customerServ "go-clean-architecture/pkg/service/customer"
)

func NewProductService(productRepository *productRepo.ProductRepo, customerService *customerServ.CustomerService) ProductService {
	return &ProductServiceImpl{
		ProductRepository: *productRepository,
		CustomerService:   *customerService,
	}
}

type ProductServiceImpl struct {
	ProductRepository productRepo.ProductRepo
	CustomerService   customerServ.CustomerService
}

func (service *ProductServiceImpl) GetAll() []entity.Product {
	products := service.ProductRepository.GetAll()
	return products
}

func (service *ProductServiceImpl) Get(name string) (entity.Product, error) {
	product, err := service.ProductRepository.Get(name)
	return product, err
}
