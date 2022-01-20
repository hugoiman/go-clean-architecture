package controllers

import (
	"encoding/json"
	"go-clean-architecture/pkg/dto"
	"net/http"

	authServ "go-clean-architecture/pkg/service/auth"

	"github.com/gorilla/mux"
)

func NewAuthController(authService *authServ.AuthService) authController {
	return authController{
		authService: *authService,
	}
}

type authController struct {
	authService authServ.AuthService
}

func (controller *authController) Route(router *mux.Router) {
	router.HandleFunc("/login", controller.Login).Methods("POST")
}

func (controller *authController) Login(w http.ResponseWriter, r *http.Request) {
	var userCredentials dto.UserCredential
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := controller.authService.Login(userCredentials)
	if err != nil {
		http.Error(w, "Username atau password salah.", http.StatusBadRequest)
		return
	}

	token := controller.authService.GenerateToken(customer)
	message, _ := json.Marshal(token)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token":` + string(message) + `}`))
}
