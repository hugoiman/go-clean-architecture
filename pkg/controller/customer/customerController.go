package controller

import (
	"encoding/json"
	"net/http"

	customerServ "go-clean-architecture/pkg/service/customer"

	"github.com/gorilla/mux"
)

func NewCustomerController(customerService *customerServ.CustomerService) customerController {
	return customerController{
		customerService: *customerService,
	}
}

type customerController struct {
	customerService customerServ.CustomerService
}

func (controller *customerController) Route(router, auth *mux.Router) {
	router.HandleFunc("/customer/{username}", controller.Get).Methods("GET")
	auth.HandleFunc("/customer", controller.GetAll).Methods("GET")
}

func (controller *customerController) GetAll(w http.ResponseWriter, r *http.Request) {
	customers := controller.customerService.GetAll()
	message, _ := json.Marshal(&customers)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (controller *customerController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	customer, err := controller.customerService.Get(username)
	if err != nil {
		http.Error(w, "not found bos"+err.Error(), http.StatusBadRequest)
		return
	}
	message, _ := json.Marshal(&customer)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
