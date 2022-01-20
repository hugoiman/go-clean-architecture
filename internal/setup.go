package internal

import (
	"go-clean-architecture/config"

	customerContr "go-clean-architecture/pkg/controller/customer"
	customerRepo "go-clean-architecture/pkg/repository/customer"
	customerServ "go-clean-architecture/pkg/service/customer"

	productContr "go-clean-architecture/pkg/controller/product"
	productRepo "go-clean-architecture/pkg/repository/product"
	productServ "go-clean-architecture/pkg/service/product"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	var db *config.DB

	// Setup Customer
	customerRepository := customerRepo.NewCustomerRepository(db)
	customerService := customerServ.NewCustomerService(&customerRepository)
	customerController := customerContr.NewCustomerController(&customerService)
	customerController.Route(router)

	// Setup Product
	productRepository := productRepo.NewProductRepository(db)
	productService := productServ.NewProductService(&productRepository)
	productController := productContr.NewProductController(&productService)
	productController.Route(router)
}
