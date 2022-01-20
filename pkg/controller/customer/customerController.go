package controller

import (
	"encoding/json"
	"net/http"

	customerServ "go-clean-architecture/pkg/service/customer"

	"github.com/gorilla/mux"
)

func NewCustomerController(customerService *customerServ.CustomerService) CustomerController {
	return CustomerController{
		CustomerService: *customerService,
	}
}

type CustomerController struct {
	CustomerService customerServ.CustomerService
}

func (controller *CustomerController) Route(router *mux.Router) {
	router.HandleFunc("/customer/{username}", controller.Get).Methods("GET")
	router.HandleFunc("/customer", controller.GetAll).Methods("GET")
}

func (controller *CustomerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers := controller.CustomerService.GetAll()
	message, _ := json.Marshal(&customers)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (controller *CustomerController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	customer, err := controller.CustomerService.Get(username)
	if err != nil {
		http.Error(w, "not found bos"+err.Error(), http.StatusBadRequest)
		return
	}
	message, _ := json.Marshal(&customer)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
