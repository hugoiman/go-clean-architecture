package controller

import (
	"encoding/json"
	"net/http"

	productServ "go-clean-architecture/pkg/service/product"

	"github.com/gorilla/mux"
)

func NewProductController(productService *productServ.ProductService) productController {
	return productController{
		productService: *productService,
	}
}

type productController struct {
	productService productServ.ProductService
}

func (controller *productController) Route(router *mux.Router) {
	router.HandleFunc("/product/{name}", controller.Get).Methods("GET")
	router.HandleFunc("/product", controller.GetAll).Methods("GET")
}

func (controller *productController) GetAll(w http.ResponseWriter, r *http.Request) {
	products := controller.productService.GetAll()
	message, _ := json.Marshal(&products)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (controller *productController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productName := vars["name"]

	product, err := controller.productService.Get(productName)
	if err != nil {
		http.Error(w, "not found bos", http.StatusBadRequest)
		return
	}
	message, _ := json.Marshal(&product)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
