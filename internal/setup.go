package internal

import (
	"go-clean-architecture/config"

	authContr "go-clean-architecture/pkg/controller/auth"
	"go-clean-architecture/pkg/middleware"
	authRepo "go-clean-architecture/pkg/repository/auth"
	authServ "go-clean-architecture/pkg/service/auth"

	customerContr "go-clean-architecture/pkg/controller/customer"
	customerRepo "go-clean-architecture/pkg/repository/customer"
	customerServ "go-clean-architecture/pkg/service/customer"

	productContr "go-clean-architecture/pkg/controller/product"
	productRepo "go-clean-architecture/pkg/repository/product"
	productServ "go-clean-architecture/pkg/service/product"

	"github.com/gorilla/mux"
)

func Setup(router *mux.Router) {
	db := config.NewDB()

	// Setup Middleware
	mw := middleware.NewMiddleware()
	auth := router.PathPrefix("").Subrouter()
	auth.Use(mw.Auth)

	// Setup Customer
	customerRepository := customerRepo.NewCustomerRepository(&db)
	customerService := customerServ.NewCustomerService(&customerRepository)
	customerController := customerContr.NewCustomerController(&customerService)
	customerController.Route(router, auth)

	// Setup Auth
	authRepository := authRepo.NewAuthRepository(&db)
	authService := authServ.NewAuthService(&authRepository, &customerService)
	authController := authContr.NewAuthController(&authService)
	authController.Route(router)

	// Setup Product
	productRepository := productRepo.NewProductRepository(&db)
	productService := productServ.NewProductService(&productRepository)
	productController := productContr.NewProductController(&productService)
	productController.Route(router)
}
